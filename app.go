package main

// TODO why do you use this path pkg?
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
	players, err := session.Players(opts.DBName, opts.PlayerCollection)
	if err != nil {
		log.Fatalf("Cannot get player collection: %v", err)
	}
	mngr := manager.Manager{DB: players}
	s := server.NewServer(mngr)
	s.Start(opts.ServerAddress)
}
