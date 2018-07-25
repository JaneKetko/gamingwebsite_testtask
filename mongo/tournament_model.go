package mongo

import (
	"github.com/Ragnar-BY/gamingwebsite_testtask/tournament"
	"github.com/globalsign/mgo/bson"
)

// TournamentModel represents Tournament model for MongoDB.
type TournamentModel struct {
	ID             bson.ObjectId `bson:"_id,omitempty"`
	TournamentID   int           `bson:"tournamentid"`
	IsFinished     bool          `bson:"isfinished"`
	Deposit        float32       `bson:"deposit"`
	ParticipantIDs []int         `bson:"participants,omitempty"`
	WinnerID       *int          `bson:"winner,omitempty"`
}

// ToTournament converts TournamentsModel to tournament.Tournament
func (tm TournamentModel) ToTournament() tournament.Tournament {
	return tournament.Tournament{
		ID:           tm.TournamentID,
		IsFinished:   tm.IsFinished,
		Deposit:      tm.Deposit,
		Participants: tm.ParticipantIDs,
		Winner:       tm.WinnerID,
	}
}
