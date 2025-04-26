package handlers

import (
	"fmt"
	"net/http"
)

type Handler struct{}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.URL.Path {

	case "/":
		h.home(w, r)

	default:
		http.Error(w, "Invalid path", http.StatusNotFound)
	}
}

func (h *Handler) home(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "Hey there! This is the home page!\n")
}
