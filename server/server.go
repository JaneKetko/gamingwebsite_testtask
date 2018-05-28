package server

import (
	"log"
	"net/http"

	"github.com/Ragnar-BY/gamingwebsite_testtask/manager"
	"github.com/gorilla/mux"
)

// Server is router.
type Server struct {
	router *mux.Router
}

// NewServer creates new Server instance.
func NewServer(mngr manager.Manager) *Server {
	s := &Server{mux.NewRouter()}
	s.router = newManagerRouter(mngr, s.router)
	return s
}

// Start starts pkg with addr.
func (s *Server) Start(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.router))
}
