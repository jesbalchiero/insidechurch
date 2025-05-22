package middleware

import (
	"net/http"
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
		limiter: rate.NewLimiter(rate.Every(time.Second), 100), // 100 requisições por segundo
		ips:     make(map[string]*rate.Limiter),
	}
}

func (m *SecurityMiddleware) SecurityHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Headers de segurança
		c.Header("X-Content-Type-Options", "nosniff")
		c.Header("X-Frame-Options", "DENY")
		c.Header("X-XSS-Protection", "1; mode=block")
		c.Header("Strict-Transport-Security", "max-age=31536000; includeSubDomains")
		c.Header("Content-Security-Policy", "default-src 'self'")
		c.Header("Referrer-Policy", "strict-origin-when-cross-origin")
		c.Header("Permissions-Policy", "geolocation=(), microphone=(), camera=()")

		c.Next()
	}
}

func (m *SecurityMiddleware) RateLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()

		m.mu.Lock()
		limiter, exists := m.ips[ip]
		if !exists {
			limiter = rate.NewLimiter(rate.Every(time.Minute), 100) // 100 requisições por minuto por IP
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
		c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length")
		c.Header("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
