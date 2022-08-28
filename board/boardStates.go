package board

// Holds all the parameters necessary to see available special moves like castling and en paesant
type BoardState struct {
	Board         Board
	Configuration Configuration
}
