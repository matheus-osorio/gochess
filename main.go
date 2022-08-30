package main

import (
	"fmt"
	"gochess/board/fen"
	"gochess/game"
	"os"
)

func main() {
	state, err := fen.CreateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	if err != nil {
		fmt.Println("Error on FEN string parsing")
		os.Exit(1)
	}

	game.Play(*state, 4, 1)
}
