package board

import (
	"gochess/board/fen"
	"testing"
)

func Test_CalculatePoints(t *testing.T) {
	t.Parallel()
	state, err := fen.CreateBoard("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	if err != nil {
		t.Error(err)
	}

	if state.CalculatePoints() != 0 {
		t.Error("Even position is not being calculated properly")
	}
}
