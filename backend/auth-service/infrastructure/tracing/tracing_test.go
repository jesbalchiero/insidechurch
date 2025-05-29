package tracing

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestInitTracer(t *testing.T) {
	// Teste com endpoint padrão
	err := InitTracer("test-service")
	if err != nil {
		t.Errorf("InitTracer falhou com endpoint padrão: %v", err)
	}

	// Teste com endpoint personalizado
	os.Setenv("OTEL_EXPORTER_OTLP_ENDPOINT", "localhost:4318")
	err = InitTracer("test-service")
	if err != nil {
		t.Errorf("InitTracer falhou com endpoint personalizado: %v", err)
	}
}

func TestTracedHandler(t *testing.T) {
	// Inicializar o tracer
	err := InitTracer("test-service")
	if err != nil {
		t.Fatalf("Falha ao inicializar tracer: %v", err)
	}

	// Criar um handler de teste
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Criar um request de teste
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Executar o handler com tracing
	handler := TracedHandler(testHandler)
	handler.ServeHTTP(w, req)

	// Verificar o status code
	if w.Code != http.StatusOK {
		t.Errorf("Status code esperado %d, obtido %d", http.StatusOK, w.Code)
	}
}

func TestTraceSpan(t *testing.T) {
	// Inicializar o tracer
	err := InitTracer("test-service")
	if err != nil {
		t.Fatalf("Falha ao inicializar tracer: %v", err)
	}

	// Teste com função que retorna erro
	err = TraceSpan(context.Background(), "test-error", func(ctx context.Context) error {
		return fmt.Errorf("erro de teste")
	})
	if err == nil {
		t.Error("Esperava erro, obteve nil")
	}

	// Teste com função que não retorna erro
	err = TraceSpan(context.Background(), "test-success", func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		t.Errorf("Não esperava erro, obteve: %v", err)
	}
}

func TestTraceSpanWithAttributes(t *testing.T) {
	// Inicializar o tracer
	err := InitTracer("test-service")
	if err != nil {
		t.Fatalf("Falha ao inicializar tracer: %v", err)
	}

	// Teste com atributos e função que retorna erro
	attributes := map[string]string{
		"test_key": "test_value",
	}
	err = TraceSpanWithAttributes(context.Background(), "test-error", attributes, func(ctx context.Context) error {
		return fmt.Errorf("erro de teste")
	})
	if err == nil {
		t.Error("Esperava erro, obteve nil")
	}

	// Teste com atributos e função que não retorna erro
	err = TraceSpanWithAttributes(context.Background(), "test-success", attributes, func(ctx context.Context) error {
		return nil
	})
	if err != nil {
		t.Errorf("Não esperava erro, obteve: %v", err)
	}
}

func TestResponseWriter(t *testing.T) {
	// Criar um ResponseWriter de teste
	w := httptest.NewRecorder()
	rw := newResponseWriter(w)

	// Testar WriteHeader
	statusCode := http.StatusNotFound
	rw.WriteHeader(statusCode)
	if rw.statusCode != statusCode {
		t.Errorf("Status code esperado %d, obtido %d", statusCode, rw.statusCode)
	}

	// Testar Write
	testData := []byte("test data")
	n, err := rw.Write(testData)
	if err != nil {
		t.Errorf("Erro ao escrever: %v", err)
	}
	if n != len(testData) {
		t.Errorf("Bytes escritos esperados %d, obtidos %d", len(testData), n)
	}
}
