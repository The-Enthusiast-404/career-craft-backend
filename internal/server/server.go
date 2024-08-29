package server

import (
	"net/http"

	"dev.theenthusiast.career-craft/internal/api/handlers"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

type Server struct {
	Router *httprouter.Router
	DB     *sqlx.DB
}

func NewServer(db *sqlx.DB) *Server {
	s := &Server{
		Router: httprouter.New(),
		DB:     db,
	}
	s.routes()
	return s
}

func (s *Server) routes() {
	jh := handlers.NewJobHandler(s.DB)
	s.Router.GET("/jobs/:company", jh.GetJobsByCompany)
	s.Router.POST("/jobs", jh.CreateJob)
	s.Router.POST("/jobs/bulk", jh.BulkCreateJobs)
}

// CORSMiddleware wraps the router with CORS functionality
func (s *Server) CORSMiddleware() http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"}, // Add your frontend URL
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		Debug:          true, // Enable debugging for testing, disable in production
	})

	return c.Handler(s.Router)
}

// Run starts the server with CORS middleware
func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.CORSMiddleware())
}
