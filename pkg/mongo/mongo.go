package mongo

import (
	"fmt"
	"log"

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
func (s *Session) Players(dbname string, players string) PlayerService {
	playerCollection := s.session.DB(dbname).C(players)
	index := mgo.Index{
		Key:    []string{"playerId"},
		Unique: true,
	}
	err := playerCollection.EnsureIndex(index)
	if err != nil {
		// TODO you print errors there, but better return this error to main.go and show all errors there.
		log.Fatalf("cannot create mongo collection index: %v", err)
	}
	// CounterCollection is collection players+counter
	CounterCollection := s.session.DB(dbname).C(fmt.Sprintf("%scounter", players))

	// dirty hack for adding counter
	type counter struct {
		ID       string `bson:"_id,omitempty"`
		PlayerID int    `bson:"playerId"`
	}
	err = CounterCollection.Insert(counter{ID: "playerIdCounter", PlayerID: 0})
	if err != nil {
		log.Fatalf("cannot create counter: %v", err)
	}
	return NewPlayerService(playerCollection, CounterCollection)
}
