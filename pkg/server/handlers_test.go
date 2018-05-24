package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"errors"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
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
			Body().Contains("wrong name") // do we need to check error message or only status?
	})
	t.Run("DBError", func(t *testing.T) {
		db.On("AddPlayer", "player2").Return(0, errors.New("cannot add new player"))
		e.Request(http.MethodPost, "/add").WithQuery("name", "player2").
			Expect().Status(http.StatusBadRequest) //same question as above
	})
	db.AssertExpectations(t)
}
