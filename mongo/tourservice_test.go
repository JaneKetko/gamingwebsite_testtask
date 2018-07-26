// +build !notmongo

package mongo

import (
	"log"
	"testing"

	"context"

	"github.com/Ragnar-BY/gamingwebsite_testtask/tournament"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var tours *TourService

func init() {
	var s Session
	err := s.Open("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	tours, err = s.Tournament("testDB", "tours")
	if err != nil {
		log.Fatal(err)
	}
}

// TODO: may be better don`t use this function.

func cleanToursCollection(t *testing.T) {
	err := tours.deleteAllTours(context.Background())
	assert.NoError(t, err)
}
func TestTourService_CreateTournament(t *testing.T) {
	defer cleanToursCollection(t)
	ctx := context.Background()
	id, err := tours.CreateTournament(ctx, 12.34)
	require.NoError(t, err)
	assert.NotZero(t, id)

	id2, err := tours.CreateTournament(ctx, 1.23)
	require.NoError(t, err)
	assert.NotZero(t, id)
	assert.Equal(t, id+1, id2)
}

func TestTourService_TournamentByID(t *testing.T) {
	defer cleanToursCollection(t)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {

		id, err := tours.CreateTournament(ctx, 1.0)
		require.NoError(t, err)
		p, err := tours.TournamentByID(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, float32(1.0), p.Deposit)
	})
	t.Run("Error", func(t *testing.T) {
		_, err := tours.TournamentByID(ctx, -1)
		require.Error(t, err)
	})
}

func TestTourService_DeleteTournament(t *testing.T) {
	defer cleanToursCollection(t)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		id, err := tours.CreateTournament(ctx, 10.0)
		require.NoError(t, err)
		_, err = tours.TournamentByID(ctx, id)
		require.NoError(t, err)
		err = tours.DeleteTournament(ctx, id)
		require.NoError(t, err)
		_, err = tours.TournamentByID(ctx, id)
		assert.Error(t, err)
	})
	t.Run("DeleteError", func(t *testing.T) {
		err := tours.DeleteTournament(ctx, -1)
		assert.Error(t, err)
	})
}
func TestTourService_UpdateTournament(t *testing.T) {
	defer cleanToursCollection(t)
	ctx := context.Background()

	t.Run("Success", func(t *testing.T) {
		deposit := float32(12.34)
		id, err := tours.CreateTournament(ctx, deposit)
		require.NoError(t, err)
		tour, err := tours.TournamentByID(ctx, id)
		require.NoError(t, err)

		winner := 12
		tour.IsFinished = true
		tour.Winner = &winner

		err = tours.UpdateTournament(ctx, id, tour)
		require.NoError(t, err)
		tour2, err := tours.TournamentByID(ctx, id)
		require.NoError(t, err)
		assert.Equal(t, true, tour2.IsFinished)
		assert.Equal(t, winner, *tour2.Winner)
	})
	t.Run("Error", func(t *testing.T) {
		err := tours.UpdateTournament(ctx, -1, tournament.Tournament{})
		assert.Error(t, err)
	})
}
func TestTourService_GetAndIncreasePlayerID(t *testing.T) {
	defer cleanToursCollection(t)
	id, err := tours.getAndIncreaseTournamentID()
	require.NoError(t, err)
	assert.NotZero(t, id)
	id2, err := tours.getAndIncreaseTournamentID()
	require.NoError(t, err)
	assert.Equal(t, id+1, id2)
}

func TestTourService_ListAllPlayers(t *testing.T) {
	defer cleanToursCollection(t)
	ctx := context.Background()
	t.Run("Success", func(t *testing.T) {
		deposits := []float32{1.0, 2.1, 3.21}
		for _, d := range deposits {
			_, err := tours.CreateTournament(ctx, d)
			require.NoError(t, err)
		}
		trs, err := tours.listAllTours(ctx)
		require.NoError(t, err)
		require.Equal(t, len(deposits), len(trs))
		for i, d := range deposits {
			assert.Equal(t, d, trs[i].Deposit)
		}
	})
}

func TestTourService_DeleteAllPlayers(t *testing.T) {
	ctx := context.Background()
	t.Run("Success", func(t *testing.T) {
		deposits := []float32{1.0, 2.1, 3.21}
		for _, d := range deposits {
			_, err := tours.CreateTournament(ctx, d)
			require.NoError(t, err)
		}
		trs, err := tours.listAllTours(ctx)
		require.NoError(t, err)
		l := len(trs)
		assert.Equal(t, len(deposits), l)
		err = tours.deleteAllTours(ctx)
		require.NoError(t, err)

		trs, err = tours.listAllTours(ctx)
		require.NoError(t, err)
		l = len(trs)
		assert.Equal(t, 0, l)
	})
}
