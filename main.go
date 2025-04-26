package main

import (
	"log"
	"net/http"

	"github.com/kazxieo/go-http-server/handlers"
)

const Addr = ":8080" // Should I name this to port? Idk

func main() {

	server := http.Server{
		Addr:    Addr,
		Handler: &handlers.Handler{},
	}

	log.Fatalln(server.ListenAndServe())
}
