package pieces

import (
	"gochess/board"
)

type Pawn struct {
	Piece
}

// The function returns the board state after all available moves from a Pawn Piece
func (p Pawn) MovesAvailable(chessboard board.Board, configs board.Configuration, coordinate int) (states []board.BoardState) {
	var diagLeftFunc func(board.Board, int, int, int) []int
	var diagRightFunc func(board.Board, int, int, int) []int
	var fowardFunc func(board.Board, int, int, int) []int

	var firstMove bool
	if p.IsWhite {
		diagLeftFunc = p.upLeftMoves
		diagRightFunc = p.upRightMoves
		fowardFunc = p.upMoves

		firstMove = coordinate >= 48 && coordinate < 56
	} else {
		diagLeftFunc = p.downLeftMoves
		diagRightFunc = p.downRightMoves
		fowardFunc = p.downMoves

		firstMove = coordinate >= 8 && coordinate < 16
	}

	coordinates := diagLeftFunc(chessboard, 1, MUSTCAPTURE, coordinate)
	coordinates = append(coordinates, diagRightFunc(chessboard, 1, MUSTCAPTURE, coordinate)...)

	if firstMove {
		coordinates = append(coordinates, fowardFunc(chessboard, 2, MUSTNOTCAPTURE, coordinate)...)
	} else {
		coordinates = append(coordinates, fowardFunc(chessboard, 1, MUSTNOTCAPTURE, coordinate)...)
	}

	return p.CreateBoardStates(coordinates, chessboard, configs, "Pawn", coordinate)
}
