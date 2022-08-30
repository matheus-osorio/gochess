package fen

import (
	"gochess/board"
	"gochess/board/fen"
	"gochess/board/pieces"
	"reflect"
	"testing"
)

func Test_CreateBoard(t *testing.T) {
	state, err := fen.CreateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBKQBNR w KQkq")
	if err != nil {
		t.Error(err)
	}

	chessboard, configs := state.Board, state.Configuration

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

	whitePart := (chessboard)[:16]
	blackPart := (chessboard)[48:]

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

func Test_CreateFEN(t *testing.T) {
	t.Parallel()
	original := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq"
	state, err := fen.CreateBoard(fen.FenStr(original))
	if err != nil {
		t.Error(err)
	}

	newFen := fen.CreateFEN(*state)
	if original != newFen {
		t.Errorf("The original FEN is different from the new one. Original: %s, New: %s", original, newFen)
	}
}
