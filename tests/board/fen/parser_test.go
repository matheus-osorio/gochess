package fen

import (
	"gochess/board"
	"gochess/board/fen"
	"gochess/board/pieces"
	"reflect"
	"testing"
)

func Test_CreateBoard(t *testing.T) {
	chessboard, configs, err := fen.CreateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBKQBNR w KQkq")
	if err != nil {
		t.Error(err)
	}

	expectedOrder := board.Board{
		pieces.Rook{},
		pieces.Knight{},
		pieces.Bishop{},
		pieces.Queen{},
		pieces.King{},
		pieces.Bishop{},
		pieces.Knight{},
		pieces.Rook{},
		pieces.Pawn{},
		pieces.Pawn{},
		pieces.Pawn{},
		pieces.Pawn{},
		pieces.Pawn{},
		pieces.Pawn{},
		pieces.Pawn{},
		pieces.Pawn{},
	}

	whitePart := (*chessboard)[:16]
	blackPart := (*chessboard)[48:]

	for i := 0; i < 16; i++ {
		whiteReflection := reflect.TypeOf(whitePart[i])
		blackReflection := reflect.TypeOf(blackPart[15-i])
		if whiteReflection != blackReflection {
			t.Errorf("Types from white and black should me mirroed. White: %s, Black: %s", whiteReflection.Name(), blackReflection.Name())
		}

		expectedReflection := reflect.TypeOf(expectedOrder[i])
		if whiteReflection != expectedReflection {
			t.Errorf("Type is different than what was expected in model. Expected: %s, Got: %s", expectedReflection.Name(), whiteReflection.Name())
		}
	}

	if !(configs.CastleRights.Black.King && configs.CastleRights.Black.Queen) {
		t.Errorf("Castle Rights not being correctly inserted for Blacks")
	}

	if !(configs.CastleRights.White.King && configs.CastleRights.White.Queen) {
		t.Errorf("Castle Rights not being correctly inserted for Whites")
	}
}
