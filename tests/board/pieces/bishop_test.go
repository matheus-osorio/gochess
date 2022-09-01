package pieces

import (
	"gochess/board/fen"
	"testing"
)

func Test_NoMovesDiag(t *testing.T) {
	state, _ := fen.CreateBoard("bnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[0]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 0)
	if len(moves) > 0 {
		t.Errorf("Expected 0 moves available. Got: %d", len(moves))
	}
}

func Test_1MovesDiag1(t *testing.T) {
	state, _ := fen.CreateBoard("bnbqkbnr/p1pppppp/2p5/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[0]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 0)
	if len(moves) != 1 {
		t.Errorf("Expected 1 move available. Got: %d", len(moves))
	}
}

func Test_CaptureDiag1(t *testing.T) {
	state, _ := fen.CreateBoard("bnbqkbnr/p1pppppp/2P5/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[0]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 0)
	if len(moves) != 2 {
		t.Errorf("Expected 2 move available. Got: %d", len(moves))
	}
}

func Test_NoMovesDiag2(t *testing.T) {
	state, _ := fen.CreateBoard("bnbqkbnb/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[7]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 7)
	if len(moves) > 0 {
		t.Errorf("Expected 0 moves available. Got: %d", len(moves))
	}
}

func Test_1MovesDiag2(t *testing.T) {
	state, _ := fen.CreateBoard("bnbqkbnb/pppppp1p/5ppp/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[7]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 7)
	if len(moves) != 1 {
		t.Errorf("Expected 0 moves available. Got: %d", len(moves))
	}
}

func Test_CaptureDiag2(t *testing.T) {
	state, _ := fen.CreateBoard("bnbqkbnb/pppppp1p/5Ppp/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	piece := state.Board[7]
	moves := piece.MovesAvailable(state.Board, state.Configuration, 7)
	if len(moves) != 2 {
		t.Errorf("Expected 0 moves available. Got: %d", len(moves))
	}
}
