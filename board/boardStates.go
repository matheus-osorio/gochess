package board

import (
	"reflect"
)

type MoveData struct {
	Piece   string
	IsWhite bool
	Origin  int
	Destiny int
}

// Holds all the parameters necessary to see available special moves like castling and en paesant
type BoardState struct {
	Board         Board
	Configuration Configuration
	MoveData      MoveData
}

func (state BoardState) CalculatePoints() float64 {
	totalPoint := 0.0
	for _, piece := range state.Board {
		if piece == nil {
			continue
		}
		var sign float64
		if piece.IsWhitePiece() {
			sign = 1.0
		} else {
			sign = -1.0
		}
		switch reflect.TypeOf(piece).Name() {
		case "Rook":
			totalPoint += 5 * sign
		case "Knight":
			totalPoint += 3 * sign
		case "Bishop":
			totalPoint += 3 * sign
		case "Queen":
			totalPoint += 8 * sign
		case "King":
			totalPoint += 1000 * sign
		case "Pawn":
			totalPoint += 1 * sign
		}
	}
	return totalPoint
}

func (state *BoardState) CreateBoardState(origin int, destiny int, chessboard Board, configs Configuration, pieceName string, isWhite bool) {
	newBoard := make(Board, 64)
	copy(newBoard, chessboard)
	newBoard[origin], newBoard[destiny] = nil, newBoard[origin]
	state.Configuration = configs.ChangeTurn()
	state.Board = newBoard
	state.MoveData = MoveData{
		Piece:   pieceName,
		IsWhite: isWhite,
		Origin:  origin,
		Destiny: destiny,
	}
}
