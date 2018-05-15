package mongo

import (
	"log"

	"github.com/globalsign/mgo"
)

// TODO take all this variables from config file, environments of flags.
var (
	// Address is Mongo address.
	Address = "127.0.0.1:27017"
	// DBName is name of MongoDB.
	DBName = "GamingDB"
	// PlayerCollectionName is name of players collection.
	PlayerCollectionName = "players"
	// CounterCollectionName is collection with counter for player.
	CounterCollectionName = "playercounter"
)

// Session is Mongo session.
type Session struct {
	session *mgo.Session
}

var session Session

func init() {
	var err error
	session.session, err = mgo.Dial(Address)
	if err != nil {
		log.Fatalf("cannot create mongo session: %v", err)
	}

}

// Players return new player service from DB.
func Players() PlayerService {
	playerCollection := session.session.DB(DBName).C(PlayerCollectionName)
	index := mgo.Index{
		Key:    []string{"playerId"},
		Unique: true,
	}
	err := playerCollection.EnsureIndex(index)
	if err != nil {
		log.Fatalf("cannot create mongo collection index: %v", err)
	}
	CounterCollection := session.session.DB(DBName).C(CounterCollectionName)
	return NewPlayerService(playerCollection, CounterCollection)
}
