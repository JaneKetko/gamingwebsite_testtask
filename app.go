package main

import (
	"gamingwebsite_testtask/server"
)

func main() {
	s := server.NewServer()
	s.Start()
}
