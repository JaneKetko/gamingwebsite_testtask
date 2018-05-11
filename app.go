package main

import (
	"gamingwebsite_testtask/pkg/database"
	"gamingwebsite_testtask/pkg/manager"
	"gamingwebsite_testtask/pkg/server"
)

func main() {

	aDB := new(database.ArrayDB)
	mngr := manager.Manager{DB: aDB}
	s := server.NewServer(mngr)
	s.Start()
}
