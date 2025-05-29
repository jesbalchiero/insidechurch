package metrics

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestMetricsMiddleware(t *testing.T) {
	// Criar um handler de teste
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Criar um request de teste
	req := httptest.NewRequest("GET", "/test", nil)
	w := httptest.NewRecorder()

	// Executar o handler com métricas
	handler := MetricsMiddleware(testHandler)
	handler.ServeHTTP(w, req)

	// Verificar o status code
	if w.Code != http.StatusOK {
		t.Errorf("Status code esperado %d, obtido %d", http.StatusOK, w.Code)
	}
}

func TestRecordHTTPDuration(t *testing.T) {
	// Testar registro de duração
	duration := 100 * time.Millisecond
	RecordHTTPDuration("/test", "GET", duration.Seconds())
}

func TestRecordHTTPRequest(t *testing.T) {
	// Testar registro de requisição
	RecordHTTPRequest("/test", "GET", http.StatusOK)
}

func TestRecordHTTPError(t *testing.T) {
	// Testar registro de erro
	RecordHTTPError("/test", "GET", http.StatusInternalServerError)
}

func TestRecordDatabaseOperation(t *testing.T) {
	// Testar registro de operação no banco de dados
	RecordDatabaseOperation("test_db", "SELECT", 100*time.Millisecond)
}

func TestRecordDatabaseError(t *testing.T) {
	// Testar registro de erro no banco de dados
	RecordDatabaseError("test_db", "SELECT")
}

func TestRecordCacheOperation(t *testing.T) {
	// Testar registro de operação no cache
	RecordCacheOperation("test_cache", "GET", 50*time.Millisecond)
}

func TestRecordCacheError(t *testing.T) {
	// Testar registro de erro no cache
	RecordCacheError("test_cache", "GET")
}

func TestRecordExternalAPICall(t *testing.T) {
	// Testar registro de chamada à API externa
	RecordExternalAPICall("test_api", "GET", 200*time.Millisecond)
}

func TestRecordExternalAPIError(t *testing.T) {
	// Testar registro de erro na chamada à API externa
	RecordExternalAPIError("test_api", "GET")
}

func TestRecordEventProcessing(t *testing.T) {
	// Testar registro de processamento de evento
	RecordEventProcessing("test_event", 150*time.Millisecond)
}

func TestRecordEventError(t *testing.T) {
	// Testar registro de erro no processamento de evento
	RecordEventError("test_event")
}

func TestRecordNotificationSent(t *testing.T) {
	// Testar registro de notificação enviada
	RecordNotificationSent("test_notification", "email")
}

func TestRecordNotificationError(t *testing.T) {
	// Testar registro de erro no envio de notificação
	RecordNotificationError("test_notification", "email")
}

func TestRecordUserAction(t *testing.T) {
	// Testar registro de ação do usuário
	RecordUserAction("test_user", "login")
}

func TestRecordUserError(t *testing.T) {
	// Testar registro de erro na ação do usuário
	RecordUserError("test_user", "login")
}
