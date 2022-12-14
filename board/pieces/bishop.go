package pieces

import "gochess/board"

type Bishop struct {
	Piece
}

// The function returns the board state after all available moves from a Bishop Piece
func (b Bishop) MovesAvailable(board board.Board, configs board.Configuration, coordinate int) (states []board.BoardState) {
	coordinates := b.upLeftMoves(board, 8, NORULE, coordinate)
	coordinates = append(coordinates, b.upRightMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, b.downLeftMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, b.downRightMoves(board, 8, NORULE, coordinate)...)
	return b.CreateBoardStates(coordinates, board, configs, "Bishop", coordinate)
}
