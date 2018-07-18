package main

import (
	"log"

	"github.com/Ragnar-BY/gamingwebsite_testtask/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/mongo"
	"github.com/Ragnar-BY/gamingwebsite_testtask/mysql"
	"github.com/Ragnar-BY/gamingwebsite_testtask/server"
)

func startWithMongo(opts settings) {
	var session mongo.Session
	err := session.Open(opts.Address)
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

func startWithMySQL(opts settings) {
	sql, err := mysql.Open(opts.User, opts.Password, opts.DBName)
	if err != nil {
		log.Fatal("Cannot start MySQL: ", err)
	}
	ps := mysql.PlayerService{DB: sql, Name: opts.PlayerCollection}

	mngr := manager.NewManager(ps)
	s := server.NewServer(mngr)
	s.Start(opts.ServerAddress)

}
func main() {
	opts := new(settings)
	err := opts.Parse()
	if err != nil {
		log.Fatalf("Cannot parse settings: %v", err)
	}
	if opts.Type == "mysql" {
		startWithMySQL(*opts)
	} else {
		startWithMongo(*opts)
	}

}
