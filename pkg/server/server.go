package server

import (
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//PORT  is default port.
// TODO port is a port :) without ":"
const PORT = ":8080"

//Server is router
type Server struct {
	router *mux.Router
}

//NewServer create new Server instance.
func NewServer(mngr manager.Manager) *Server {
	s := &Server{mux.NewRouter()}
	s.router = NewManagerRouter(mngr, s.router)
	return s
}

//Start  start pkg with PORT.
func (s *Server) Start() {
	// TODO fmt.Sprintf(":%s", PORT)
	// it is better call like log.Fatal(http.ListenAndServe(PORT, s.router))
	err := http.ListenAndServe(PORT, s.router)
	if err != nil {
		log.Fatal(err)
	}
}
