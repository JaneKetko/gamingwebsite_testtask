package tournament

// Tournament represents tournament.
type Tournament struct {
	ID           int
	IsFinished   bool
	Deposit      float32
	Fund         float32
	Participants []int
	Winner       *int
}
