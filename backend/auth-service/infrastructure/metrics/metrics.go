package metrics

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Métricas HTTP
	HTTPRequestsInFlight = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "http_requests_in_flight",
			Help: "Número de requisições HTTP em andamento",
		},
	)

	HTTPDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duração das requisições HTTP em segundos",
			Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
		},
		[]string{"path", "method", "status"},
	)

	HTTPRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP",
		},
		[]string{"path", "method", "status"},
	)

	httpDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "http_request_duration_seconds",
			Help: "Duração das requisições HTTP em segundos",
		},
		[]string{"path", "method"},
	)

	httpRequests = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total de requisições HTTP",
		},
		[]string{"path", "method", "status"},
	)

	httpErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_errors_total",
			Help: "Total de erros HTTP",
		},
		[]string{"path", "method", "status"},
	)

	// Métricas de Rate Limiting
	RateLimitHits = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "rate_limit_hits_total",
			Help: "Total de hits no rate limiting",
		},
		[]string{"endpoint", "status"},
	)

	// Métricas de Autenticação
	LoginAttempts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "login_attempts_total",
			Help: "Total de tentativas de login",
		},
		[]string{"status"},
	)

	RegisterAttempts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "register_attempts_total",
			Help: "Total de tentativas de registro",
		},
		[]string{"status"},
	)

	ActiveUsers = promauto.NewGauge(
		prometheus.GaugeOpts{
			Name: "active_users",
			Help: "Número de usuários ativos",
		},
	)

	TokenGenerations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "token_generations_total",
			Help: "Total de gerações de tokens",
		},
		[]string{"type", "status"},
	)

	TokenValidations = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "token_validations_total",
			Help: "Total de validações de tokens",
		},
		[]string{"type", "status"},
	)

	TokenRefreshes = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "token_refreshes_total",
			Help: "Total de refreshs de tokens",
		},
		[]string{"status"},
	)

	// Métricas de Banco de Dados
	dbDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "db_operation_duration_seconds",
			Help: "Duração das operações no banco de dados em segundos",
		},
		[]string{"database", "operation"},
	)

	dbErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "db_errors_total",
			Help: "Total de erros no banco de dados",
		},
		[]string{"database", "operation"},
	)

	// Métricas de Cache
	cacheDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "cache_operation_duration_seconds",
			Help: "Duração das operações no cache em segundos",
		},
		[]string{"cache", "operation"},
	)

	cacheErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "cache_errors_total",
			Help: "Total de erros no cache",
		},
		[]string{"cache", "operation"},
	)

	// Métricas de API Externa
	externalAPIDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "external_api_duration_seconds",
			Help: "Duração das chamadas à API externa em segundos",
		},
		[]string{"api", "method"},
	)

	externalAPIErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "external_api_errors_total",
			Help: "Total de erros em chamadas à API externa",
		},
		[]string{"api", "method"},
	)

	// Métricas de Eventos
	eventProcessingDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "event_processing_duration_seconds",
			Help: "Duração do processamento de eventos em segundos",
		},
		[]string{"event"},
	)

	eventErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "event_errors_total",
			Help: "Total de erros no processamento de eventos",
		},
		[]string{"event"},
	)

	// Métricas de Notificações
	notificationsSent = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "notifications_sent_total",
			Help: "Total de notificações enviadas",
		},
		[]string{"notification", "channel"},
	)

	notificationErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "notification_errors_total",
			Help: "Total de erros no envio de notificações",
		},
		[]string{"notification", "channel"},
	)

	// Métricas de Usuário
	userActions = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_actions_total",
			Help: "Total de ações de usuário",
		},
		[]string{"user", "action"},
	)

	userErrors = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "user_errors_total",
			Help: "Total de erros em ações de usuário",
		},
		[]string{"user", "action"},
	)
)

// Funções para registrar métricas HTTP
func RecordHTTPDuration(path, method string, duration float64) {
	httpDuration.WithLabelValues(path, method).Observe(duration)
}

func RecordHTTPRequest(path, method string, status int) {
	httpRequests.WithLabelValues(path, method, strconv.Itoa(status)).Inc()
}

func RecordHTTPError(path, method string, status int) {
	httpErrors.WithLabelValues(path, method, strconv.Itoa(status)).Inc()
}

// Funções para registrar métricas de banco de dados
func RecordDatabaseOperation(database, operation string, duration time.Duration) {
	dbDuration.WithLabelValues(database, operation).Observe(duration.Seconds())
}

func RecordDatabaseError(database, operation string) {
	dbErrors.WithLabelValues(database, operation).Inc()
}

// Funções para registrar métricas de cache
func RecordCacheOperation(cache, operation string, duration time.Duration) {
	cacheDuration.WithLabelValues(cache, operation).Observe(duration.Seconds())
}

func RecordCacheError(cache, operation string) {
	cacheErrors.WithLabelValues(cache, operation).Inc()
}

// Funções para registrar métricas de API externa
func RecordExternalAPICall(api, method string, duration time.Duration) {
	externalAPIDuration.WithLabelValues(api, method).Observe(duration.Seconds())
}

func RecordExternalAPIError(api, method string) {
	externalAPIErrors.WithLabelValues(api, method).Inc()
}

// Funções para registrar métricas de eventos
func RecordEventProcessing(event string, duration time.Duration) {
	eventProcessingDuration.WithLabelValues(event).Observe(duration.Seconds())
}

func RecordEventError(event string) {
	eventErrors.WithLabelValues(event).Inc()
}

// Funções para registrar métricas de notificações
func RecordNotificationSent(notification, channel string) {
	notificationsSent.WithLabelValues(notification, channel).Inc()
}

func RecordNotificationError(notification, channel string) {
	notificationErrors.WithLabelValues(notification, channel).Inc()
}

// Funções para registrar métricas de usuário
func RecordUserAction(user, action string) {
	userActions.WithLabelValues(user, action).Inc()
}

func RecordUserError(user, action string) {
	userErrors.WithLabelValues(user, action).Inc()
}
