package pieces

import (
	"fmt"
	"gochess/board"
)

type King struct {
	Piece
}

// The function returns the board state after all available moves from the King
func (k King) MovesAvailable(board board.Board, configs board.Configuration, coordinate int) []board.BoardState {
	coordinates := k.upMoves(board, 1, NORULE, coordinate)
	coordinates = append(coordinates, k.downMoves(board, 1, NORULE, coordinate)...)
	coordinates = append(coordinates, k.leftMoves(board, 1, NORULE, coordinate)...)
	coordinates = append(coordinates, k.rightMoves(board, 1, NORULE, coordinate)...)
	coordinates = append(coordinates, k.upLeftMoves(board, 1, NORULE, coordinate)...)
	coordinates = append(coordinates, k.upRightMoves(board, 1, NORULE, coordinate)...)
	coordinates = append(coordinates, k.downLeftMoves(board, 1, NORULE, coordinate)...)
	coordinates = append(coordinates, k.downRightMoves(board, 1, NORULE, coordinate)...)
	states := k.CreateBoardStates(coordinates, board, configs, "King", coordinate)
	states = append(states, k.castling(board, configs, coordinate)...)
	for _, state := range states {
		state.Configuration = state.Configuration.RemoveCastlingRights(k.IsWhite)
	}
	return states
}

func (k King) castleKingSide(chessboard board.Board, configs board.Configuration, coordinate int) board.BoardState {
	return board.BoardState{
		Board:         chessboard.CastleKingSide(coordinate),
		Configuration: configs.ChangeTurn(),
	}
}

func (k King) castleQueenSide(chessboard board.Board, configs board.Configuration, coordinate int) board.BoardState {
	if coordinate == 3 {
		fmt.Println(configs)
	}
	return board.BoardState{
		Board:         chessboard.CastleQueenSide(coordinate),
		Configuration: configs.ChangeTurn(),
	}
}

func (k King) castling(chessboard board.Board, configs board.Configuration, coordinate int) (states []board.BoardState) {
	castling := configs.GetCastlingRights(k.IsWhite)
	if castling.King {
		if moves := k.rightMoves(chessboard, 3, MUSTNOTCAPTURE, coordinate); len(moves) == 2 {
			states = append(states, k.castleKingSide(chessboard, configs, coordinate))
		}
	}

	if castling.Queen {
		if moves := k.leftMoves(chessboard, 4, MUSTNOTCAPTURE, coordinate); len(moves) == 3 {
			states = append(states, k.castleQueenSide(chessboard, configs, coordinate))
		}
	}

	return
}
