package main

import (
	"log"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/mongo"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/server"
)

func main() {
	opts := new(settings)
	err := opts.Parse()
	if err != nil {
		log.Fatalf("Cannot parse settings: %v", err)
	}
	var session mongo.Session
	err = session.Open(opts.Address)
	if err != nil {
		log.Fatalf("Cannot start MongoDB: %v", err)
	}
	players := session.Players(opts.DBName, opts.PlayerCollection)
	mngr := manager.Manager{DB: &players}
	s := server.NewServer(mngr)
	s.Start(opts.ServerAddress)
}
