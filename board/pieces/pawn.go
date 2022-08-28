package pieces

import "gochess/board"

type Pawn struct {
	Piece
}

// The function returns the board state after all available moves from a Pawn Piece
func (p Pawn) MovesAvailable(chessboard board.Board, configs board.Configuration) (states []board.BoardState) {

	var diagLeftFunc func(board.Board, int, int) []int
	var diagRightFunc func(board.Board, int, int) []int
	var fowardFunc func(board.Board, int, int) []int

	var firstMove bool

	if p.IsWhite {
		diagLeftFunc = p.upLeftMoves
		diagRightFunc = p.upRightMoves
		fowardFunc = p.upMoves

		firstMove = p.Coordinate >= 48 && p.Coordinate < 56
	} else {
		diagLeftFunc = p.downLeftMoves
		diagRightFunc = p.downRightMoves
		fowardFunc = p.downMoves

		firstMove = p.Coordinate >= 8 && p.Coordinate < 16
	}

	coordinates := diagLeftFunc(chessboard, 1, MUSTCAPTURE)
	coordinates = append(coordinates, diagRightFunc(chessboard, 1, MUSTCAPTURE)...)

	if firstMove {
		coordinates = append(coordinates, fowardFunc(chessboard, 2, MUSTNOTCAPTURE)...)
	} else {
		coordinates = append(coordinates, fowardFunc(chessboard, 1, MUSTNOTCAPTURE)...)
	}

	return p.CreateBoardStates(coordinates, chessboard, configs)
}
