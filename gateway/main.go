package main

import (
	"log"
	"net/http"
)

const (
	httpAddr = ":3000"
)

func main() {

	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting the server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to Start the server")
	}

}
