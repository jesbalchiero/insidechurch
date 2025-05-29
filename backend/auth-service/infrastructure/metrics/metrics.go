package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Duração das requisições HTTP
	HTTPDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_duration_seconds",
			Help:    "Duração das requisições HTTP em segundos",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"path", "method", "status"},
	)

	// Total de requisições HTTP
	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP",
		},
		[]string{"path", "method", "status"},
	)

	// Requisições HTTP em andamento
	HTTPRequestsInFlight = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Número de requisições HTTP em andamento",
		},
	)

	// Métricas de autenticação
	LoginAttempts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_login_attempts_total",
			Help: "Total de tentativas de login",
		},
		[]string{"status"}, // success, failure
	)

	RegisterAttempts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_register_attempts_total",
			Help: "Total de tentativas de registro",
		},
		[]string{"status"}, // success, failure
	)

	TokenValidations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_token_validations_total",
			Help: "Total de validações de token",
		},
		[]string{"status"}, // valid, invalid
	)

	// Métricas de rate limiting
	RateLimitHits = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rate_limit_hits_total",
			Help: "Total de hits no rate limiter",
		},
		[]string{"status"}, // allowed, rejected
	)

	// Métricas de usuários
	ActiveUsers = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "auth_active_users",
			Help: "Número de usuários ativos",
		},
	)

	// Métricas de tokens
	TokenGenerations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "auth_token_generations_total",
			Help: "Total de gerações de token",
		},
		[]string{"type"}, // access, refresh
	)

	TokenRefreshes = promauto.NewCounter(
		prometheus.CounterOpts{
			Name: "auth_token_refreshes_total",
			Help: "Total de refreshs de token",
		},
	)
)
