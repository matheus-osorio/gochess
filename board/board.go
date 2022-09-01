package board

const FREESPOT = 0
const WHITEPIECE = 1
const BLACKPIECE = 2

// defines everything that moves in the game
type Movable interface {
	MovesAvailable(Board, Configuration, int) []BoardState
	IsWhitePiece() bool
}

// The board. Has methods for board management
type Board []Movable

// This struct removes all the drawbacks from having a 1 dimensional array as the board
type BoardPlace struct {
	Up    int
	Down  int
	Left  int
	Right int
}

// Creates a new instance of the board so that it can be modified independently
func (b Board) Copy() Board {
	board := make(Board, 64)
	copy(board, b)
	return board
}

// The function recieves a coordinate and returns how many squares there are on all imaediate positions
func (b Board) GetPlaceInBoard(coordinate int) BoardPlace {
	place := BoardPlace{}
	place.Up = coordinate / 8
	place.Left = coordinate % 8
	place.Right = 7 - (coordinate % 8)
	place.Down = 7 - (coordinate / 8)
	return place
}

// The function returns simple data about the given square
func (b Board) GetCoordinateInfo(coordinate int) int {
	piece := b[coordinate]
	if piece == nil {
		return FREESPOT
	}

	if piece.IsWhitePiece() {
		return WHITEPIECE
	}

	return BLACKPIECE
}

// This function changes the board after a King side Castle
func (b Board) CastleKingSide(coordinate int) Board {
	var board Board
	copy(board, b)
	board[coordinate], board[coordinate+1], board[coordinate+2], board[coordinate+3] = nil, board[coordinate+3], board[coordinate], nil
	return board
}

// This function changes the board after a Queen side Castle
func (b Board) CastleQueenSide(coordinate int) Board {
	board := b.Copy()
	board[coordinate], board[coordinate-1], board[coordinate-2], board[coordinate-4] = nil, board[coordinate-4], board[coordinate], nil
	return board
}
