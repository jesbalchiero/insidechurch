package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"

	"github.com/insidechurch/auth-service/infrastructure/logger"
	"github.com/insidechurch/auth-service/infrastructure/metrics"
	"github.com/insidechurch/auth-service/infrastructure/tracing"
)

var (
	log *logger.JSONLogger
)

func init() {
	var err error
	log, err = logger.NewJSONLogger("logs/app.log")
	if err != nil {
		fmt.Printf("Erro ao inicializar logger: %v\n", err)
		os.Exit(1)
	}
}

// Tipo personalizado para chaves do contexto
type contextKey string

const (
	userIDKey contextKey = "user_id"
)

// Estrutura para armazenar os limitadores por usuário
type userRateLimiter struct {
	limiters map[string]*rate.Limiter
	mu       sync.RWMutex
}

// Configurações do rate limiter
const (
	requestsPerMinute = 60 // 60 requisições por minuto
	burstSize         = 10 // Permite até 10 requisições em burst
)

var (
	limiter = &userRateLimiter{
		limiters: make(map[string]*rate.Limiter),
	}
)

// Obtém ou cria um limitador para um usuário específico
func (rl *userRateLimiter) getLimiter(userID string) *rate.Limiter {
	rl.mu.RLock()
	limiter, exists := rl.limiters[userID]
	rl.mu.RUnlock()

	if exists {
		return limiter
	}

	rl.mu.Lock()
	defer rl.mu.Unlock()

	// Verificar novamente após adquirir o lock (double-check pattern)
	limiter, exists = rl.limiters[userID]
	if exists {
		return limiter
	}

	// Criar novo limitador
	limiter = rate.NewLimiter(rate.Every(time.Minute/requestsPerMinute), burstSize)
	rl.limiters[userID] = limiter
	return limiter
}

// Middleware de rate limiting
func rateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obter o ID do usuário do contexto (se disponível)
		userID := "anonymous"
		if claims, ok := r.Context().Value(userIDKey).(string); ok {
			userID = claims
		}

		// Obter o limitador para o usuário
		limiter := limiter.getLimiter(userID)

		// Verificar se a requisição pode prosseguir
		if !limiter.Allow() {
			metrics.RateLimitHits.WithLabelValues("rejected").Inc()
			http.Error(w, "Limite de requisições excedido", http.StatusTooManyRequests)
			return
		}

		metrics.RateLimitHits.WithLabelValues("allowed").Inc()
		next.ServeHTTP(w, r)
	}
}

type User struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	UserID string `json:"user_id"`
	Type   string `json:"type"` // "access" ou "refresh"
	jwt.RegisteredClaims
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type ValidateRequest struct {
	Token string `json:"token"`
}

type ValidateResponse struct {
	Valid bool `json:"valid"`
}

var jwtKey = []byte(os.Getenv("JWT_SECRET"))
var users = []User{
	{
		ID:       "1",
		Email:    "admin@insidechurch.com",
		Password: "$2a$10$XOPbrlUPQdwdJUpSrIF6X.LbE14qsMmKGhM1A8W9iqDp0JxX5QKHy", // "admin123"
	},
}

func generateTokenPair(userID string) (*TokenPair, error) {
	// Gerar access token (15 minutos)
	accessClaims := &Claims{
		UserID: userID,
		Type:   "access",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
	accessTokenString, err := accessToken.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	// Gerar refresh token (7 dias)
	refreshClaims := &Claims{
		UserID: userID,
		Type:   "refresh",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	refreshTokenString, err := refreshToken.SignedString(jwtKey)
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := tracing.TraceSpanWithAttributes(ctx, "register", map[string]string{
		"method": r.Method,
		"path":   r.URL.Path,
		"ip":     r.RemoteAddr,
	}, func(ctx context.Context) error {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return nil
		}

		var req RegisterRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("Erro ao decodificar requisição de registro", err,
				logger.String("method", r.Method),
				logger.String("path", r.URL.Path),
				logger.String("ip", r.RemoteAddr),
			)
			metrics.RegisterAttempts.WithLabelValues("failure").Inc()
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return nil
		}

		// Verificar se o email já está em uso
		for _, u := range users {
			if u.Email == req.Email {
				log.Info("Tentativa de registro com email já existente",
					logger.String("email", req.Email),
					logger.String("ip", r.RemoteAddr),
				)
				metrics.RegisterAttempts.WithLabelValues("failure").Inc()
				http.Error(w, "Email já está em uso", http.StatusBadRequest)
				return nil
			}
		}

		// Gerar hash da senha
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Error("Erro ao gerar hash da senha", err,
				logger.String("email", req.Email),
			)
			metrics.RegisterAttempts.WithLabelValues("failure").Inc()
			http.Error(w, "Erro ao processar senha", http.StatusInternalServerError)
			return err
		}

		// Criar novo usuário
		newUser := User{
			ID:       fmt.Sprintf("%d", len(users)+1),
			Email:    req.Email,
			Password: string(hashedPassword),
		}
		users = append(users, newUser)

		// Gerar tokens
		tokenPair, err := generateTokenPair(newUser.ID)
		if err != nil {
			log.Error("Erro ao gerar tokens", err,
				logger.String("user_id", newUser.ID),
			)
			metrics.RegisterAttempts.WithLabelValues("failure").Inc()
			http.Error(w, "Erro ao gerar tokens", http.StatusInternalServerError)
			return err
		}

		metrics.RegisterAttempts.WithLabelValues("success").Inc()
		metrics.ActiveUsers.Inc()
		metrics.TokenGenerations.WithLabelValues("access").Inc()
		metrics.TokenGenerations.WithLabelValues("refresh").Inc()

		log.Info("Usuário registrado com sucesso",
			logger.String("user_id", newUser.ID),
			logger.String("email", newUser.Email),
		)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenPair)
		return nil
	})

	if err != nil {
		log.Error("Erro no handler de registro", err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := tracing.TraceSpanWithAttributes(ctx, "login", map[string]string{
		"method": r.Method,
		"path":   r.URL.Path,
		"ip":     r.RemoteAddr,
	}, func(ctx context.Context) error {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return nil
		}

		var req LoginRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Error("Erro ao decodificar requisição de login", err,
				logger.String("method", r.Method),
				logger.String("path", r.URL.Path),
				logger.String("ip", r.RemoteAddr),
			)
			metrics.LoginAttempts.WithLabelValues("failure").Inc()
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return nil
		}

		// Encontrar usuário
		var user *User
		for _, u := range users {
			if u.Email == req.Email {
				user = &u
				break
			}
		}

		if user == nil {
			log.Info("Tentativa de login com email não encontrado",
				logger.String("email", req.Email),
				logger.String("ip", r.RemoteAddr),
			)
			metrics.LoginAttempts.WithLabelValues("failure").Inc()
			http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
			return nil
		}

		// Verificar senha
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
			log.Info("Tentativa de login com senha inválida",
				logger.String("email", req.Email),
				logger.String("ip", r.RemoteAddr),
			)
			metrics.LoginAttempts.WithLabelValues("failure").Inc()
			http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
			return nil
		}

		// Gerar tokens
		tokenPair, err := generateTokenPair(user.ID)
		if err != nil {
			log.Error("Erro ao gerar tokens", err,
				logger.String("user_id", user.ID),
			)
			metrics.LoginAttempts.WithLabelValues("failure").Inc()
			http.Error(w, "Erro ao gerar tokens", http.StatusInternalServerError)
			return err
		}

		metrics.LoginAttempts.WithLabelValues("success").Inc()
		metrics.TokenGenerations.WithLabelValues("access").Inc()
		metrics.TokenGenerations.WithLabelValues("refresh").Inc()

		log.Info("Login realizado com sucesso",
			logger.String("user_id", user.ID),
			logger.String("email", user.Email),
		)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenPair)
		return nil
	})

	if err != nil {
		log.Error("Erro no handler de login", err)
	}
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := tracing.TraceSpanWithAttributes(ctx, "refresh", map[string]string{
		"method": r.Method,
		"path":   r.URL.Path,
		"ip":     r.RemoteAddr,
	}, func(ctx context.Context) error {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return nil
		}

		var req RefreshRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			metrics.TokenValidations.WithLabelValues("invalid").Inc()
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return nil
		}

		// Validar refresh token
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			metrics.TokenValidations.WithLabelValues("invalid").Inc()
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return nil
		}

		// Verificar se é um refresh token
		if claims.Type != "refresh" {
			metrics.TokenValidations.WithLabelValues("invalid").Inc()
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return nil
		}

		// Gerar novo par de tokens
		tokenPair, err := generateTokenPair(claims.UserID)
		if err != nil {
			metrics.TokenValidations.WithLabelValues("invalid").Inc()
			http.Error(w, "Erro ao gerar tokens", http.StatusInternalServerError)
			return err
		}

		metrics.TokenValidations.WithLabelValues("valid").Inc()
		metrics.TokenRefreshes.Inc()
		metrics.TokenGenerations.WithLabelValues("access").Inc()
		metrics.TokenGenerations.WithLabelValues("refresh").Inc()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenPair)
		return nil
	})

	if err != nil {
		log.Error("Erro no handler de refresh", err)
	}
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	err := tracing.TraceSpanWithAttributes(ctx, "validate", map[string]string{
		"method": r.Method,
		"path":   r.URL.Path,
		"ip":     r.RemoteAddr,
	}, func(ctx context.Context) error {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return nil
		}

		var req struct {
			Token string `json:"token"`
		}

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			metrics.TokenValidations.WithLabelValues("invalid").Inc()
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return nil
		}

		claims := &Claims{}
		token, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			metrics.TokenValidations.WithLabelValues("invalid").Inc()
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return nil
		}

		metrics.TokenValidations.WithLabelValues("valid").Inc()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(claims)
		return nil
	})

	if err != nil {
		log.Error("Erro no handler de validação", err)
	}
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Token não fornecido", http.StatusUnauthorized)
			return
		}

		// Extrair o token do header "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "Formato de token inválido", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		claims := &Claims{}
		parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !parsedToken.Valid {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Verificar se é um access token
		if claims.Type != "access" {
			http.Error(w, "Token inválido", http.StatusUnauthorized)
			return
		}

		// Adicionar o ID do usuário ao contexto da requisição usando o tipo personalizado
		ctx := context.WithValue(r.Context(), userIDKey, claims.UserID)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

func main() {
	// Definir JWT_SECRET se não estiver definido
	if os.Getenv("JWT_SECRET") == "" {
		os.Setenv("JWT_SECRET", "sua_chave_secreta_aqui")
	}

	// Inicializar o logger de auditoria
	if err := initAuditLogger(); err != nil {
		log.Error("Erro ao inicializar logger de auditoria", err)
		os.Exit(1)
	}
	defer auditLogger.Close()

	// Inicializar o tracer
	if err := tracing.InitTracer("auth-service"); err != nil {
		log.Error("Erro ao inicializar tracer", err)
		os.Exit(1)
	}

	// Rotas públicas com rate limiting, auditoria, métricas e tracing
	http.Handle("/auth/register", tracing.TracedHandler(metrics.MetricsMiddleware(auditMiddleware(rateLimitMiddleware(registerHandler)))))
	http.Handle("/auth/login", tracing.TracedHandler(metrics.MetricsMiddleware(auditMiddleware(rateLimitMiddleware(loginHandler)))))
	http.Handle("/auth/refresh", tracing.TracedHandler(metrics.MetricsMiddleware(auditMiddleware(rateLimitMiddleware(refreshHandler)))))
	http.Handle("/health", tracing.TracedHandler(metrics.MetricsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	}))))

	// Rotas protegidas com rate limiting, autenticação, auditoria, métricas e tracing
	http.Handle("/auth/validate", tracing.TracedHandler(metrics.MetricsMiddleware(auditMiddleware(rateLimitMiddleware(authMiddleware(validateHandler))))))

	// Endpoint do Prometheus
	http.Handle("/metrics", promhttp.Handler())

	log.Info("Serviço de autenticação iniciado",
		logger.String("port", "8081"),
	)

	fmt.Println("Auth Service rodando na porta 8081")
	http.ListenAndServe(":8081", nil)
}
