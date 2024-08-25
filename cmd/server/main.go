package main

import (
	"log"
	"net/http"

	"dev.theenthusiast.career-craft/internal/api"
)

func main() {
	router := api.SetupRoutes()

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
