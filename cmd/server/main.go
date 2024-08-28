package main

import (
	"log"
	"net/http"

	"dev.theenthusiast.career-craft/internal/database"
	"dev.theenthusiast.career-craft/internal/server"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}
	defer db.Close()

	s := server.NewServer(db)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
