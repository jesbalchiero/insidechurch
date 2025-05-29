package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEventsHandler(t *testing.T) {
	t.Run("GET /events deve retornar lista de eventos", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/events", nil)
		rec := httptest.NewRecorder()
		eventsHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var events []Event
		json.NewDecoder(rec.Body).Decode(&events)
		if len(events) == 0 {
			t.Error("Lista de eventos vazia")
		}
	})

	t.Run("POST /events deve criar novo evento", func(t *testing.T) {
		reqBody := Event{
			Title: "Culto de Domingo",
		}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/events", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		eventsHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var newEvent Event
		json.NewDecoder(rec.Body).Decode(&newEvent)
		if newEvent.ID == "" {
			t.Error("ID não foi gerado")
		}
	})

	t.Run("POST /events com JSON inválido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/events", strings.NewReader("json inválido"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		eventsHandler(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Esperava status 400, recebeu %d", rec.Code)
		}
	})

	t.Run("Método não permitido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/events", nil)
		rec := httptest.NewRecorder()
		eventsHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Esperava status 405, recebeu %d", rec.Code)
		}
	})
}

func TestEventByIDHandler(t *testing.T) {
	t.Run("GET /events/1 deve retornar evento existente", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/events/1", nil)
		rec := httptest.NewRecorder()
		eventByIDHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var event Event
		json.NewDecoder(rec.Body).Decode(&event)
		if event.ID != "1" {
			t.Error("ID incorreto")
		}
	})

	t.Run("GET /events/999 deve retornar 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/events/999", nil)
		rec := httptest.NewRecorder()
		eventByIDHandler(rec, req)
		if rec.Code != http.StatusNotFound {
			t.Errorf("Esperava status 404, recebeu %d", rec.Code)
		}
	})

	t.Run("PUT /events/1 deve atualizar evento", func(t *testing.T) {
		reqBody := Event{
			Title: "Culto de Domingo Atualizado",
		}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/events/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		eventByIDHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var updatedEvent Event
		json.NewDecoder(rec.Body).Decode(&updatedEvent)
		if updatedEvent.Title != "Culto de Domingo Atualizado" {
			t.Error("Título não foi atualizado")
		}
	})

	t.Run("DELETE /events/1 deve remover evento", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/events/1", nil)
		rec := httptest.NewRecorder()
		eventByIDHandler(rec, req)
		if rec.Code != http.StatusNoContent {
			t.Errorf("Esperava status 204, recebeu %d", rec.Code)
		}
	})

	t.Run("Método não permitido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/events/1", nil)
		rec := httptest.NewRecorder()
		eventByIDHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Esperava status 405, recebeu %d", rec.Code)
		}
	})
}
