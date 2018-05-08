package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

//PORT  default port
const PORT = ":8080"

type Server struct {
	router *mux.Router
}

//NewServer create new Server instance
func NewServer() *Server {
	return &Server{mux.NewRouter()}
}

//Start  start server with PORT
func (s *Server) Start() {
	http.ListenAndServe(PORT, s.router)
}
