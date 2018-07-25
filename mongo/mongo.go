package mongo

import (
	"fmt"

	"github.com/globalsign/mgo"
)

// Session is Mongo session.
type Session struct {
	session *mgo.Session
}

// Open opens new session with address.
func (s *Session) Open(address string) error {
	var err error
	s.session, err = mgo.Dial(address)
	return err
}

// Players return new player service from DB.
func (s *Session) Players(dbname string, players string) (*PlayerService, error) {
	playerCollection := s.session.DB(dbname).C(players)
	index := mgo.Index{
		Key:    []string{"playerId"},
		Unique: true,
	}
	err := playerCollection.EnsureIndex(index)
	if err != nil {
		return nil, fmt.Errorf("cannot create mongo collection index: %v", err)
	}
	// CounterCollection is collection players+counter
	CounterCollection := s.session.DB(dbname).C(fmt.Sprintf("%scounter", players))

	// dirty hack for adding counter
	type counter struct {
		ID       string `bson:"_id,omitempty"`
		PlayerID int    `bson:"playerId"`
	}

	count, err := CounterCollection.FindId("playerIdCounter").Count() //check if counter already exists
	if err != nil {
		return nil, fmt.Errorf("cannot check if counter exists: %v", err)
	}
	if count == 0 {
		err = CounterCollection.Insert(counter{ID: "playerIdCounter", PlayerID: 1})

		if err != nil {
			return nil, fmt.Errorf("cannot create counter: %v", err)
		}
	}
	ps := NewPlayerService(playerCollection, CounterCollection)
	return &ps, nil
}

func (s *Session) Tournament(dbname string, tours string) (*TourService, error) {
	tourCollection := s.session.DB(dbname).C(tours)
	index := mgo.Index{
		Key:    []string{"tournamentid"},
		Unique: true,
	}
	err := tourCollection.EnsureIndex(index)
	if err != nil {
		return nil, fmt.Errorf("cannot create mongo collection index: %v", err)
	}
	// CounterCollection is collection players+counter
	CounterCollection := s.session.DB(dbname).C(fmt.Sprintf("%scounter", tours))

	// dirty hack for adding counter
	type counter struct {
		ID           string `bson:"_id,omitempty"`
		TournamentID int    `bson:"tournamentid"`
	}

	count, err := CounterCollection.FindId("tournamentIdCounter").Count() //check if counter already exists
	if err != nil {
		return nil, fmt.Errorf("cannot check if counter exists: %v", err)
	}
	if count == 0 {
		err = CounterCollection.Insert(counter{ID: "tournamentIdCounter", TournamentID: 1})

		if err != nil {
			return nil, fmt.Errorf("cannot create counter: %v", err)
		}
	}
	ts := NewTourService(tourCollection, CounterCollection)
	return &ts, nil

}
