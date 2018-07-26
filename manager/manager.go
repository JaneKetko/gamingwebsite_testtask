package manager

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"

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
	PlayerByID(ctx context.Context, id int) (*player.Player, error)
	AddPlayer(ctx context.Context, name string) (int, error)
	DeletePlayer(ctx context.Context, id int) error
	UpdatePlayer(ctx context.Context, id int, player player.Player) error
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
	mutePlayers map[int]*sync.Mutex
	PlayerDB    PlayerDB
	muteTours   map[int]*sync.Mutex
	TourDB      TournamentDB
	random      *rand.Rand
}

// NewManager is new manager
func NewManager(players PlayerDB, tours TournamentDB) Manager {
	return Manager{
		mutePlayers: make(map[int]*sync.Mutex),
		PlayerDB:    players,
		muteTours:   make(map[int]*sync.Mutex),
		TourDB:      tours,
		random:      rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// CreateNewPlayer creates new player in PlayerDB.
func (m *Manager) CreateNewPlayer(ctx context.Context, name string) (int, error) {
	id, err := m.PlayerDB.AddPlayer(ctx, name)
	if err != nil {
		return 0, err
	}
	m.createPlayerMutexIfNotExist(id)
	return id, nil
}

// GetPlayerPoints gets player points.
func (m *Manager) GetPlayerPoints(ctx context.Context, playerID int) (float32, error) {
	m.createPlayerMutexIfNotExist(playerID)
	m.mutePlayers[playerID].Lock()
	defer m.mutePlayers[playerID].Unlock()
	pl, err := m.PlayerDB.PlayerByID(ctx, playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	return pl.Balance, nil
}

// TakePointsFromPlayer takes points from player.
func (m *Manager) TakePointsFromPlayer(ctx context.Context, playerID int, points float32) (float32, error) {
	m.createPlayerMutexIfNotExist(playerID)
	m.mutePlayers[playerID].Lock()
	defer m.mutePlayers[playerID].Unlock()
	pl, err := m.PlayerDB.PlayerByID(ctx, playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	if pl.Balance < points {
		return 0, ErrNotEnoughBalance
	}
	pl.Balance -= points
	return pl.Balance, m.PlayerDB.UpdatePlayer(ctx, playerID, *pl)
}

// FundPointsToPlayer funds points to player.
func (m *Manager) FundPointsToPlayer(ctx context.Context, playerID int, points float32) (float32, error) {
	m.createPlayerMutexIfNotExist(playerID)
	m.mutePlayers[playerID].Lock()
	defer m.mutePlayers[playerID].Unlock()
	pl, err := m.PlayerDB.PlayerByID(ctx, playerID)
	if err != nil {
		return 0, fmt.Errorf("cannot get player ID: %v", err)
	}
	pl.Balance += points
	return pl.Balance, m.PlayerDB.UpdatePlayer(ctx, playerID, *pl)
}

// RemovePlayer removes player.
func (m *Manager) RemovePlayer(ctx context.Context, playerID int) error {
	//TODO: if we remove player, should we remove mutex from map, and how if should
	m.createPlayerMutexIfNotExist(playerID)
	m.mutePlayers[playerID].Lock()
	defer m.mutePlayers[playerID].Unlock()
	err := m.PlayerDB.DeletePlayer(ctx, playerID)
	if err != nil {
		return fmt.Errorf("cannot delete player with ID %v: %v", playerID, err)
	}
	return nil
}

func (m *Manager) createPlayerMutexIfNotExist(playerID int) {
	if _, ok := m.mutePlayers[playerID]; !ok {
		m.mutePlayers[playerID] = &sync.Mutex{}
	}
}

// AnnounceTournament creates tournament in tournamentDB.
func (m *Manager) AnnounceTournament(deposit float32) (int, error) {
	if deposit <= 0 {
		return 0, errors.New("cannot announce tournament: deposit is negative")
	}
	id, err := m.TourDB.CreateTournament(deposit)
	if err != nil {
		return 0, fmt.Errorf("cannot announce tournament: %v", err)
	}
	m.createToursMutexIfNotExist(id)
	return id, nil
}

// JoinTournament add player to tournament if possible.
// Player add tour.Deposit to tour.Fund.
func (m *Manager) JoinTournament(tourID int, playerID int) error {
	tour, err := m.TourDB.TournamentByID(tourID)
	if err != nil {
		return fmt.Errorf("cannot join to tournament: %v", err)
	}
	if tour.IsFinished {
		return errors.New("cannot join to tournament: tournament is finished")
	}

	_, err = m.TakePointsFromPlayer(playerID, tour.Deposit)
	if err != nil {
		return fmt.Errorf("cannot join to tournament: %v", err)
	}
	m.createToursMutexIfNotExist(tourID)
	m.muteTours[tourID].Lock()
	defer m.muteTours[tourID].Unlock()

	if tour.Participants == nil {
		tour.Participants = make([]int, 0)
	}
	tour.Participants = append(tour.Participants, playerID)
	tour.Fund += tour.Deposit
	err = m.TourDB.UpdateTournament(tourID, tour)
	if err != nil {
		return fmt.Errorf("cannot join to tournament: %v", err)
	}
	return nil
}

// ResultTournament count tournament result, if it is unknown
// or return known results. Returns winner, prize.
func (m *Manager) ResultTournament(tourID int) (*player.Player, float32, error) {
	tour, err := m.TourDB.TournamentByID(tourID)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot get tournament results: %v", err)
	}
	if tour.IsFinished {
		if tour.Winner == nil {
			return nil, 0, errors.New("cannot get tournament results: tournament is finished, by winner is unknown")
		}
		winner, err := m.PlayerDB.PlayerByID(*tour.Winner)
		if err != nil {
			return nil, 0, fmt.Errorf("cannot get tournament results: %v", err)
		}
		return winner, tour.Fund, nil
	}
	count := len(tour.Participants)
	if count == 0 {
		return nil, 0, errors.New("cannot get tournament results: there are not participants")
	}
	winnerID := m.random.Intn(count)
	winner, err := m.PlayerDB.PlayerByID(winnerID)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot get tournament results: %v", err)
	}
	m.createToursMutexIfNotExist(tourID)
	m.muteTours[tourID].Lock()
	defer m.muteTours[tourID].Unlock()

	tour.Winner = &winnerID
	tour.IsFinished = true
	err = m.TourDB.UpdateTournament(tourID, tour)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot get tournament results: %v", err)
	}
	_, err = m.FundPointsToPlayer(winnerID, tour.Fund)
	if err != nil {
		return nil, 0, fmt.Errorf("cannot get tournament results: %v", err)
	}
	return winner, tour.Fund, nil
}
func (m *Manager) createToursMutexIfNotExist(tourID int) {
	if _, ok := m.muteTours[tourID]; !ok {
		m.muteTours[tourID] = &sync.Mutex{}
	}
}
