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
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/time/rate"
)

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
			http.Error(w, "Limite de requisições excedido", http.StatusTooManyRequests)
			return
		}

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
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Verificar se o email já está em uso
	for _, u := range users {
		if u.Email == req.Email {
			http.Error(w, "Email já está em uso", http.StatusBadRequest)
			return
		}
	}

	// Gerar hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao processar senha", http.StatusInternalServerError)
		return
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
		http.Error(w, "Erro ao gerar tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenPair)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
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
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Verificar senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		http.Error(w, "Credenciais inválidas", http.StatusUnauthorized)
		return
	}

	// Gerar tokens
	tokenPair, err := generateTokenPair(user.ID)
	if err != nil {
		http.Error(w, "Erro ao gerar tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenPair)
}

func refreshHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req RefreshRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	// Validar refresh token
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.RefreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	// Verificar se é um refresh token
	if claims.Type != "refresh" {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	// Gerar novo par de tokens
	tokenPair, err := generateTokenPair(claims.UserID)
	if err != nil {
		http.Error(w, "Erro ao gerar tokens", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tokenPair)
}

func validateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Token string `json:"token"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(req.Token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(claims)
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

	// Rotas públicas com rate limiting
	http.HandleFunc("/auth/register", rateLimitMiddleware(registerHandler))
	http.HandleFunc("/auth/login", rateLimitMiddleware(loginHandler))
	http.HandleFunc("/auth/refresh", rateLimitMiddleware(refreshHandler))
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})

	// Rotas protegidas com rate limiting e autenticação
	http.HandleFunc("/auth/validate", rateLimitMiddleware(authMiddleware(validateHandler)))

	fmt.Println("Auth Service rodando na porta 8081")
	http.ListenAndServe(":8081", nil)
}
