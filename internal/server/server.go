package server

import (
	"net/http"

	"dev.theenthusiast.career-craft/internal/api/handlers"
	"dev.theenthusiast.career-craft/internal/api/middleware"
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
	ch := handlers.NewCompanyHandler(s.DB)

	// Public routes
	s.Router.GET("/health", s.healthCheck)

	// Protected routes
	s.Router.GET("/jobs/:company", s.wrapWithAuth(jh.GetJobsByCompany))
	s.Router.POST("/jobs", s.wrapWithAuth(jh.CreateJob))
	s.Router.POST("/jobs/bulk", jh.BulkCreateJobs)
	s.Router.GET("/company/:company", s.wrapWithAuth(ch.GetCompanyDetails))
}

func (s *Server) healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *Server) wrapWithAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		middleware.AuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
			h(w, r, ps)
		})(w, r)
	}
}

// CORSMiddleware wraps the router with CORS functionality
func (s *Server) CORSMiddleware() http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Add your frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            true, // Enable debugging for testing, disable in production
	})
	return c.Handler(s.Router)
}

// Run starts the server with CORS middleware
func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.CORSMiddleware())
}
