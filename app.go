package main

import (
	"log"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/database"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/database/mongo"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/server"
)

// ManagerWithArrayDB creates manager with arrayDB as database.
func ManagerWithArrayDB() manager.Manager {
	return manager.Manager{DB: new(database.ArrayDB)}
}

func main() {

	var session mongo.Session
	err := session.Open()
	defer session.Close()
	if err != nil {
		log.Fatalf("cannot create mongo session: %v", err)
	}
	players := session.Players()
	mngr := manager.Manager{DB: &players}
	s := server.NewServer(mngr)
	s.Start()
}
