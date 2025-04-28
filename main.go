package main

import (
	"log"
	"net/http"

	"github.com/kazxieo/go-http-server/handlers"
	"github.com/kazxieo/go-http-server/middleware"
)

const Addr = ":8080" // Should I name this to port? Idk

// The middlewares are called in reverse order
var middlewares = []func(http.Handler) http.Handler{
	middleware.Authentication,
	middleware.Logging, // This will be called first
}

func main() {

	// Wrap the handler with middlewares
	var h http.Handler = &handlers.Handler{}
	for _, m := range middlewares {
		h = m(h)
	}

	server := http.Server{
		Addr:    Addr,
		Handler: h,
	}

	log.Fatalln(server.ListenAndServe())
}
