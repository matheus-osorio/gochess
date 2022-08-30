package pieces

import (
	"gochess/board/fen"
	"testing"
)

func Test_FirstMove(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("8/p7/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[8]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 8)
	if len(moves) != 2 {
		t.Errorf("Pawn has 2 moves available on the tested move!")
	}
}

func Test_FowardMove(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("8/p7/p7/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[16]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 16)
	if len(moves) != 1 {
		t.Errorf("Pawn has 1 move available on the tested move!")
	}
}

func Test_CaptureMove(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("8/p7/pP6/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[8]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 8)
	if len(moves) != 1 {
		t.Errorf("Pawn has 1 moves available on the tested move!")
	}
}

func Test_NoMove(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("8/p7/Pp6/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[8]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 8)
	if len(moves) != 0 {
		t.Errorf("Pawn has NO moves available on the tested move!")
	}
}
