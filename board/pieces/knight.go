package pieces

import "gochess/board"

const LEFT2UP1 = LEFTMOVE*2 + UPMOVE
const LEFT2DOWN1 = LEFTMOVE*2 + DOWNMOVE

const LEFT1UP2 = LEFTMOVE + UPMOVE*2
const RIGHT1UP2 = RIGHTMOVE + UPMOVE*2

const RIGHT2UP1 = RIGHTMOVE*2 + UPMOVE
const RIGHT2DOWN1 = RIGHTMOVE*2 + DOWNMOVE

const LEFT1DOWN2 = LEFTMOVE + DOWNMOVE*2
const RIGHT1DOWN2 = RIGHTMOVE + DOWNMOVE*2

type Knight struct {
	Piece
}

// The function returns the board state after all available moves from a knight Piece
func (k Knight) MovesAvailable(board board.Board, configs board.Configuration) (states []board.BoardState) {
	coordinate := k.Coordinate
	coordinates := []int{}
	position := board.GetPlaceInBoard(coordinate)

	if position.Left >= 2 && position.Up >= 1 {
		coordinates = append(coordinates, coordinate+LEFT2UP1)
	}
	if position.Left >= 2 && position.Down >= 1 {
		coordinates = append(coordinates, coordinate+LEFT2DOWN1)
	}
	if position.Left >= 1 && position.Up >= 2 {
		coordinates = append(coordinates, coordinate+LEFT1UP2)
	}
	if position.Right >= 1 && position.Up >= 2 {
		coordinates = append(coordinates, coordinate+RIGHT1UP2)
	}
	if position.Right >= 2 && position.Up >= 1 {
		coordinates = append(coordinates, coordinate+RIGHT2UP1)
	}
	if position.Right >= 2 && position.Down >= 1 {
		coordinates = append(coordinates, coordinate+RIGHT2DOWN1)
	}
	if position.Left >= 1 && position.Down >= 2 {
		coordinates = append(coordinates, coordinate+LEFT1DOWN2)
	}
	if position.Right >= 1 && position.Down >= 2 {
		coordinates = append(coordinates, coordinate+RIGHT1DOWN2)
	}

	return k.CreateBoardStates(coordinates, board, configs)
}
