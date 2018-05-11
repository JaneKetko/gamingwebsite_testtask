package database

import (
	"errors"
	"gamingwebsite_testtask/pkg/player"
)

var (
	//ErrWrongID is error for wrong id.
	ErrWrongID = errors.New("wrong ID")
)

//DB is interface for database.
type DB interface {
	GetPlayerByID(id int) (player.Player, error)
	AddPlayer(name string) (int, error)
	DeletePlayer(id int) error
	UpdatePlayer(id int, player player.Player) error
}
