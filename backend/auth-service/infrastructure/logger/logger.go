package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
	"time"
)

// Field representa um campo de log
type Field struct {
	Key   string
	Value interface{}
}

// LogEntry representa uma entrada de log
type LogEntry struct {
	Timestamp string                 `json:"timestamp"`
	Level     string                 `json:"level"`
	Message   string                 `json:"message"`
	Fields    map[string]interface{} `json:"fields,omitempty"`
	Error     string                 `json:"error,omitempty"`
}

// Logger é a interface para logging
type Logger interface {
	Info(msg string, fields ...Field)
	Error(msg string, err error, fields ...Field)
	Debug(msg string, fields ...Field)
}

// JSONLogger implementa Logger com saída em JSON
type JSONLogger struct {
	mu    sync.Mutex
	file  *os.File
	level string
}

// NewJSONLogger cria uma nova instância do JSONLogger
func NewJSONLogger(logPath string) (*JSONLogger, error) {
	// Criar diretório de logs se não existir
	if err := os.MkdirAll("logs", 0755); err != nil {
		return nil, fmt.Errorf("erro ao criar diretório de logs: %v", err)
	}

	// Abrir arquivo de log
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo de log: %v", err)
	}

	return &JSONLogger{
		file:  file,
		level: "info",
	}, nil
}

// Info registra uma mensagem de nível INFO
func (l *JSONLogger) Info(msg string, fields ...Field) {
	l.log("INFO", msg, nil, fields...)
}

// Error registra uma mensagem de nível ERROR
func (l *JSONLogger) Error(msg string, err error, fields ...Field) {
	l.log("ERROR", msg, err, fields...)
}

// Debug registra uma mensagem de nível DEBUG
func (l *JSONLogger) Debug(msg string, fields ...Field) {
	l.log("DEBUG", msg, nil, fields...)
}

// log registra uma entrada de log
func (l *JSONLogger) log(level, msg string, err error, fields ...Field) {
	entry := LogEntry{
		Timestamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   msg,
		Fields:    make(map[string]interface{}),
	}

	// Adicionar campos
	for _, field := range fields {
		entry.Fields[field.Key] = field.Value
	}

	// Adicionar erro se houver
	if err != nil {
		entry.Error = err.Error()
	}

	// Serializar para JSON
	l.mu.Lock()
	defer l.mu.Unlock()

	encoder := json.NewEncoder(l.file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(entry); err != nil {
		fmt.Printf("Erro ao escrever log: %v\n", err)
	}
}

// Close fecha o arquivo de log
func (l *JSONLogger) Close() error {
	return l.file.Close()
}

// Funções auxiliares para criar campos
func String(key, value string) Field {
	return Field{Key: key, Value: value}
}

func Int(key string, value int) Field {
	return Field{Key: key, Value: value}
}

func Float(key string, value float64) Field {
	return Field{Key: key, Value: value}
}

func Bool(key string, value bool) Field {
	return Field{Key: key, Value: value}
}

func Time(key string, value time.Time) Field {
	return Field{Key: key, Value: value.Format(time.RFC3339)}
}

func Error(err error) Field {
	return Field{Key: "error", Value: err.Error()}
}
