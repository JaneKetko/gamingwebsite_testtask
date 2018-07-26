// +build !notmysql

package mysql

import (
	"context"
	"log"
	"testing"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var players *PlayerService

func init() {

	s, err := Open("root", "12345678", "TestDB")
	if err != nil {
		log.Fatal(err)
	}
	players = &PlayerService{DB: s, Name: "players"}

}

// TODO: may be better don`t use this function.

func cleanPlayerTable(t *testing.T) {
	err := players.deleteAllPlayers(context.Background())
	assert.NoError(t, err)
}
func TestPlayerService_AddPlayer(t *testing.T) {
	defer cleanPlayerTable(t)
	id, err := players.AddPlayer(context.Background(), "player1")
	require.NoError(t, err)
	assert.NotZero(t, id)

	id2, err := players.AddPlayer(context.Background(), "player2")
	require.NoError(t, err)
	assert.NotZero(t, id)
	assert.Equal(t, id+1, id2)
}

func TestPlayerService_PlayerByID(t *testing.T) {
	defer cleanPlayerTable(t)

	t.Run("Success", func(t *testing.T) {
		id, err := players.AddPlayer(context.Background(), "PlayerByID")
		require.NoError(t, err)
		p, err := players.PlayerByID(context.Background(), id)
		require.NoError(t, err)
		assert.Equal(t, "PlayerByID", p.Name)
	})
	t.Run("Error", func(t *testing.T) {
		_, err := players.PlayerByID(context.Background(), -1)
		require.Error(t, err)
	})
}

func TestPlayerService_DeletePlayer(t *testing.T) {
	defer cleanPlayerTable(t)
	id, err := players.AddPlayer(context.Background(), "player1")
	require.NoError(t, err)

	t.Run("Success", func(t *testing.T) {
		_, err = players.PlayerByID(context.Background(), id)
		require.NoError(t, err)
		err = players.DeletePlayer(context.Background(), id)
		require.NoError(t, err)
		_, err = players.PlayerByID(context.Background(), id)
		assert.Error(t, err)
	})
	t.Run("DeleteError", func(t *testing.T) {
		err = players.DeletePlayer(context.Background(), -1)
		assert.Error(t, err)
	})

}
func TestPlayerService_UpdatePlayer(t *testing.T) {
	defer cleanPlayerTable(t)

	t.Run("Success", func(t *testing.T) {
		balance := float32(12.34)
		name := "playerUpdate"
		id, err := players.AddPlayer(context.Background(), name)
		require.NoError(t, err)
		p, err := players.PlayerByID(context.Background(), id)
		require.NoError(t, err)
		p.Balance = balance
		err = players.UpdatePlayer(context.Background(), id, *p)
		require.NoError(t, err)
		p2, err := players.PlayerByID(context.Background(), id)
		require.NoError(t, err)
		assert.Equal(t, balance, p2.Balance)
		assert.Equal(t, name, p2.Name)
	})
	t.Run("Error", func(t *testing.T) {
		err := players.UpdatePlayer(context.Background(), -1, player.Player{})
		assert.Error(t, err)
	})
}

func TestPlayerService_ListAllPlayers(t *testing.T) {
	defer cleanPlayerTable(t)
	t.Run("Success", func(t *testing.T) {
		names := []string{"p1", "p2", "p3"}
		for _, n := range names {
			_, err := players.AddPlayer(context.Background(), n)
			require.NoError(t, err)
		}
		pls, err := players.listAllPlayers(context.Background())
		require.NoError(t, err)
		require.Equal(t, len(names), len(pls))
		for i, name := range names {
			assert.Equal(t, name, pls[i].Name)
		}
	})
}

func TestPlayerService_DeleteAllPlayers(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		names := []string{"p1", "p2", "p3"}
		for _, n := range names {
			_, err := players.AddPlayer(context.Background(), n)
			require.NoError(t, err)
		}
		pls, err := players.listAllPlayers(context.Background())
		require.NoError(t, err)
		l := len(pls)
		assert.Equal(t, len(names), l)
		err = players.deleteAllPlayers(context.Background())
		require.NoError(t, err)

		pls, err = players.listAllPlayers(context.Background())
		require.NoError(t, err)
		l = len(pls)
		assert.Equal(t, 0, l)
	})
}
