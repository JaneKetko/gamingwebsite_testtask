package mysql

import (
	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
)

// PlayerModel represents Player model for MongoDB.
type PlayerModel struct {
	ID      int     `db:"id"`
	Name    string  `db:"name"`
	Balance float32 `db:"balance"`
}

// NewPlayerModel creates new playerModel from player.Player.
func NewPlayerModel(pl player.Player) PlayerModel {
	return PlayerModel{Name: pl.Name, Balance: pl.Balance}
}

// ToPlayer converts PlayerModel to player.Player.
func (pm PlayerModel) ToPlayer() *player.Player {
	return &player.Player{ID: pm.ID, Name: pm.Name, Balance: pm.Balance}
}
