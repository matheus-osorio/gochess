package pieces

import "gochess/board"

type Queen struct {
	Piece
}

// The function returns the board state after all available moves from the Queen
func (q Queen) MovesAvailable(board board.Board, configs board.Configuration, coordinate int) (states []board.BoardState) {
	coordinates := q.upMoves(board, 8, NORULE, coordinate)
	coordinates = append(coordinates, q.downMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, q.leftMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, q.rightMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, q.upLeftMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, q.upRightMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, q.downLeftMoves(board, 8, NORULE, coordinate)...)
	coordinates = append(coordinates, q.downRightMoves(board, 8, NORULE, coordinate)...)
	return q.CreateBoardStates(coordinates, board, configs, "Queen", coordinate)
}
