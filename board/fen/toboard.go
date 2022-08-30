package fen

import (
	"errors"
	"fmt"
	"gochess/board"
	"gochess/board/pieces"
	"unicode"
)

func createNulls(number int) board.Board {
	board := board.Board{}
	for i := 0; i < number; i++ {
		board = append(board, nil)
	}
	return board
}

func createPiece(piece rune, position int) board.Movable {
	isWhite := unicode.IsUpper(piece)
	piece = unicode.ToLower(piece)
	switch piece {
	case 'r':
		return pieces.Rook{
			Piece: pieces.Piece{
				IsWhite: isWhite,
			},
		}
	case 'n':
		return pieces.Knight{
			Piece: pieces.Piece{
				IsWhite: isWhite,
			},
		}
	case 'b':
		return pieces.Bishop{
			Piece: pieces.Piece{
				IsWhite: isWhite,
			},
		}
	case 'k':
		return pieces.King{
			Piece: pieces.Piece{
				IsWhite: isWhite,
			},
		}
	case 'q':
		return pieces.Queen{
			Piece: pieces.Piece{
				IsWhite: isWhite,
			},
		}
	case 'p':
		return pieces.Pawn{
			Piece: pieces.Piece{
				IsWhite: isWhite,
			},
		}
	default:
		panic(fmt.Sprintf("This shouldn't have happened. No piece for %c", piece))
	}
}

func createPositions(boardString string) (board.Board, error) {
	chessboard := board.Board{}
	for _, str := range boardString {
		if str == '/' {
			continue
		}
		if '1' <= str && str <= '8' {
			chessboard = append(chessboard, createNulls(int(str-'0'))...)
			continue
		}

		chessboard = append(chessboard, createPiece(str, len(chessboard)))
	}

	if len(chessboard) != 64 {
		return chessboard, fmt.Errorf("expecting board size of 64. Found size %d", len(chessboard))
	}

	return chessboard, nil
}

func createConfigs(turn string, castling string) board.Configuration {
	config := board.Configuration{
		WhiteTurn: turn == "w",
	}

	for _, letter := range castling {
		switch letter {
		case 'k':
			config.CastleRights.Black.King = true
		case 'q':
			config.CastleRights.Black.Queen = true
		case 'K':
			config.CastleRights.White.King = true
		case 'Q':
			config.CastleRights.White.Queen = true
		}
	}

	return config
}

// Create a Board and a config object from a FEN String. The function takes only the first 3 parameters (board position, turn and castling rights)
// You can read about FEN Strings here: https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
func CreateBoard(fen FenStr) (*board.BoardState, error) {
	if !fen.Validate() {
		return nil, errors.New("there was something wrong with the FEN String")
	}

	boardFEN, turnFEN, castlingFEN := fen.Separate()

	chessboard, err := createPositions(boardFEN)
	if err != nil {
		return nil, err
	}

	configuration := createConfigs(turnFEN, castlingFEN)
	state := board.BoardState{
		Board:         chessboard,
		Configuration: configuration,
	}
	return &state, nil
}
