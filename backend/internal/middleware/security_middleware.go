package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/time/rate"
)

type SecurityMiddleware struct {
	limiter *rate.Limiter
	ips     map[string]*rate.Limiter
	mu      sync.Mutex
}

func NewSecurityMiddleware() *SecurityMiddleware {
	return &SecurityMiddleware{
		limiter: rate.NewLimiter(rate.Every(time.Second), 100),
		ips:     make(map[string]*rate.Limiter),
	}
}

func (m *SecurityMiddleware) SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Headers básicos de segurança
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		// Proteção contra ataques de timing
		c.Header("X-DNS-Prefetch-Control", "off")
		c.Header("X-Download-Options", "noopen")

		// Proteção contra cache
		c.Header("Cache-Control", "no-store, no-cache, must-revalidate, proxy-revalidate")
		c.Header("Pragma", "no-cache")
		c.Header("Expires", "0")

		// CSP específico para o Swagger UI
		if strings.HasPrefix(c.Request.URL.Path, "/api/swagger") {
			c.Header("Content-Security-Policy", "default-src 'self'; "+
				"script-src 'self' 'unsafe-inline' 'unsafe-eval'; "+
				"style-src 'self' 'unsafe-inline'; "+
				"img-src 'self' data:; "+
				"font-src 'self' data:; "+
				"connect-src 'self'; "+
				"frame-ancestors 'none'; "+
				"form-action 'self'; "+
				"base-uri 'self'; "+
				"object-src 'none'")
		} else {
			c.Header("Content-Security-Policy", "default-src 'self'; "+
				"script-src 'self'; "+
				"style-src 'self'; "+
				"img-src 'self'; "+
				"font-src 'self'; "+
				"connect-src 'self'; "+
				"frame-ancestors 'none'; "+
				"form-action 'self'; "+
				"base-uri 'self'; "+
				"object-src 'none'")
		}

		c.Next()
	}
}

func (m *SecurityMiddleware) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		m.mu.Lock()
		limiter, exists := m.ips[ip]
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Minute), 100)
			m.ips[ip] = limiter
		}
		m.mu.Unlock()

		if !limiter.Allow() {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"error": "muitas requisições, tente novamente mais tarde",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func (m *SecurityMiddleware) CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Configuração mais restritiva do CORS
		origin := c.Request.Header.Get("Origin")
		allowedOrigins := []string{"http://localhost:3000"}

		// Verifica se a origem está na lista de permitidas
		allowed := false
		for _, allowedOrigin := range allowedOrigins {
			if origin == allowedOrigin {
				allowed = true
				break
			}
		}

		if allowed {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			c.Header("Access-Control-Expose-Headers", "Content-Length")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Max-Age", "86400") // 24 horas
		}

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// SwaggerSecurityHeaders retorna um middleware com headers de segurança específicos para o Swagger UI
func (m *SecurityMiddleware) SwaggerSecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Headers básicos de segurança
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "SAMEORIGIN")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")

		// CSP específico para o Swagger UI
		c.Header("Content-Security-Policy", "default-src 'self'; "+
			"script-src 'self' 'unsafe-inline' 'unsafe-eval'; "+
			"style-src 'self' 'unsafe-inline'; "+
			"img-src 'self' data:; "+
			"font-src 'self' data:; "+
			"connect-src 'self'")

		c.Next()
	}
}
