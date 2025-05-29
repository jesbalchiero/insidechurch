package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{ID: "1", Name: "João", Email: "joao@email.com"},
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		u.ID = fmt.Sprintf("%d", len(users)+1)
		users = append(users, u)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(u)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func userByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	for i, u := range users {
		if u.ID == id {
			switch r.Method {
			case http.MethodGet:
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(u)
				return
			case http.MethodPut:
				var updated User
				if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				users[i].Name = updated.Name
				users[i].Email = updated.Email
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(users[i])
				return
			case http.MethodDelete:
				users = append(users[:i], users[i+1:]...)
				w.WriteHeader(http.StatusNoContent)
				return
			default:
				http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
				return
			}
		}
	}
	http.Error(w, "Usuário não encontrado", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/users/", userByIDHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	fmt.Println("User Service rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
