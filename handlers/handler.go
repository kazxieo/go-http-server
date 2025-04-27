package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/kazxieo/go-http-server/database"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {

	case "/":
		h.home(w, r)

	case "/users":
		h.getUsers(w, r)

	default:
		http.Error(w, "Invalid path", http.StatusNotFound)
	}
}

func (h *Handler) home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hey there! This is the home page!\n")
}

func (h *Handler) getUsers(w http.ResponseWriter, _ *http.Request) {

	if err := json.NewEncoder(w).Encode(database.SampleDatabase); err != nil {
		http.Error(w, fmt.Sprintf("Error processing JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
