package chessprinter

import (
	"fmt"
	"gochess/board"
	"reflect"
	"strings"
)

func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}

func PrintMoveData(data board.MoveData) {
	originLetter := 'A' + data.Origin%8
	originNumber := 8 - data.Origin/8
	destinyLetter := 'A' + data.Destiny%8
	destinyNumber := 8 - data.Destiny/8
	fmt.Printf("Move:\nPiece:%s\nWhite:%t\nFrom:%c%d\nTo: %c%d\n\n", data.Piece, data.IsWhite, originLetter, originNumber, destinyLetter, destinyNumber)
}

func PrintBoard(board board.Board) {
	boardString := "\n     A   B   C   D   E   F   G   H"
	for i, piece := range board {

		if i%8 == 0 {
			extra := ""
			if i > 0 {
				extra = fmt.Sprint(8 - i/8 + 1)
			}
			boardString += fmt.Sprintf(" %s\n   ---------------------------------\n%d  |", extra, 8-i/8)
		}

		var pieceStr string
		if piece == nil {
			pieceStr = " "
		} else {
			switch reflect.TypeOf(piece).Name() {
			case "Rook":
				pieceStr = "r"

			case "Knight":
				pieceStr = "k"

			case "Bishop":
				pieceStr = "b"

			case "Queen":
				pieceStr = "q"

			case "King":
				pieceStr = "k"

			case "Pawn":
				pieceStr = "p"
			}

			if piece.IsWhitePiece() {
				pieceStr = strings.ToUpper(pieceStr)
			}
		}

		boardString += fmt.Sprintf(" %s |", pieceStr)
	}
	boardString += " 1\n   ---------------------------------\n"
	boardString += "     A   B   C   D   E   F   G   H"
	fmt.Println(boardString)
}
