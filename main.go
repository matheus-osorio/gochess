package main

import (
	"fmt"
	"gochess/board/chessprinter"
	"gochess/board/fen"
	"os"
)

func main() {
	board, _, err := fen.CreateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	if err != nil {
		fmt.Println("Error on FEN string parsing")
		os.Exit(1)
	}

	chessprinter.PrintBoard(*board)
}
