package database

import "gamingwebsite_testtask/server/player"

//DB is interface for database.
type DB interface {
	GetPlayerByID(id int) (player.Player, error)
	AddPlayer(name string) (int, error)
	DeletePlayer(id int) error
	UpdatePlayer(id int, player player.Player) error
}
