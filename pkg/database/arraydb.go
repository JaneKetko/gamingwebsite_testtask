package database

import (
	"gamingwebsite_testtask/pkg/player"
)

//ArrayDB is database from slice.
type ArrayDB struct {
	players []player.Player
}

//GetPlayerByID return player by ID.
func (a ArrayDB) GetPlayerByID(id int) (player.Player, error) {
	if id < 0 || id >= len(a.players) {
		return player.Player{}, ErrWrongID
	}
	return a.players[id], nil
}

//AddPlayer add player.
func (a *ArrayDB) AddPlayer(name string) (int, error) {
	id := len(a.players)
	pl := player.Player{ID: id, Name: name, Balance: 0}
	a.players = append(a.players, pl)
	return id, nil
}

//DeletePlayer delete player if possible.
func (a *ArrayDB) DeletePlayer(id int) error {
	if id < 0 || id >= len(a.players) {
		return ErrWrongID
	}
	a.players = append(a.players[:id], a.players[id+1:]...)
	return nil
}

//UpdatePlayer update player with id.
func (a *ArrayDB) UpdatePlayer(id int, player player.Player) error {
	if id < 0 || id >= len(a.players) {
		return ErrWrongID
	}
	player.ID = id
	a.players[id] = player
	return nil
}
