package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Event struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var events = []Event{
	{ID: "1", Title: "Culto de Domingo"},
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(events)
	case http.MethodPost:
		var e Event
		if err := json.NewDecoder(r.Body).Decode(&e); err != nil {
			http.Error(w, "JSON inválido", http.StatusBadRequest)
			return
		}
		e.ID = fmt.Sprintf("%d", len(events)+1)
		events = append(events, e)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(e)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func eventByIDHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/events/")
	for _, e := range events {
		if e.ID == id {
			if r.Method == http.MethodGet {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(e)
				return
			} else if r.Method == http.MethodPut {
				var updated Event
				if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
					http.Error(w, "JSON inválido", http.StatusBadRequest)
					return
				}
				e.Title = updated.Title
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(e)
				return
			} else if r.Method == http.MethodDelete {
				for i, event := range events {
					if event.ID == id {
						events = append(events[:i], events[i+1:]...)
						w.WriteHeader(http.StatusNoContent)
						return
					}
				}
			}
		}
	}
	http.Error(w, "Evento não encontrado", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/events", eventsHandler)
	http.HandleFunc("/events/", eventByIDHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ok")
	})
	fmt.Println("Event Service rodando na porta 8080")
	http.ListenAndServe(":8080", nil)
}
