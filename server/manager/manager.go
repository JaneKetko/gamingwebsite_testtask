package manager

import (
	"errors"
	"gamingwebsite_testtask/server/database"
)

//Manager manage players
type Manager struct {
	db database.DB
}

//CreateNewPlayer create new player in DB
func (m *Manager) CreateNewPlayer(name string) (int, error) {
	return m.db.AddPlayer(name)
}

//GetPlayerPoints get player points
func (m *Manager) GetPlayerPoints(playerID int) (int, error) {
	player, err := m.db.GetPlayerByID(playerID)
	if err != nil {
		return 0, err
	}
	return player.Balance, nil
}

//TakePointsFromPlayer take points from player
func (m *Manager) TakePointsFromPlayer(playerID int, points int) (int, error) {

	player, err := m.db.GetPlayerByID(playerID)
	if err != nil {
		return 0, err
	}
	if player.Balance < points {
		return 0, errors.New("player has not enough balance")
	}
	balance := player.Balance - points
	player.Balance = balance

	err = m.db.UpdatePlayer(playerID, player)
	return balance, err
}

//FundPointsToPlayer fund points to player
func (m *Manager) FundPointsToPlayer(playerID int, points int) (int, error) {
	player, err := m.db.GetPlayerByID(playerID)
	if err != nil {
		return 0, err
	}
	balance := player.Balance + points
	player.Balance = balance

	err = m.db.UpdatePlayer(playerID, player)
	return balance, err
}
