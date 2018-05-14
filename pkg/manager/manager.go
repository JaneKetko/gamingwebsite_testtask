package manager

import (
	"errors"
	"fmt"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/database"
)

var (
	// ErrNotEnoughBalance is error for not enough balance.
	ErrNotEnoughBalance = errors.New("player has not enough balance")
)

// Manager manages players.
type Manager struct {
	DB database.DB
}

// CreateNewPlayer creates new player in DB.
func (m *Manager) CreateNewPlayer(name string) (int, error) {
	return m.DB.AddPlayer(name)
}

// GetPlayerPoints gets player points.
func (m *Manager) GetPlayerPoints(playerID int) (float32, error) {
	player, err := m.DB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	return player.Balance, nil
}

// TakePointsFromPlayer takes points from player.
func (m *Manager) TakePointsFromPlayer(playerID int, points float32) (float32, error) {

	player, err := m.DB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	if player.Balance < points {
		return 0, ErrNotEnoughBalance
	}
	player.Balance -= points
	return player.Balance, m.DB.UpdatePlayer(playerID, *player)
}

// FundPointsToPlayer funds points to player.
func (m *Manager) FundPointsToPlayer(playerID int, points float32) (float32, error) {
	player, err := m.DB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	player.Balance += points
	return player.Balance, m.DB.UpdatePlayer(playerID, *player)
}
