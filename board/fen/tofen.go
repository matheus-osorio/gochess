package fen

import (
	"fmt"
	"gochess/board"
	"reflect"
	"strings"
)

func createBoardPart(chessboard board.Board) string {
	var vacantSpot int
	var fenBoard string
	for i, piece := range chessboard {
		if i%8 == 0 {
			if vacantSpot > 0 {
				fenBoard += fmt.Sprint(vacantSpot)
				vacantSpot = 0
			}
			fenBoard += "/"
		}

		if piece == nil {
			vacantSpot++
			continue
		}

		if vacantSpot > 0 {
			fenBoard += fmt.Sprint(vacantSpot)
			vacantSpot = 0
		}
		pieceInitial := ""
		pieceTypeName := reflect.TypeOf(piece).Name()
		switch pieceTypeName {
		case "Knight":
			pieceInitial = "N"
		default:
			pieceInitial = string(pieceTypeName[0])
		}

		if !piece.IsWhitePiece() {
			pieceInitial = strings.ToLower(pieceInitial)
		}

		fenBoard += pieceInitial
	}

	return fenBoard[1:]
}

func createTurnPart(configs board.Configuration) string {
	configString := ""
	if configs.WhiteTurn {
		configString += "w"
	} else {
		configString += "b"
	}
	return configString
}

func createCastlingPart(configs board.Configuration) string {
	castleString := ""
	turns := func(castles board.BoardSide) string {
		castleString := ""
		if castles.King {
			castleString += "K"
		}
		if castles.Queen {
			castleString += "Q"
		}
		return castleString
	}
	castleString += turns(configs.CastleRights.White)
	castleString += strings.ToLower(turns(configs.CastleRights.Black))
	if len(castleString) == 0 {
		return "-"
	}

	return castleString
}

func CreateFEN(boardState board.BoardState) string {
	chessboard := boardState.Board
	configs := boardState.Configuration
	return createBoardPart(chessboard) + " " + createTurnPart(configs) + " " + createCastlingPart(configs)
}
