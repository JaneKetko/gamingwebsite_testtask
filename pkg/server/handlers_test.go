package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/player"
	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
)

func TestManagerRouter_AddPlayerHandler(t *testing.T) {
	db := &manager.MockDB{}
	m := newManagerRouter(manager.Manager{DB: db}, mux.NewRouter())

	server := httptest.NewServer(m)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	t.Run("Success", func(t *testing.T) {
		db.On("AddPlayer", "player1").Return(1, nil)
		e.Request(http.MethodPost, "/add").WithQuery("name", "player1").
			Expect().Status(http.StatusCreated).JSON().Number().Equal(1)
	})
	t.Run("WrongName", func(t *testing.T) {
		e.Request(http.MethodPost, "/add").WithQuery("name", "").
			Expect().Status(http.StatusBadRequest).
			Body().Contains("wrong name") // TODO:do we need to check error message or only status?
	})
	t.Run("DBError", func(t *testing.T) {
		db.On("AddPlayer", "player2").Return(0, errors.New("cannot add new player"))
		e.Request(http.MethodPost, "/add").WithQuery("name", "player2").
			Expect().Status(http.StatusBadRequest) //same question as above
	})
	db.AssertExpectations(t)
}

func TestManagerRouter_balancePlayerHandler(t *testing.T) {
	db := &manager.MockDB{}
	m := newManagerRouter(manager.Manager{DB: db}, mux.NewRouter())
	server := httptest.NewServer(m)
	defer server.Close()
	e := httpexpect.New(t, server.URL)

	t.Run("Success", func(t *testing.T) {
		db.On("PlayerByID", 1).Return(&player.Player{
			ID:      1,
			Balance: 1.5,
		}, nil)
		e.Request(http.MethodGet, "/balance/1").
			Expect().Status(http.StatusOK).JSON().Number().Equal(1.5)
	})

	t.Run("DBError", func(t *testing.T) {
		db.On("PlayerByID", 2).Return(nil, errors.New("some error"))
		e.Request(http.MethodGet, "/balance/2").
			Expect().Status(http.StatusBadRequest)
	})

	t.Run("PlayerParseIDError", func(t *testing.T) {
		e.Request(http.MethodGet, "/balance/98765432109876543210").
			Expect().Status(http.StatusBadRequest)
	})
	//TODO: do we need to check such cases?
	t.Run("PlayerWrongIDError", func(t *testing.T) {
		e.Request(http.MethodGet, "/balance/wrongid").
			Expect().Status(http.StatusNotFound)
	})

}
