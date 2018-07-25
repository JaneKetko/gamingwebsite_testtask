package manager

import (
	"errors"
	"fmt"
	"sync"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
	"github.com/Ragnar-BY/gamingwebsite_testtask/tournament"
)

//go:generate mockery -name=PlayerDB -name=TournamentDB -inpkg
var (
	// ErrNotEnoughBalance is error for not enough balance.
	ErrNotEnoughBalance = errors.New("player has not enough balance")
)

// PlayerDB is interface for database.
type PlayerDB interface {
	PlayerByID(id int) (*player.Player, error)
	AddPlayer(name string) (int, error)
	DeletePlayer(id int) error
	UpdatePlayer(id int, player player.Player) error
}

// TournamentDB is interface for tournament database.
type TournamentDB interface {
	TournamentByID(id int) (tournament.Tournament, error)
	CreateTournament(deposit float32) (int, error)
	DeleteTournament(id int) error
	UpdateTournament(id int, t tournament.Tournament) error
}

// Manager manages players and tournaments.
type Manager struct {
	mute     map[int]*sync.Mutex
	PlayerDB PlayerDB
	TourDB   TournamentDB
}

// NewManager is new manager
func NewManager(players PlayerDB, tours TournamentDB) Manager {
	return Manager{mute: make(map[int]*sync.Mutex), PlayerDB: players, TourDB: tours}
}

// CreateNewPlayer creates new player in PlayerDB.
func (m *Manager) CreateNewPlayer(name string) (int, error) {
	id, err := m.PlayerDB.AddPlayer(name)
	if err != nil {
		return 0, err
	}
	m.createMutexIfNotExist(id)
	return id, nil
}

// GetPlayerPoints gets player points.
func (m *Manager) GetPlayerPoints(playerID int) (float32, error) {
	m.createMutexIfNotExist(playerID)
	m.mute[playerID].Lock()
	defer m.mute[playerID].Unlock()
	pl, err := m.PlayerDB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	return pl.Balance, nil
}

// TakePointsFromPlayer takes points from player.
func (m *Manager) TakePointsFromPlayer(playerID int, points float32) (float32, error) {
	m.createMutexIfNotExist(playerID)
	m.mute[playerID].Lock()
	defer m.mute[playerID].Unlock()
	pl, err := m.PlayerDB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	if pl.Balance < points {
		return 0, ErrNotEnoughBalance
	}
	pl.Balance -= points
	return pl.Balance, m.PlayerDB.UpdatePlayer(playerID, *pl)
}

// FundPointsToPlayer funds points to player.
func (m *Manager) FundPointsToPlayer(playerID int, points float32) (float32, error) {
	m.createMutexIfNotExist(playerID)
	m.mute[playerID].Lock()
	defer m.mute[playerID].Unlock()
	pl, err := m.PlayerDB.PlayerByID(playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	pl.Balance += points
	return pl.Balance, m.PlayerDB.UpdatePlayer(playerID, *pl)
}

// RemovePlayer removes player.
func (m *Manager) RemovePlayer(playerID int) error {
	//TODO: if we remove player, should we remove mutex from map, and how if should
	m.createMutexIfNotExist(playerID)
	m.mute[playerID].Lock()
	defer m.mute[playerID].Unlock()
	err := m.PlayerDB.DeletePlayer(playerID)
	if err != nil {
		return fmt.Errorf("cannot delete player with ID %v: %v", playerID, err)
	}
	return nil
}

func (m *Manager) createMutexIfNotExist(playerID int) {
	if _, ok := m.mute[playerID]; !ok {
		m.mute[playerID] = &sync.Mutex{}
	}
}
