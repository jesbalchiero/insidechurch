package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLoginHandler(t *testing.T) {
	reqBody := LoginRequest{
		Email:    "joao@email.com",
		Password: "senha123",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	loginHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Esperava status 200, recebeu %d", rec.Code)
	}
	var resp LoginResponse
	json.NewDecoder(rec.Body).Decode(&resp)
	if resp.Token == "" {
		t.Error("Token não foi gerado")
	}
}

func TestRefreshHandler(t *testing.T) {
	reqBody := RefreshRequest{
		RefreshToken: "token-antigo",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/refresh", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	refreshHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Esperava status 200, recebeu %d", rec.Code)
	}
	var resp LoginResponse
	json.NewDecoder(rec.Body).Decode(&resp)
	if resp.Token == "" {
		t.Error("Novo token não foi gerado")
	}
}

func TestValidateHandler(t *testing.T) {
	reqBody := ValidateRequest{
		Token: "token-jwt",
	}
	jsonBody, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/auth/validate", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	validateHandler(rec, req)
	if rec.Code != http.StatusOK {
		t.Errorf("Esperava status 200, recebeu %d", rec.Code)
	}
	var resp ValidateResponse
	json.NewDecoder(rec.Body).Decode(&resp)
	if !resp.Valid {
		t.Error("Token deveria ser válido")
	}
}
