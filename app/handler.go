package app

import (
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	repository Repository
}

func (h *Handler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := h.repository.FindAllTodos()
	if err != nil {
		log.Println("Error finding all todos", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(todos)
	if err != nil {
		log.Println("Error marshalling todos", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
