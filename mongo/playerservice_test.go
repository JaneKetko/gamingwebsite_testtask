// +build !notmongo

package mongo

import (
	"log"
	"testing"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var players *PlayerService

func init() {
	var s Session
	err := s.Open("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	players, err = s.Players("testDB", "players")
	if err != nil {
		log.Fatal(err)
	}
}

func cleanCollection(t *testing.T) {
	err := players.deleteAllPlayers()
	assert.NoError(t, err)
}

func TestPlayerService_AddPlayer(t *testing.T) {
	defer cleanCollection(t)
	id, err := players.AddPlayer("player1")
	require.NoError(t, err)
	assert.NotZero(t, id)
}

//TODO do we need to make test table?

func TestPlayerService_PlayerByID(t *testing.T) {
	defer cleanCollection(t)

	t.Run("Success", func(t *testing.T) {
		id, err := players.AddPlayer("PlayerByID")
		require.NoError(t, err)
		p, err := players.PlayerByID(id)
		require.NoError(t, err)
		assert.Equal(t, "PlayerByID", p.Name)
	})
	t.Run("Error", func(t *testing.T) {
		_, err := players.PlayerByID(-1)
		require.Error(t, err)
	})
}

func TestPlayerService_DeletePlayer(t *testing.T) {
	defer cleanCollection(t)
	id, err := players.AddPlayer("player1")
	require.NoError(t, err)

	t.Run("Success", func(t *testing.T) {
		_, err = players.PlayerByID(id)
		require.NoError(t, err)
		err = players.DeletePlayer(id)
		require.NoError(t, err)
		_, err = players.PlayerByID(id)
		assert.Error(t, err)
	})
	t.Run("DeleteError", func(t *testing.T) {
		err = players.DeletePlayer(-1)
		assert.Error(t, err)
	})

}
func TestPlayerService_UpdatePlayer(t *testing.T) {
	defer cleanCollection(t)

	t.Run("Success", func(t *testing.T) {
		balance := float32(12.34)
		name := "playerUpdate"
		id, err := players.AddPlayer(name)
		require.NoError(t, err)
		p, err := players.PlayerByID(id)
		require.NoError(t, err)
		p.Balance = balance
		err = players.UpdatePlayer(id, *p)
		require.NoError(t, err)
		p2, err := players.PlayerByID(id)
		require.NoError(t, err)
		assert.Equal(t, balance, p2.Balance)
		assert.Equal(t, name, p2.Name)
	})
	t.Run("Error", func(t *testing.T) {
		err := players.UpdatePlayer(-1, player.Player{})
		assert.Error(t, err)
	})
}
func TestPlayerService_GetAndIncreasePlayerID(t *testing.T) {
	defer cleanCollection(t)
	id, err := players.getAndIncreasePlayerID()
	require.NoError(t, err)
	assert.NotZero(t, id)
	id2, err := players.getAndIncreasePlayerID()
	require.NoError(t, err)
	assert.Equal(t, id+1, id2)
}

func TestPlayerService_ListAllPlayers(t *testing.T) {
	defer cleanCollection(t)
	t.Run("Success", func(t *testing.T) {
		names := []string{"p1", "p2", "p3"}
		for _, n := range names {
			_, err := players.AddPlayer(n)
			require.NoError(t, err)
		}
		pls, err := players.listAllPlayers()
		require.NoError(t, err)
		require.Equal(t, len(names), len(pls))
		for i, name := range names {
			assert.Equal(t, name, pls[i].Name)
		}
	})
}

//TODO how to create test for deleteAllPlayers?
