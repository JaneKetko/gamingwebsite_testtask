package mongo

import (
	"github.com/Ragnar-BY/gamingwebsite_testtask/pkg/player"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// PlayerService is for using Mongo Collection "players".
type PlayerService struct {
	players *mgo.Collection
	//TODO find way to remove counter from here
	counter int
}

// NewPlayerService returns new service from Collection.
func NewPlayerService(players *mgo.Collection) PlayerService {
	return PlayerService{players: players, counter: 0}
}

// PlayerByID returns player by id, if exist.
func (ps *PlayerService) PlayerByID(id int) (*player.Player, error) {
	model := PlayerModel{}
	err := ps.players.Find(bson.M{"playerId": id}).One(&model)
	return model.ToPlayer(), err
}

// AddPlayer inserts new player in collection.
func (ps *PlayerService) AddPlayer(name string) (int, error) {
	model := PlayerModel{PlayerID: ps.counter, Name: name, Balance: 0}
	err := ps.players.Insert(&model)
	if err != nil {
		return 0, err
	}
	ps.counter++
	return ps.counter, nil
}

// DeletePlayer deletes player by id from collection, if possible.
func (ps *PlayerService) DeletePlayer(id int) error {
	return ps.players.Remove(bson.M{"playerId": id})
}

// UpdatePlayer updates player with id from collection with player.Player, if possible.
func (ps *PlayerService) UpdatePlayer(id int, player player.Player) error {
	model := PlayerModel{}
	err := ps.players.Find(bson.M{"playerId": id}).One(&model)
	if err != nil {
		return err
	}
	// TODO add index to playerId in mongoDB and update players using just this playerID
	return ps.players.Update(bson.M{"_id": model.ID}, bson.M{"$set": bson.M{"balance": player.Balance}})
}
