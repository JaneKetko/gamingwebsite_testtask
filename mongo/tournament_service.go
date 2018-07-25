package mongo

import (
	"fmt"

	"github.com/Ragnar-BY/gamingwebsite_testtask/tournament"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// TourService is for using Mongo Collection "tours".
type TourService struct {
	tours   *mgo.Collection
	counter *mgo.Collection
}

// NewTourService returns new service from Collection.
func NewTourService(tours *mgo.Collection, counter *mgo.Collection) TourService {
	return TourService{tours: tours, counter: counter}
}

// TournamentByID returns tournament by id, if exist.
func (ts *TourService) TournamentByID(id int) (tournament.Tournament, error) {
	model := TournamentModel{}
	err := ts.tours.Find(bson.M{"tournamentid": id}).One(&model)
	return model.ToTournament(), err
}

// CreateTournament inserts new tournament in collection.
func (ts *TourService) CreateTournament(deposit float32) (int, error) {
	tourID, err := ts.getAndIncreaseTournamentID()
	if err != nil {
		return 0, fmt.Errorf("cannot get new id: %v", err)
	}
	model := TournamentModel{TournamentID: tourID, Deposit: deposit, IsFinished: false}
	err = ts.tours.Insert(&model)
	if err != nil {
		return 0, fmt.Errorf("cannot add new player: %v", err)
	}
	return tourID, nil
}

// DeleteTournament deletes Tournament by id from collection, if possible.
func (ts *TourService) DeleteTournament(id int) error {
	return ts.tours.Remove(bson.M{"tournamentid": id})
}

// UpdateTournament updates player with player id from collection with player.Player, if possible.
func (ts *TourService) UpdateTournament(id int, tour tournament.Tournament) error {
	m := bson.M{
		"isfinished": tour.IsFinished,
		"deposit":    tour.Deposit,
	}
	if tour.Participants != nil {
		m["participiants"] = *tour.Participants
	}
	if tour.Winner != nil {
		m["winner"] = *tour.Winner
	}
	return ts.tours.Update(bson.M{"tournamentid": id}, bson.M{"$set": m})
}

// getAndIncreaseTournamentID return last Tournament ID and increase it in collection
func (ts *TourService) getAndIncreaseTournamentID() (int, error) {
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"tournamentid": 1}},
		ReturnNew: false,
	}
	var result bson.M
	_, err := ts.counter.Find(bson.M{"_id": "tournamentIdCounter"}).Apply(change, &result)
	if err != nil {
		return 0, err
	}
	return result["tournamentid"].(int), nil
}

func (ts *TourService) deleteAllTours() error {
	_, err := ts.tours.RemoveAll(nil)
	if err != nil {
		return fmt.Errorf("cannot remove all tours")
	}
	return nil
}

func (ts *TourService) listAllTours() ([]tournament.Tournament, error) {
	var tournamentModels []TournamentModel
	err := ts.tours.Find(nil).All(&tournamentModels)
	if err != nil {
		return nil, fmt.Errorf("cannot list all tours: %v", err)
	}
	tours := make([]tournament.Tournament, len(tournamentModels))
	for i, pm := range tournamentModels {
		tours[i] = pm.ToTournament()
	}
	return tours, nil
}
