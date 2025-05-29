package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Notification struct {
	UserID  string `json:"user_id"`
	Message string `json:"message"`
}

var notifications = []Notification{}

func notificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	var n Notification
	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}
	notifications = append(notifications, n)
	w.WriteHeader(http.StatusNoContent)
}

func notificationsByUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}
	userID := strings.TrimPrefix(r.URL.Path, "/notifications/")
	var userNotifications []Notification
	for _, n := range notifications {
		if n.UserID == userID {
			userNotifications = append(userNotifications, n)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userNotifications)
}

func main() {
	http.HandleFunc("/notifications", notificationHandler)
	http.HandleFunc("/notifications/", notificationsByUserHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	fmt.Println("Notification Service rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
