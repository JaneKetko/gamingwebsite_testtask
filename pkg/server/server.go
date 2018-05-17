package server

import (
	"log"
	"net/http"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/gorilla/mux"
)

// Server is router.
type Server struct {
	router *mux.Router
}

// NewServer creates new Server instance.
func NewServer(mngr manager.Manager) *Server {
	s := &Server{mux.NewRouter()}
	s.router = NewManagerRouter(mngr, s.router)
	return s
}

// Start starts pkg with addr.
func (s *Server) Start(addr string) {
	// TODO log.Fatal(http.ListenAndServe(addr, s.router)) - is the same and more popular.
	err := http.ListenAndServe(addr, s.router)
	if err != nil {
		log.Fatal(err)
	}
}
