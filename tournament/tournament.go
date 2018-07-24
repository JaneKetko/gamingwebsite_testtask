package tournament

import "github.com/Ragnar-BY/gamingwebsite_testtask/player"

// Tournament represents tournament.
type Tournament struct {
	ID           int
	IsFinished   bool
	Deposit      float32
	Participants []*player.Player
	Winner       *player.Player
}
