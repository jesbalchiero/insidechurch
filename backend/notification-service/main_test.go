package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotificationHandler(t *testing.T) {
	t.Run("POST /notifications deve enviar notificação", func(t *testing.T) {
		reqBody := Notification{
			UserID:  "1",
			Message: "Bem-vindo ao Inside Church!",
		}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/notifications", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		notificationHandler(rec, req)
		if rec.Code != http.StatusNoContent {
			t.Errorf("Esperava status 204, recebeu %d", rec.Code)
		}
	})

	t.Run("POST /notifications com JSON inválido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/notifications", strings.NewReader("json inválido"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		notificationHandler(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Esperava status 400, recebeu %d", rec.Code)
		}
	})

	t.Run("Método não permitido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/notifications", nil)
		rec := httptest.NewRecorder()
		notificationHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Esperava status 405, recebeu %d", rec.Code)
		}
	})
}

func TestNotificationsByUserHandler(t *testing.T) {
	t.Run("GET /notifications/1 deve retornar notificações do usuário", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/notifications/1", nil)
		rec := httptest.NewRecorder()
		notificationsByUserHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var notifications []Notification
		json.NewDecoder(rec.Body).Decode(&notifications)
		if len(notifications) == 0 {
			t.Error("Lista de notificações vazia")
		}
	})

	t.Run("GET /notifications/999 deve retornar lista vazia", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/notifications/999", nil)
		rec := httptest.NewRecorder()
		notificationsByUserHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var notifications []Notification
		json.NewDecoder(rec.Body).Decode(&notifications)
		if len(notifications) != 0 {
			t.Error("Lista de notificações não deveria estar vazia")
		}
	})

	t.Run("Método não permitido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/notifications/1", nil)
		rec := httptest.NewRecorder()
		notificationsByUserHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Esperava status 405, recebeu %d", rec.Code)
		}
	})
}
