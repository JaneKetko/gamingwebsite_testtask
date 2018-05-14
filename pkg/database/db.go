package database

import (
	"errors"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/player"
)

var (
	//ErrWrongID is error for wrong id.
	ErrWrongID = errors.New("wrong ID")
)

//DB is interface for database.
type DB interface {
	// TODO it is better return pointers, not object
	// TODO it is better don't use verbs like get, set, put and etc., if we can (go way)
	GetPlayerByID(id int) (player.Player, error)
	AddPlayer(name string) (int, error)
	DeletePlayer(id int) error
	UpdatePlayer(id int, player player.Player) error
}
