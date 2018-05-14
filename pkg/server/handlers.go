package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/gorilla/mux"
)

// ManagerRouter is for manager.
type ManagerRouter struct {
	manager manager.Manager
}

// NewManagerRouter returns new ManagerRouter.
func NewManagerRouter(manager manager.Manager, router *mux.Router) *mux.Router {

	managerRouter := ManagerRouter{manager}
	router.HandleFunc("/add/{name}", managerRouter.addPlayerHandler).
		Methods("POST")
	router.HandleFunc("/balance/{playerId:[0-9]+}", managerRouter.balancePlayerHandler).
		Methods("GET")
	router.HandleFunc("/fund", managerRouter.fundPointsHandler).
		Methods("PUT").
		Queries("playerId", "{playerId:[0-9]+}", "points", "{points:[0-9]+}")
	router.HandleFunc("/take", managerRouter.takePointsHandler).
		Methods("PUT").
		Queries("playerId", "{playerId:[0-9]+}", "points", "{points:[0-9]+}")
	return router
}

// addPlayerHandler creates new player, returns id.
func (m *ManagerRouter) addPlayerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
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
func (m *ManagerRouter) balancePlayerHandler(w http.ResponseWriter, r *http.Request) {
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
func (m *ManagerRouter) fundPointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get playerId: %v", err))
		return
	}
	points, err := getFloatValue(r, "points")
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get points: %v", err))
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
func (m *ManagerRouter) takePointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get playerId: %v", err))
		return
	}
	points, err := getFloatValue(r, "points")
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot get points: %v", err))
		return
	}
	balance, err := m.manager.TakePointsFromPlayer(playerID, points)
	if err != nil {
		Error(w, http.StatusBadRequest, fmt.Sprintf("cannot take points from player: %v", err))
		return
	}
	JSON(w, http.StatusOK, balance)
}

// TODO it is better move this function to other file.
func getIntValue(r *http.Request, key string) (int, error) {
	val := r.FormValue(key)
	return strconv.Atoi(val)
}

func getFloatValue(r *http.Request, key string) (float32, error) {
	val := r.FormValue(key)
	f64, err := strconv.ParseFloat(val, 32)
	return float32(f64), err
}
