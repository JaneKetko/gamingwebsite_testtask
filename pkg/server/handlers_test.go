package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/manager"
	"github.com/gavv/httpexpect"
	"github.com/gorilla/mux"
)

func TestManagerRouter_AddPlayerHandler(t *testing.T) {

	db := &manager.MockDB{}
	m := newManagerRouter(manager.Manager{db}, mux.NewRouter())

	server := httptest.NewServer(m)
	defer server.Close()

	db.On("AddPlayer", "player1").Return(1, nil)
	e := httpexpect.New(t, server.URL)

	e.Request(http.MethodPost, "/add").WithQuery("name", "player1").
		Expect().Status(http.StatusCreated).JSON().Number().Equal(1)

	db.AssertExpectations(t)

}
