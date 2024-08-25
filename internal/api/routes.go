package api

import (
	"net/http"

	"dev.theenthusiast.career-craft/internal/api/handlers"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HelloHandler)
	return mux
}
