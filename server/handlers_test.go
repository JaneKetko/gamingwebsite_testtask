package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Ragnar-BY/gamingwebsite_testtask/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	badPlayerID = "98765432109876543210" // too long for int.
)

func TestManagerRouter_AddPlayerHandler(t *testing.T) {
	db := &manager.MockDB{}
	m := newManagerRouter(manager.NewManager(db), mux.NewRouter())

	server := httptest.NewServer(m)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	type dbArguments struct {
		playerName  string
		returnID    int
		returnError error
	}
	tt := []struct {
		name           string
		dbArgs         *dbArguments
		expectedStatus int
		expectedValue  int
	}{
		{
			name:           "Success",
			dbArgs:         &dbArguments{playerName: "player1", returnID: 1, returnError: nil},
			expectedStatus: http.StatusCreated,
			expectedValue:  1,
		},
		{
			name:           "WrongName",
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "DBError",
			dbArgs:         &dbArguments{playerName: "player2", returnID: 0, returnError: errors.New("cannot add new player")},
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			playerName := ""
			if tc.dbArgs != nil {
				db.On("AddPlayer", mock.Anything, tc.dbArgs.playerName).Return(tc.dbArgs.returnID, tc.dbArgs.returnError)
				playerName = tc.dbArgs.playerName
			}
			res := e.Request(http.MethodPost, "/add").WithQuery("name", playerName).Expect()

			res.Status(tc.expectedStatus)
			if tc.expectedValue != 0 {
				res.JSON().Number().Equal(tc.expectedValue)
			}
		})
	}
	db.AssertExpectations(t)
}

func TestManagerRouter_balancePlayerHandler(t *testing.T) {
	db := &manager.MockDB{}
	m := newManagerRouter(manager.NewManager(db), mux.NewRouter())
	server := httptest.NewServer(m)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	type dbArguments struct {
		playerID     int
		returnPlayer *player.Player
		returnError  error
	}
	tt := []struct {
		name            string
		path            string
		dbArgs          *dbArguments
		expectedStatus  int
		expectedBalance float32
	}{
		{
			name: "Success",
			path: "1",
			dbArgs: &dbArguments{
				playerID:     1,
				returnPlayer: &player.Player{ID: 1, Balance: 1.5},
				returnError:  nil},
			expectedStatus:  http.StatusOK,
			expectedBalance: 1.5,
		},
		{
			name:           "DBError",
			path:           "2",
			dbArgs:         &dbArguments{playerID: 2, returnPlayer: nil, returnError: errors.New("some error")},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "PlayerParseIDError",
			path:           badPlayerID,
			dbArgs:         nil,
			expectedStatus: http.StatusBadRequest,
		},
		{
			name:           "WrongID",
			path:           "wrongid",
			dbArgs:         nil,
			expectedStatus: http.StatusNotFound,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.dbArgs != nil {
				db.On("PlayerByID", mock.Anything, tc.dbArgs.playerID).Return(tc.dbArgs.returnPlayer, tc.dbArgs.returnError)
			}
			expect := e.Request(http.MethodGet, "/balance/"+tc.path).Expect()
			expect.Status(tc.expectedStatus)
			if tc.expectedBalance > 0 {
				expect.JSON().Number().Equal(tc.expectedBalance)
			}
		})
	}
	db.AssertExpectations(t)
}

func TestManagerRouter_fundPointsHandler(t *testing.T) {
	db := &manager.MockDB{}
	m := newManagerRouter(manager.NewManager(db), mux.NewRouter())
	server := httptest.NewServer(m)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	t.Run("Success", func(t *testing.T) {
		db.On("PlayerByID", mock.Anything, 1).Return(&player.Player{
			ID:      1,
			Balance: 1.5,
		}, nil)
		db.On("UpdatePlayer", mock.Anything, 1, player.Player{
			ID:      1,
			Balance: 4.0,
		}).Return(nil)
		e.Request(http.MethodPut, "/fund/1").WithQuery("points", 2.5).
			Expect().Status(http.StatusOK).JSON().Number().Equal(4.0)
	})
	t.Run("PlayerParseIDOrPointsError", func(t *testing.T) {
		e.Request(http.MethodPut, "/fund/"+badPlayerID).WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
	t.Run("DBError", func(t *testing.T) {
		db.On("PlayerByID", mock.Anything, 3).Return(nil, errors.New("some error"))
		e.Request(http.MethodPut, "/fund/3").WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
}

func TestManagerRouter_takePointsHandler(t *testing.T) {
	db := &manager.MockDB{}
	m := newManagerRouter(manager.NewManager(db), mux.NewRouter())
	server := httptest.NewServer(m)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	t.Run("Success", func(t *testing.T) {
		db.On("PlayerByID", mock.Anything, 1).Return(&player.Player{
			ID:      1,
			Balance: 4.0,
		}, nil)
		db.On("UpdatePlayer", mock.Anything, 1, player.Player{
			ID:      1,
			Balance: 1.5,
		}).Return(nil)
		e.Request(http.MethodPut, "/take/1").WithQuery("points", 2.5).
			Expect().Status(http.StatusOK).JSON().Number().Equal(1.5)
	})
	t.Run("PlayerParseIDOrFloatError", func(t *testing.T) {
		e.Request(http.MethodPut, "/take/"+badPlayerID).WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
	t.Run("DBManagerError", func(t *testing.T) {
		db.On("PlayerByID", mock.Anything, 3).Return(nil, errors.New("some error"))
		e.Request(http.MethodPut, "/take/3").WithQuery("points", 2.5).
			Expect().Status(http.StatusBadRequest)
	})
}

func TestManagerRouter_RemovePlayer(t *testing.T) {
	db := &manager.MockDB{}
	m := newManagerRouter(manager.NewManager(db), mux.NewRouter())
	server := httptest.NewServer(m)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	tt := []struct {
		name           string
		playerID       int
		returnError    error
		expectedStatus int
	}{
		{
			name:           "Success",
			playerID:       1,
			returnError:    nil,
			expectedStatus: http.StatusOK,
		},
		{
			name:           "DBError",
			playerID:       2,
			returnError:    errors.New("wrong id"),
			expectedStatus: http.StatusBadRequest,
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			db.On("DeletePlayer", mock.Anything, tc.playerID).Return(tc.returnError)
			e.Request(http.MethodDelete, "/remove/"+strconv.Itoa(tc.playerID)).
				Expect().Status(tc.expectedStatus)
		})
	}
	t.Run("ParseIDError", func(t *testing.T) {
		e.Request(http.MethodDelete, "/remove/"+badPlayerID).
			Expect().Status(http.StatusBadRequest)
	})
}

func TestGetPlayerIDAndPoints(t *testing.T) {
	tt := []struct {
		name             string
		playerID         string
		points           string
		expectError      bool
		expectedPlayerID int
		expectedPoints   float32
	}{
		{name: "Success", playerID: "1", points: "1.5", expectError: false, expectedPlayerID: 1, expectedPoints: 1.5},
		{name: "BadPlayerID", playerID: badPlayerID, points: "1.5", expectError: true},
		{name: "WrongPoints", playerID: "2", points: "9876543210987654321098765432109876543210.91", expectError: true},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest("", "", nil)
			req = mux.SetURLVars(req, map[string]string{
				"playerId": tc.playerID,
				"points":   tc.points,
			})
			playerID, points, err := getPlayerIDAndPoints(req)
			if !tc.expectError {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedPlayerID, playerID)
				assert.Equal(t, tc.expectedPoints, points)
			} else {
				assert.Error(t, err)
			}
		})
	}
}
