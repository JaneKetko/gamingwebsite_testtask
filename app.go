package main

import (
	"log"

	"github.com/Ragnar-BY/gamingwebsite_testtask/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/mongo"
	"github.com/Ragnar-BY/gamingwebsite_testtask/server"
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
		log.Fatalf("Cannot start MongoDB on the %s: %v", opts.Address, err)
	}
	players, err := session.Players(opts.DBName, opts.PlayerCollection)
	if err != nil {
		log.Fatalf("Cannot get player collection: %v", err)
	}
	mngr := manager.NewManager(players)
	s := server.NewServer(mngr)
	s.Start(opts.ServerAddress)
}
