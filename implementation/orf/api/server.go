package api

import (
	"fmt"
	"net/http"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

// Server exposes AOS HTTP API.
type Server struct {
	System *repository.ObjectSystem
}

// NewServer creates API server.
func NewServer(
	system *repository.ObjectSystem,
) *Server {

	return &Server{

		System: system,
	}
}

// Router builds HTTP routes.
func (s *Server) Router() http.Handler {

	mux :=
		http.NewServeMux()

	// system

	mux.HandleFunc(
		"/system/status",
		s.handleSystemStatus,
	)

	mux.HandleFunc(
		"/system/metrics",
		s.handleSystemMetrics,
	)

	// objects

	mux.HandleFunc(
		"/objects",
		s.handleObjects,
	)

	mux.HandleFunc(
		"/objects/",
		s.handleObject,
	)

	// graph

	mux.HandleFunc(
		"/graph",
		s.handleGraph,
	)

	mux.HandleFunc(
		"/graph/",
		s.handleGraphQuery,
	)

	// lifecycle rules

	mux.HandleFunc(
		"/rules",
		s.HandleRules,
	)

	return mux
}

// Start starts HTTP server.
func (s *Server) Start(
	addr string,
) error {

	fmt.Println(
		"Adaptive Object System API Server",
	)

	fmt.Println(
		"listen:",
		addr,
	)

	return http.ListenAndServe(
		addr,
		s.Router(),
	)
}

// Handler returns http handler.
func (s *Server) Handler() http.Handler {

	return s.Router()
}
