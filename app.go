package main

import (
	"log"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/database/mongo"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/server"
)

func main() {

	var session mongo.Session
	// TODO you can process this error in init function of the package.
	err := session.Open()
	if err != nil {
		log.Fatalf("cannot create mongo session: %v", err)
	}
	players := session.Players()
	mngr := manager.Manager{DB: &players}
	s := server.NewServer(mngr)
	s.Start()
}
