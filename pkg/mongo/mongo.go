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
		log.Fatalf("cannot create mongo collection index: %v", err)
	}
	// CounterCollection is collection players+counter
	CounterCollection := s.session.DB(dbname).C(fmt.Sprintf("%scounter", players))
	return NewPlayerService(playerCollection, CounterCollection)
}
