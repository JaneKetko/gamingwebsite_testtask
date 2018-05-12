package server

import (
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//ManagerRouter for manager
type ManagerRouter struct {
	manager manager.Manager
}

//NewManagerRouter return new ManagerRouter
func NewManagerRouter(manager manager.Manager, router *mux.Router) *mux.Router {

	managerRouter := ManagerRouter{manager}

	router.HandleFunc("/add", managerRouter.addPlayerHandler).
		Methods("POST").
		Queries("name", "{name}")
	router.HandleFunc("/balance", managerRouter.balancePlayerHandler).
		Methods("GET").
		Queries("playerId", "{playerId:[0-9]+}")
	router.HandleFunc("/fund", managerRouter.fundPointsHandler).
		Methods("PUT").
		Queries("playerId", "{playerId:[0-9]+}", "points", "{points:[0-9]+}")
	router.HandleFunc("/take", managerRouter.takePointsHandler).
		Methods("PUT").
		Queries("playerId", "{playerId:[0-9]+}", "points", "{points:[0-9]+}")

	return router
}

//addPlayerHandler create new player, return id.
func (m *ManagerRouter) addPlayerHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		Error(w, http.StatusBadRequest, "wrong name")
		return
	}
	id, err := m.manager.CreateNewPlayer(name)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	JSON(w, http.StatusCreated, id)
}

//balancePlayerHandler return player balance
func (m *ManagerRouter) balancePlayerHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	balance, err := m.manager.GetPlayerPoints(playerID)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	JSON(w, http.StatusOK, balance)
}

//fundPointsHandler give points to player, return new balance
func (m *ManagerRouter) fundPointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	points, err := getIntValue(r, "points")
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	balance, err := m.manager.FundPointsToPlayer(playerID, points)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	JSON(w, http.StatusOK, balance)
}

///takePointsHandler take points if possible from player, return new balance
func (m *ManagerRouter) takePointsHandler(w http.ResponseWriter, r *http.Request) {
	playerID, err := getIntValue(r, "playerId")
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	points, err := getIntValue(r, "points")
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	balance, err := m.manager.TakePointsFromPlayer(playerID, points)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}
	JSON(w, http.StatusOK, balance)
}

func getIntValue(r *http.Request, key string) (int, error) {
	val := r.FormValue(key)
	return strconv.Atoi(val)
}
