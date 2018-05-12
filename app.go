package main

import (
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/database"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/server"

)

func main() {

	aDB := new(database.ArrayDB)
	mngr := manager.Manager{DB: aDB}
	s := server.NewServer(mngr)
	s.Start()
}
