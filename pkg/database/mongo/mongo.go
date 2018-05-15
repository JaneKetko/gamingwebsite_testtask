package mongo

import (
	"github.com/globalsign/mgo"
)

var (
	// Address is Mongo address.
	Address = "127.0.0.1:27017"
	// DBName is name of MongoDB.
	DBName = "GamingDB"
	//PlayerCollectionName is name of players collection.
	PlayerCollectionName = "players"
)

// Session is Mongo session.
type Session struct {
	session *mgo.Session
}

// Open opens Mongo session.
func (s *Session) Open() error {
	var err error
	s.session, err = mgo.Dial(Address)
	return err
}

// Close closes Mongo session.
func (s *Session) Close() {
	if s.session != nil {
		s.session.Close()
	}
}

// Players return new player service from DB.
func (s *Session) Players() PlayerService {
	playerCollection := s.session.DB(DBName).C(PlayerCollectionName)
	return NewPlayerService(playerCollection)
}
