package mongo

import (
	"fmt"

	"github.com/Ragnar-BY/gamingwebsite_testtask/player"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// PlayerService is for using Mongo Collection "players".
type PlayerService struct {
	players *mgo.Collection
	counter *mgo.Collection
}

// NewPlayerService returns new service from Collection.
func NewPlayerService(players *mgo.Collection, counter *mgo.Collection) PlayerService {
	return PlayerService{players: players, counter: counter}
}

// PlayerByID returns player by id, if exist.
func (ps *PlayerService) PlayerByID(id int) (*player.Player, error) {
	model := PlayerModel{}
	err := ps.players.Find(bson.M{"playerId": id}).One(&model)
	return model.ToPlayer(), err
}

// AddPlayer inserts new player in collection.
func (ps *PlayerService) AddPlayer(name string) (int, error) {
	playerID, err := ps.getAndIncreasePlayerID()
	if err != nil {
		return 0, fmt.Errorf("cannot get new id: %v", err)
	}
	model := PlayerModel{PlayerID: playerID, Name: name, Balance: 0}
	err = ps.players.Insert(&model)
	if err != nil {
		return 0, fmt.Errorf("cannot add new player: %v", err)
	}
	return playerID, nil
}

// DeletePlayer deletes player by id from collection, if possible.
func (ps *PlayerService) DeletePlayer(id int) error {
	return ps.players.Remove(bson.M{"playerId": id})
}

// UpdatePlayer updates player with player id from collection with player.Player, if possible.
func (ps *PlayerService) UpdatePlayer(id int, player player.Player) error {
	return ps.players.Update(bson.M{"playerId": id}, bson.M{"$set": bson.M{"balance": player.Balance, "name": player.Name}})
}

// getAndIncreasePlayerID return last player ID and increase it in collection
func (ps *PlayerService) getAndIncreasePlayerID() (int, error) {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"playerId": 1}},
		ReturnNew: false,
	}
	var result bson.M
	_, err := ps.counter.Find(bson.M{"_id": "playerIdCounter"}).Apply(change, &result)
	if err != nil {
		return 0, err
	}
	return result["playerId"].(int), nil
}

func (ps *PlayerService) deleteAllPlayers() error {
	_, err := ps.players.RemoveAll(nil)
	if err != nil {
		return fmt.Errorf("cannot remove all players")
	}
	return nil
}
