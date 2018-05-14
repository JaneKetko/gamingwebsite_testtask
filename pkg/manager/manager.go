package manager

// TODO after all standard libs you should add new line. (go way)
import (
	"errors"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/database"
)

//Manager manage players.
type Manager struct {
	DB database.DB
}

//CreateNewPlayer create new player in DB.
func (m *Manager) CreateNewPlayer(name string) (int, error) {
	return m.DB.AddPlayer(name)
}

//GetPlayerPoints get player points.
func (m *Manager) GetPlayerPoints(playerID int) (int, error) {
	player, err := m.DB.GetPlayerByID(playerID)
	if err != nil {
		// TODO it is better add error wrapping.
		return 0, err
	}
	return player.Balance, nil
}

//TakePointsFromPlayer take points from player.
func (m *Manager) TakePointsFromPlayer(playerID int, points int) (int, error) {

	player, err := m.DB.GetPlayerByID(playerID)
	if err != nil {
		return 0, err
	}
	if player.Balance < points {
		// TODO it is better move all custom error like global variable.
		return 0, errors.New("player has not enough balance")
	}
	// TODO why not player.Balance -= points?
	balance := player.Balance - points
	player.Balance = balance

	// TODO why not "return balance, m.DB.UpdatePlayer(playerID, player)"?
	err = m.DB.UpdatePlayer(playerID, player)
	return balance, err
}

//FundPointsToPlayer fund points to player.
func (m *Manager) FundPointsToPlayer(playerID int, points int) (int, error) {
	player, err := m.DB.GetPlayerByID(playerID)
	if err != nil {
		return 0, err
	}
	// TODO -//-.
	balance := player.Balance + points
	player.Balance = balance

	// TODO -//-.
	err = m.DB.UpdatePlayer(playerID, player)
	return balance, err
}
