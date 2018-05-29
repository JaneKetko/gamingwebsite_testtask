package manager

import (
	"errors"
	"fmt"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
)

//go:generate mockery -name=DB -inpkg
var (
	// ErrNotEnoughBalance is error for not enough balance.
	ErrNotEnoughBalance = errors.New("player has not enough balance")
)

// DB is interface for database.
type DB interface {
	PlayerByID(id int) (*player.Player, error)
	AddPlayer(name string) (int, error)
	DeletePlayer(id int) error
	UpdatePlayerBalance(id int, player player.Player) error
}

// Manager manages players.
type Manager struct {
	DB DB
}

// CreateNewPlayer creates new player in DB.
func (m *Manager) CreateNewPlayer(name string) (int, error) {
	return m.DB.AddPlayer(name)
}

// GetPlayerPoints gets player points.
func (m *Manager) GetPlayerPoints(playerID int) (float32, error) {
	pl, err := m.DB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	return pl.Balance, nil
}

// TakePointsFromPlayer takes points from player.
func (m *Manager) TakePointsFromPlayer(playerID int, points float32) (float32, error) {
	pl, err := m.DB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	if pl.Balance < points {
		return 0, ErrNotEnoughBalance
	}
	pl.Balance -= points
	return pl.Balance, m.DB.UpdatePlayerBalance(playerID, *pl)
}

// FundPointsToPlayer funds points to player.
func (m *Manager) FundPointsToPlayer(playerID int, points float32) (float32, error) {
	pl, err := m.DB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	pl.Balance += points
	return pl.Balance, m.DB.UpdatePlayerBalance(playerID, *pl)
}
