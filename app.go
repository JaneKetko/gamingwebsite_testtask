package main

import (
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/mongo"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/server"
)

func main() {

	players := mongo.Players()
	mngr := manager.Manager{DB: &players}
	s := server.NewServer(mngr)
	s.Start()
}
