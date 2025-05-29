package tracing

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
)

var (
	tracer trace.Tracer
)

// InitTracer inicializa o tracer do OpenTelemetry
func InitTracer(serviceName string) error {
	// Configurar o endpoint do coletor
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if endpoint == "" {
		endpoint = "localhost:4317" // Endpoint padrão do OpenTelemetry Collector
	}

	// Criar o cliente gRPC
	client := otlptracegrpc.NewClient(
		otlptracegrpc.WithInsecure(),
		otlptracegrpc.WithEndpoint(endpoint),
		otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)

	// Criar o exportador
	exporter, err := otlptrace.New(context.Background(), client)
	if err != nil {
		return fmt.Errorf("falha ao criar exportador OTLP: %v", err)
	}

	// Criar o recurso com informações do serviço
	res, err := resource.New(
		context.Background(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String("1.0.0"),
		),
	)
	if err != nil {
		return fmt.Errorf("falha ao criar recurso: %v", err)
	}

	// Configurar o provedor de traces
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)

	// Configurar propagação de contexto
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	// Configurar o tracer global
	otel.SetTracerProvider(tp)
	tracer = tp.Tracer(serviceName)

	return nil
}

// TracedHandler é um middleware que adiciona tracing às requisições HTTP
func TracedHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extrair o contexto do trace da requisição
		ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))

		// Criar um novo span
		ctx, span := tracer.Start(ctx, r.URL.Path,
			trace.WithAttributes(
				attribute.String("http.method", r.Method),
				attribute.String("http.url", r.URL.String()),
				attribute.String("http.user_agent", r.UserAgent()),
				attribute.String("http.remote_addr", r.RemoteAddr),
			),
		)
		defer span.End()

		// Criar um ResponseWriter personalizado para capturar o status code
		rw := newResponseWriter(w)

		// Chamar o próximo handler
		next.ServeHTTP(rw, r.WithContext(ctx))

		// Adicionar o status code ao span
		span.SetAttributes(attribute.Int("http.status_code", rw.statusCode))
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

// TraceSpan cria um novo span para uma operação
func TraceSpan(ctx context.Context, name string, fn func(context.Context) error) error {
	ctx, span := tracer.Start(ctx, name)
	defer span.End()

	err := fn(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}

// TraceSpanWithAttributes cria um novo span com atributos para uma operação
func TraceSpanWithAttributes(ctx context.Context, name string, attributes map[string]string, fn func(context.Context) error) error {
	ctx, span := tracer.Start(ctx, name)
	defer span.End()

	// Adicionar atributos ao span
	for k, v := range attributes {
		span.SetAttributes(attribute.String(k, v))
	}

	err := fn(ctx)
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}

	return err
}
