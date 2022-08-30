package game

import (
	"gochess/board"
	"gochess/board/chessprinter"
	"gochess/bot"
)

func Play(state board.BoardState, depth, threads int) {
	for {
		// chessprinter.ClearScreen()
		chessprinter.PrintMoveData(state.MoveData)
		chessprinter.PrintBoard(state.Board)
		state = bot.Play(state, depth, threads)
	}
}
