package server

import (
	"dev.theenthusiast.career-craft/internal/api/handlers"
	"github.com/jmoiron/sqlx"
	"github.com/julienschmidt/httprouter"
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
