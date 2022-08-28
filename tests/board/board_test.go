package board

import (
	"gochess/board"
	"gochess/board/pieces"
	"reflect"
	"testing"
)

func Test_GetPlaceInBoard(t *testing.T) {
	t.Parallel()
	chessboard := board.Board{}
	expected_position := board.BoardPlace{
		Up:    4,
		Left:  0,
		Right: 7,
		Down:  3,
	}
	given_position := chessboard.GetPlaceInBoard(32)

	if !reflect.DeepEqual(expected_position, given_position) {
		t.Errorf("Expected position %v differs from given position %v", expected_position, given_position)
	}
}

func Test_GetCoordinateInfo(t *testing.T) {
	t.Parallel()
	chessboard := board.Board{nil, pieces.Bishop{Piece: pieces.Piece{IsWhite: true}}, pieces.Rook{Piece: pieces.Piece{IsWhite: false}}}
	if info := chessboard.GetCoordinateInfo(0); info != board.FREESPOT {
		t.Errorf("Expected type of coordinate to be FREESPOT (int value: %d), found %d", board.FREESPOT, info)
	}
	if info := chessboard.GetCoordinateInfo(1); info != board.WHITEPIECE {
		t.Errorf("Expected type of coordinate to be WHITEPIECE (int value: %d), found %d", board.WHITEPIECE, info)
	}
	if info := chessboard.GetCoordinateInfo(2); info != board.BLACKPIECE {
		t.Errorf("Expected type of coordinate to be BLACKPIECE (int value: %d), found %d", board.BLACKPIECE, info)
	}
}
