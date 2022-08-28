package pieces

import "gochess/board"

type Rook struct {
	Piece
}

// The function returns the board state after all available moves from a Rook Piece
func (r Rook) MovesAvailable(board board.Board, configs board.Configuration) (states []board.BoardState) {
	coordinates := r.upMoves(board, 8, NORULE)
	coordinates = append(coordinates, r.downMoves(board, 8, NORULE)...)
	coordinates = append(coordinates, r.leftMoves(board, 8, NORULE)...)
	coordinates = append(coordinates, r.rightMoves(board, 8, NORULE)...)
	states = r.CreateBoardStates(coordinates, board, configs)
	for _, state := range states {
		state.Configuration = state.Configuration.RemoveCastlingRights(r.IsWhite)
	}
	return
}
