package main

import (
	"gamingwebsite_testtask/pkg/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
