package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/gorilla/mux"
)

// TODO set this port via configuration.
// PORT is default port.
const PORT = "8080"

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

// Start  starts pkg with PORT.
func (s *Server) Start() {
	err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), s.router)
	if err != nil {
		log.Fatal(err)
	}
}
