package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ragnar-BY/gamingwebsite_testtask/manager"
	"github.com/gorilla/mux"
)

// managerRouter is for manager.
type managerRouter struct {
	manager manager.Manager
}

func newManagerRouter(manager manager.Manager, router *mux.Router) *mux.Router {
	mngrRouter := managerRouter{manager}
	router.HandleFunc("/add", mngrRouter.addPlayerHandler).
		Methods(http.MethodPost).
		Queries("name", "{name}")
	router.HandleFunc("/balance/{playerId:[0-9]+}", mngrRouter.balancePlayerHandler).
		Methods(http.MethodGet)
	router.HandleFunc("/fund/{playerId:[0-9]+}", mngrRouter.fundPointsHandler).
		Methods(http.MethodPut).
		Queries("points", "{points:[0-9]+\\.?[0-9]{0,2}}")
	router.HandleFunc("/take/{playerId:[0-9]+}", mngrRouter.takePointsHandler).
		Methods(http.MethodPut).
		Queries("points", "{points:[0-9]+\\.?[0-9]{0,2}}")
	return router
}

// addPlayerHandler creates new player, returns id.
func (m *managerRouter) addPlayerHandler(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]
	if name == "" {
		Error(w, http.StatusBadRequest, "wrong name")
		return
	}
	id, err := m.manager.CreateNewPlayer(name)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot create new player: %v", err))
		return
	}
	JSON(w, http.StatusCreated, id)
}

// balancePlayerHandler returns player balance.
func (m *managerRouter) balancePlayerHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get playerId: %v", err))
		return
	}
	balance, err := m.manager.GetPlayerPoints(playerID)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get player points: %v", err))
		return
	}
	JSON(w, http.StatusOK, balance)
}

// fundPointsHandler gives points to player, returns new balance.
func (m *managerRouter) fundPointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, points, err := getPlayerIDAndPoints(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	balance, err := m.manager.FundPointsToPlayer(playerID, points)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot fund points to player: %v", err))
		return
	}
	JSON(w, http.StatusOK, balance)
}

///takePointsHandler takes points if possible from player, returns new balance.
func (m *managerRouter) takePointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, points, err := getPlayerIDAndPoints(r)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	balance, err := m.manager.TakePointsFromPlayer(playerID, points)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot take points from player: %v", err))
		return
	}
	JSON(w, http.StatusOK, balance)
}

func getPlayerIDAndPoints(r *http.Request) (int, float32, error) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get playerId: %v", err)
	}
	points, err := getFloatValue(r, "points")
	if err != nil {
		return 0, 0, fmt.Errorf("cannot get points: %v", err)
	}
	return playerID, points, nil
}

// TODO it is better move this functions to other file.
func getIntValue(r *http.Request, key string) (int, error) {
	val := mux.Vars(r)[key]
	return strconv.Atoi(val)
}

func getFloatValue(r *http.Request, key string) (float32, error) {
	val := mux.Vars(r)[key]
	f64, err := strconv.ParseFloat(val, 32)
	return float32(f64), err
}
