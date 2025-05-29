package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUsersHandler(t *testing.T) {
	t.Run("GET /users deve retornar lista de usuários", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		rec := httptest.NewRecorder()
		usersHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var users []User
		json.NewDecoder(rec.Body).Decode(&users)
		if len(users) == 0 {
			t.Error("Lista de usuários vazia")
		}
	})

	t.Run("POST /users deve criar novo usuário", func(t *testing.T) {
		reqBody := User{
			Name:  "Maria",
			Email: "maria@email.com",
		}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		usersHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var newUser User
		json.NewDecoder(rec.Body).Decode(&newUser)
		if newUser.ID == "" {
			t.Error("ID não foi gerado")
		}
	})

	t.Run("POST /users com JSON inválido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader("json inválido"))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		usersHandler(rec, req)
		if rec.Code != http.StatusBadRequest {
			t.Errorf("Esperava status 400, recebeu %d", rec.Code)
		}
	})

	t.Run("Método não permitido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPut, "/users", nil)
		rec := httptest.NewRecorder()
		usersHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Esperava status 405, recebeu %d", rec.Code)
		}
	})
}

func TestUserByIDHandler(t *testing.T) {
	t.Run("GET /users/1 deve retornar usuário existente", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/1", nil)
		rec := httptest.NewRecorder()
		userByIDHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var user User
		json.NewDecoder(rec.Body).Decode(&user)
		if user.ID != "1" {
			t.Error("ID incorreto")
		}
	})

	t.Run("GET /users/999 deve retornar 404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/users/999", nil)
		rec := httptest.NewRecorder()
		userByIDHandler(rec, req)
		if rec.Code != http.StatusNotFound {
			t.Errorf("Esperava status 404, recebeu %d", rec.Code)
		}
	})

	t.Run("PUT /users/1 deve atualizar usuário", func(t *testing.T) {
		reqBody := User{
			Name:  "João Atualizado",
			Email: "joao.novo@email.com",
		}
		jsonBody, _ := json.Marshal(reqBody)
		req := httptest.NewRequest(http.MethodPut, "/users/1", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		userByIDHandler(rec, req)
		if rec.Code != http.StatusOK {
			t.Errorf("Esperava status 200, recebeu %d", rec.Code)
		}
		var updatedUser User
		json.NewDecoder(rec.Body).Decode(&updatedUser)
		if updatedUser.Name != "João Atualizado" {
			t.Error("Nome não foi atualizado")
		}
	})

	t.Run("DELETE /users/1 deve remover usuário", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodDelete, "/users/1", nil)
		rec := httptest.NewRecorder()
		userByIDHandler(rec, req)
		if rec.Code != http.StatusNoContent {
			t.Errorf("Esperava status 204, recebeu %d", rec.Code)
		}
	})

	t.Run("Método não permitido deve retornar erro", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodPost, "/users/1", nil)
		rec := httptest.NewRecorder()
		userByIDHandler(rec, req)
		if rec.Code != http.StatusMethodNotAllowed {
			t.Errorf("Esperava status 405, recebeu %d", rec.Code)
		}
	})
}
