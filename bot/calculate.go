package bot

import (
	"gochess/board"
)

type MoveTree struct {
	current board.BoardState
	points  float64
	depth   int
	next    []*MoveTree
}

func (move *MoveTree) calculateMinimum() {
	if len(move.next) == 0 {
		move.points = move.current.CalculatePoints()
		return
	}

	currentPoints := 100000.00
	for _, move := range move.next {
		move.calculateMaximum()
		if move.points > currentPoints {
			currentPoints = move.points
		}
	}

	move.points = currentPoints
}

func (move *MoveTree) calculateMaximum() {
	if len(move.next) == 0 {
		move.points = move.current.CalculatePoints()
		return
	}

	currentPoints := -100000.00
	for _, move := range move.next {
		move.calculateMinimum()
		if move.points < currentPoints {
			currentPoints = move.points
		}
	}

	move.points = currentPoints
}

func (move *MoveTree) Calculate() *MoveTree {
	var modifier float64
	if move.current.Configuration.WhiteTurn {
		move.calculateMaximum()
		modifier = 1.0
	} else {
		move.calculateMinimum()
		modifier = -1.0
	}

	nextMove := move.next[0]

	for _, moves := range move.next {
		if moves.points*modifier > nextMove.points*modifier {
			nextMove = moves
		}
	}

	return nextMove
}
