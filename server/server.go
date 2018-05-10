package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//PORT  is default port.
const PORT = ":8080"

//Server is router
type Server struct {
	router *mux.Router
}

//NewServer create new Server instance.
func NewServer() *Server {
	return &Server{mux.NewRouter()}
}

//Start  start server with PORT.
func (s *Server) Start() {
	err := http.ListenAndServe(PORT, s.router)
	if err != nil {
		log.Fatal(err)
	}
}
