package pieces

import (
	"gochess/board/fen"
	"testing"
)

func Test_TopLeftSquare(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("n7/pp6/1P6/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	knight := state.Board[0]
	moves := knight.MovesAvailable(state.Board, state.Configuration, 0)
	if len(moves) != 2 {
		t.Errorf("Top Left Knight has 2 moves when on top left")
	}

}

func Test_TopRightSquare(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("7n/pp6/1P6/8/8/8/PPPPPPPP/RNBQKBNR w KQkq")
	knight := state.Board[7]
	moves := knight.MovesAvailable(state.Board, state.Configuration, 7)
	if len(moves) != 2 {
		t.Errorf("Top Left Knight has 2 moves when on top left")
	}

}

func Test_BottomLeftSquare(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("7n/pp6/1P6/8/8/8/PPpPPPPP/NRBQKBNR w KQkq")
	knight := state.Board[56]
	moves := knight.MovesAvailable(state.Board, state.Configuration, 56)
	if len(moves) != 2 {
		t.Errorf("Top Left Knight has 2 moves when on top left")
	}
}

func Test_BottomRightSquare(t *testing.T) {
	t.Parallel()
	state, _ := fen.CreateBoard("7n/pp6/1P6/8/8/8/PPPPP1PP/RNBQKBRN w KQkq")
	knight := state.Board[63]
	moves := knight.MovesAvailable(state.Board, state.Configuration, 63)
	if len(moves) != 2 {
		t.Errorf("Top Left Knight has 2 moves when on top left")
	}
}
