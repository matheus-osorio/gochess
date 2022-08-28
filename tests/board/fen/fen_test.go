package fen

import (
	"gochess/board/fen"
	"testing"
)

func Test_Validate(t *testing.T) {
	t.Parallel()
	f := fen.FenStr("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	if !f.Validate() {
		t.Errorf("Regex classified a valid string as invalid!")
	}

	f = fen.FenStr("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBLKBNR w KQkq")
	if f.Validate() {
		t.Errorf("Regex could not detect error on a invalid FEN string")
	}
}

func Test_Separate(t *testing.T) {
	t.Parallel()
	f := fen.FenStr("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	first_part, second_part, third_part := f.Separate()

	if first_part != "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR" {
		t.Errorf("Error while spliting FEN's first part!")
	}

	if second_part != "w" {
		t.Errorf("Error while spliting FEN's second part!")
	}

	if third_part != "KQkq" {
		t.Errorf("Error while spliting FEN's third part!")
	}
}
