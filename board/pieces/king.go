package pieces

import (
	"gochess/board"
)

type King struct {
	Piece
}

// The function returns the board state after all available moves from the King
func (k King) MovesAvailable(board board.Board, configs board.Configuration) []board.BoardState {
	coordinates := k.upMoves(board, 1, NORULE)
	coordinates = append(coordinates, k.downMoves(board, 1, NORULE)...)
	coordinates = append(coordinates, k.leftMoves(board, 1, NORULE)...)
	coordinates = append(coordinates, k.rightMoves(board, 1, NORULE)...)
	coordinates = append(coordinates, k.upLeftMoves(board, 1, NORULE)...)
	coordinates = append(coordinates, k.upRightMoves(board, 1, NORULE)...)
	coordinates = append(coordinates, k.downLeftMoves(board, 1, NORULE)...)
	coordinates = append(coordinates, k.downRightMoves(board, 1, NORULE)...)
	states := k.CreateBoardStates(coordinates, board, configs)
	states = append(states, k.castling(board, configs)...)
	for _, state := range states {
		state.Configuration = state.Configuration.RemoveCastlingRights(k.IsWhite)
	}
	return states
}

func (k King) castleKingSide(chessboard board.Board, configs board.Configuration) board.BoardState {
	return board.BoardState{
		Board:         chessboard.CastleKingSide(k.Coordinate),
		Configuration: configs.ChangeTurn(),
	}
}

func (k King) castleQueenSide(chessboard board.Board, configs board.Configuration) board.BoardState {
	return board.BoardState{
		Board:         chessboard.CastleQueenSide(k.Coordinate),
		Configuration: configs.ChangeTurn(),
	}
}

func (k King) castling(chessboard board.Board, configs board.Configuration) (states []board.BoardState) {
	castling := configs.GetCastlingRights(k.IsWhite)
	if castling.King {
		if moves := k.rightMoves(chessboard, 3, MUSTNOTCAPTURE); len(moves) == 2 {
			states = append(states, k.castleKingSide(chessboard, configs))
		}
	}

	if castling.Queen {
		if moves := k.leftMoves(chessboard, 4, MUSTNOTCAPTURE); len(moves) == 3 {
			states = append(states, k.castleQueenSide(chessboard, configs))
		}
	}

	return
}
