package manager

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
	"github.com/Ragnar-BY/gamingwebsite_testtask/tournament"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestManager_CreateNewPlayer(t *testing.T) {
	players := &MockPlayerDB{}
	tours := &MockTournamentDB{}
	m := NewManager(players, tours)
	ctx := context.Background()

	tests := []struct {
		name          string
		playerName    string
		expectedID    int
		expectedError error
	}{
		{
			name:          "CreateNewPlayer Success",
			playerName:    "player1",
			expectedID:    1,
			expectedError: nil,
		},
		{
			name:          "CreateNewPlayer Error",
			playerName:    "player2",
			expectedError: errors.New("wrong id"),
		},
	}

	for _, tt := range tests {
		players.On("AddPlayer", ctx, tt.playerName).Return(tt.expectedID, tt.expectedError)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			id, err := m.CreateNewPlayer(ctx, tt.playerName)
			if tt.expectedError != nil {
				assert.Error(t, err, tt.expectedError.Error())

			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedID, id)
			}
		})
	}
	players.AssertExpectations(t)
}

func TestManager_GetPlayerPoints(t *testing.T) {
	players := &MockPlayerDB{}
	tours := &MockTournamentDB{}
	m := NewManager(players, tours)
	ctx := context.Background()

	tests := []struct {
		name            string
		playerID        int
		expectedPlayer  *player.Player
		expectedDBError error
	}{
		{
			name:     "GetPlayerPoints Success",
			playerID: 1,
			expectedPlayer: &player.Player{
				ID:      1,
				Balance: 1.5,
			},
			expectedDBError: nil,
		},
		{
			name:            "GetPlayerPoints Error",
			playerID:        2,
			expectedPlayer:  nil,
			expectedDBError: errors.New("wrong id"),
		},
	}
	for _, tt := range tests {
		players.On("PlayerByID", ctx, tt.playerID).Return(tt.expectedPlayer, tt.expectedDBError)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance, err := m.GetPlayerPoints(ctx, tt.playerID)
			if tt.expectedDBError != nil {
				assert.Error(t, err, fmt.Sprintf("cannot get player ID: %v", tt.expectedDBError))
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedPlayer.Balance, balance)
			}
		})
	}
	players.AssertExpectations(t)
}

func TestManager_TakePointsFromPlayer(t *testing.T) {
	players := &MockPlayerDB{}
	tours := &MockTournamentDB{}
	m := NewManager(players, tours)
	ctx := context.Background()

	tests := []struct {
		testName                  string
		playerID                  int
		points                    float32
		expectedPlayerByID        *player.Player
		expectedPlayerByIDError   error
		updatePlayer              *player.Player
		expectedUpdatePlayerError error
		expectedBalance           float32
		expectedError             string
	}{
		{
			testName: "Success",
			playerID: 1,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      1,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      1,
				Balance: 2.5,
			},
			expectedBalance: 2.5,
		},
		{
			testName:                "PlayerByIDError",
			playerID:                2,
			points:                  1.5,
			expectedPlayerByID:      nil,
			expectedPlayerByIDError: errors.New("wrong id"),
			expectedError:           fmt.Sprint("cannot get player ID: wrong id"),
		},
		{
			testName: "BalanceError",
			playerID: 3,
			points:   10,
			expectedPlayerByID: &player.Player{
				ID:      3,
				Balance: 4.0,
			},
			expectedError: ErrNotEnoughBalance.Error(),
		},
		{
			testName: "UpdatePlayerError",
			playerID: 4,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      4,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      4,
				Balance: 2.5,
			},
			expectedUpdatePlayerError: errors.New("update error"),
			expectedError:             fmt.Sprintf("update error"),
		},
	}
	for _, tt := range tests {
		players.On("PlayerByID", ctx, tt.playerID).Return(tt.expectedPlayerByID, tt.expectedPlayerByIDError)
		if tt.updatePlayer != nil {
			players.On("UpdatePlayer", ctx, tt.playerID, *tt.updatePlayer).Return(tt.expectedUpdatePlayerError)
		}
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			balance, err := m.TakePointsFromPlayer(ctx, tt.playerID, tt.points)
			if tt.expectedError != "" {
				assert.Error(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPlayerByID.Balance, balance)
			}
		})
	}
	players.AssertExpectations(t)
}

func TestManager_FundPointsToPlayer(t *testing.T) {
	players := &MockPlayerDB{}
	tours := &MockTournamentDB{}
	m := NewManager(players, tours)
	ctx := context.Background()

	tests := []struct {
		testName                  string
		playerID                  int
		points                    float32
		expectedPlayerByID        *player.Player
		expectedPlayerByIDError   error
		updatePlayer              *player.Player
		expectedUpdatePlayerError error
		expectedBalance           float32
		expectedError             string
	}{
		{
			testName: "Success",
			playerID: 1,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      1,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      1,
				Balance: 5.5,
			},
			expectedBalance: 5.5,
		},
		{
			testName:                "PlayerByIDError",
			playerID:                2,
			points:                  1.5,
			expectedPlayerByID:      nil,
			expectedPlayerByIDError: errors.New("wrong id"),
			expectedError:           fmt.Sprint("cannot get player ID: wrong id"),
		},
		{
			testName: "UpdatePlayerError",
			playerID: 4,
			points:   1.5,
			expectedPlayerByID: &player.Player{
				ID:      4,
				Balance: 4.0,
			},
			updatePlayer: &player.Player{
				ID:      4,
				Balance: 5.5,
			},
			expectedUpdatePlayerError: errors.New("update error"),
			expectedError:             fmt.Sprintf("update error"),
		},
	}
	for _, tt := range tests {
		players.On("PlayerByID", ctx, tt.playerID).Return(tt.expectedPlayerByID, tt.expectedPlayerByIDError)
		if tt.updatePlayer != nil {
			players.On("UpdatePlayer", ctx, tt.playerID, *tt.updatePlayer).Return(tt.expectedUpdatePlayerError)
		}
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			balance, err := m.FundPointsToPlayer(ctx, tt.playerID, tt.points)
			if tt.expectedError != "" {
				assert.Error(t, err, tt.expectedError)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedPlayerByID.Balance, balance)
			}
		})
	}
	players.AssertExpectations(t)
}

func TestManager_RemovePlayer(t *testing.T) {
	players := &MockPlayerDB{}
	tours := &MockTournamentDB{}
	m := NewManager(players, tours)
	ctx := context.Background()

	tests := []struct {
		testName      string
		playerID      int
		expectedError error
	}{
		{
			testName:      "Success",
			playerID:      1,
			expectedError: nil,
		},
		{
			testName:      "Error",
			playerID:      -1,
			expectedError: errors.New("wrong id"),
		},
	}
	for _, tt := range tests {
		players.On("DeletePlayer", ctx, tt.playerID).Return(tt.expectedError)
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			err := m.RemovePlayer(ctx, tt.playerID)
			if tt.expectedError != nil {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
	players.AssertExpectations(t)
}

func TestManager_AnnounceTournament(t *testing.T) {
	players := &MockPlayerDB{}
	tours := &MockTournamentDB{}
	m := NewManager(players, tours)

	tt := []struct {
		name          string
		deposit       float32
		expectedID    int
		expectedError error
	}{
		{
			name:          "Success",
			deposit:       1.0,
			expectedID:    1,
			expectedError: nil,
		},
		{
			name:          "Error_negativeDeposit",
			deposit:       -1.0,
			expectedError: errors.New("cannot announce tournament: deposit is negative"),
		},
		{
			name:          "Error_dbError",
			deposit:       1.1,
			expectedError: errors.New("some error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tours.On("CreateTournament", tc.deposit).Return(tc.expectedID, tc.expectedError)

			id, err := m.AnnounceTournament(tc.deposit)
			if tc.expectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expectedID, id)
			}
		})
	}
}

func TestManager_JoinTournament(t *testing.T) {
	players := &MockPlayerDB{}
	tours := &MockTournamentDB{}
	m := NewManager(players, tours)

	tt := []struct {
		name                          string
		tourID                        int
		playerID                      int
		expectedTourByID              tournament.Tournament
		expectedTourByIDError         error
		expectedPlayerByID            *player.Player
		expectedPlayerByIDError       error
		updatePlayer                  *player.Player
		expectedUpdatePlayerError     error
		updateTournamentTour          *tournament.Tournament
		expectedUpdateTournamentError error
		expectedError                 error
	}{
		{
			name:     "Success",
			tourID:   1,
			playerID: 11,
			expectedTourByID: tournament.Tournament{
				ID:           1,
				IsFinished:   false,
				Deposit:      1.0,
				Fund:         0,
				Participants: nil,
			},
			expectedPlayerByID: &player.Player{
				ID:      11,
				Balance: 10.0,
			},
			updatePlayer: &player.Player{
				ID:      11,
				Balance: 9.0,
			},
			updateTournamentTour: &tournament.Tournament{
				ID:           1,
				IsFinished:   false,
				Deposit:      1.0,
				Fund:         1.0,
				Participants: []int{11},
			},
		},
		{
			name:     "Error_UpdateTournament",
			tourID:   2,
			playerID: 22,
			expectedTourByID: tournament.Tournament{
				ID:           2,
				IsFinished:   false,
				Deposit:      1.0,
				Fund:         0,
				Participants: nil,
			},
			expectedPlayerByID: &player.Player{
				ID:      22,
				Balance: 10.0,
			},
			updatePlayer: &player.Player{
				ID:      22,
				Balance: 9.0,
			},
			updateTournamentTour: &tournament.Tournament{
				ID:           2,
				IsFinished:   false,
				Deposit:      1.0,
				Fund:         1.0,
				Participants: []int{22},
			},
			expectedUpdateTournamentError: errors.New("error"),
			expectedError:                 errors.New("cannot join to tournament: error"),
		},
		{
			name:     "Error_TakePointsFromPlayer",
			tourID:   3,
			playerID: 33,
			expectedTourByID: tournament.Tournament{
				ID:           3,
				IsFinished:   false,
				Deposit:      1.0,
				Fund:         0,
				Participants: nil,
			},
			expectedPlayerByID: &player.Player{
				ID:      33,
				Balance: 10.0,
			},
			updatePlayer: &player.Player{
				ID:      33,
				Balance: 9.0,
			},
			expectedUpdatePlayerError: errors.New("error"),
			expectedError:             errors.New("cannot join to tournament: error"),
		},
		{
			name:     "Error_TourIsFinished",
			tourID:   4,
			playerID: 44,
			expectedTourByID: tournament.Tournament{
				ID:           4,
				IsFinished:   true,
				Deposit:      1.0,
				Fund:         0,
				Participants: nil,
			},
			expectedPlayerByID: &player.Player{
				ID:      11,
				Balance: 10.0,
			},
			expectedError: errors.New("cannot join to tournament: tournament is finished"),
		},
		{
			name:                  "Error_TourIsNotFound",
			tourID:                5,
			playerID:              55,
			expectedTourByID:      tournament.Tournament{},
			expectedTourByIDError: errors.New("error"),
			expectedError:         errors.New("cannot join to tournament: error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tours.On("TournamentByID", tc.tourID).Return(tc.expectedTourByID, tc.expectedTourByIDError)
			if tc.expectedPlayerByID != nil || tc.expectedPlayerByIDError != nil {
				players.On("PlayerByID", tc.playerID).Return(tc.expectedPlayerByID, tc.expectedPlayerByIDError)
			}
			if tc.updatePlayer != nil {
				players.On("UpdatePlayer", tc.playerID, *tc.updatePlayer).Return(tc.expectedUpdatePlayerError)
			}
			if tc.updateTournamentTour != nil {
				tours.On("UpdateTournament", tc.tourID, *tc.updateTournamentTour).Return(tc.expectedUpdateTournamentError)
			}
			err := m.JoinTournament(tc.tourID, tc.playerID)
			if tc.expectedError != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
