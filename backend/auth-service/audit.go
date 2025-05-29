package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"
	"time"
)

// AuditLog representa um registro de auditoria
type AuditLog struct {
	UserID    string    `json:"user_id"`
	Action    string    `json:"action"`
	Resource  string    `json:"resource"`
	Timestamp time.Time `json:"timestamp"`
	IP        string    `json:"ip"`
	Details   string    `json:"details,omitempty"`
}

// AuditLogger gerencia os logs de auditoria
type AuditLogger struct {
	mu    sync.Mutex
	file  *os.File
	queue chan AuditLog
}

var (
	auditLogger *AuditLogger
)

// Inicializa o logger de auditoria
func initAuditLogger() error {
	// Criar diretório de logs se não existir
	if err := os.MkdirAll("logs", 0755); err != nil {
		return fmt.Errorf("erro ao criar diretório de logs: %v", err)
	}

	// Abrir arquivo de log
	file, err := os.OpenFile("logs/audit.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("erro ao abrir arquivo de log: %v", err)
	}

	auditLogger = &AuditLogger{
		file:  file,
		queue: make(chan AuditLog, 1000), // Buffer de 1000 logs
	}

	// Iniciar worker para processar logs
	go auditLogger.processLogs()

	return nil
}

// Processa os logs em background
func (l *AuditLogger) processLogs() {
	for log := range l.queue {
		l.mu.Lock()
		encoder := json.NewEncoder(l.file)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(log); err != nil {
			fmt.Printf("Erro ao escrever log: %v\n", err)
		}
		l.mu.Unlock()
	}
}

// Log registra uma nova entrada de auditoria
func (l *AuditLogger) Log(userID, action, resource, ip, details string) {
	log := AuditLog{
		UserID:    userID,
		Action:    action,
		Resource:  resource,
		Timestamp: time.Now(),
		IP:        ip,
		Details:   details,
	}

	select {
	case l.queue <- log:
		// Log enviado com sucesso
	default:
		// Buffer cheio, log perdido
		fmt.Printf("Buffer de logs cheio, log perdido: %+v\n", log)
	}
}

// Fecha o logger de auditoria
func (l *AuditLogger) Close() error {
	close(l.queue)
	return l.file.Close()
}

// Middleware para registrar logs de auditoria
func auditMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Obter o ID do usuário do contexto (se disponível)
		userID := "anonymous"
		if claims, ok := r.Context().Value(userIDKey).(string); ok {
			userID = claims
		}

		// Registrar o acesso
		auditLogger.Log(
			userID,
			r.Method,
			r.URL.Path,
			r.RemoteAddr,
			fmt.Sprintf("User-Agent: %s", r.UserAgent()),
		)

		next.ServeHTTP(w, r)
	}
}
