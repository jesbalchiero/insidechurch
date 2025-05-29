package metrics

import (
	"net/http"
	"strconv"
	"time"
)

// MetricsMiddleware é um middleware que coleta métricas do Prometheus
func MetricsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Incrementar contador de requisições em andamento
		HTTPRequestsInFlight.Inc()
		defer HTTPRequestsInFlight.Dec()

		// Criar um ResponseWriter personalizado para capturar o status code
		rw := newResponseWriter(w)

		// Registrar início da requisição
		start := time.Now()

		// Chamar o próximo handler
		next.ServeHTTP(rw, r)

		// Registrar duração da requisição
		duration := time.Since(start).Seconds()
		HTTPDuration.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(rw.statusCode)).Observe(duration)

		// Incrementar contador total de requisições
		HTTPRequestsTotal.WithLabelValues(r.URL.Path, r.Method, strconv.Itoa(rw.statusCode)).Inc()
	})
}

// responseWriter é um wrapper para http.ResponseWriter que captura o status code
type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

// newResponseWriter cria um novo responseWriter
func newResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{w, http.StatusOK}
}

// WriteHeader sobrescreve o método WriteHeader para capturar o status code
func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
