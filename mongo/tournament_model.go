package mongo

import "github.com/globalsign/mgo/bson"

// TournamentModel represents Tournament model for MongoDB.
type TournamentModel struct {
	ID             bson.ObjectId `bson:"_id"`
	TournamentID   int           `bson:"tournamentid"`
	IsFinished     bool          `bson:"isfinished"`
	Deposit        float32       `bson:"deposit"`
	ParticipantIDs []int         `bson:"participants,omitempty"`
	WinnerID       int           `bson:"winner,omitempty"`
}
