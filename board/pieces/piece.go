package pieces

import (
	"gochess/board"
	"math"
)

const LEFTMOVE = -1
const RIGHTMOVE = 1
const UPMOVE = -8
const DOWNMOVE = 8
const UPLEFTMOVE = UPMOVE + LEFTMOVE
const UPRIGHTMOVE = UPMOVE + RIGHTMOVE
const DOWNLEFTMOVE = DOWNMOVE + LEFTMOVE
const DOWNRIGHTMOVE = DOWNMOVE + RIGHTMOVE

const NORULE = 0
const MUSTCAPTURE = 100
const MUSTNOTCAPTURE = 200

type Piece struct {
	IsWhite bool
}

func (p Piece) IsWhitePiece() bool {
	return p.IsWhite
}

func (p Piece) checkMoves(chessboard board.Board, increment int, coordinate int, color int, moveRange int, moveRule int) (coordinates []int) {
	for i := 1; i <= moveRange; i++ {
		current := coordinate + i*increment
		squareStatus := chessboard.GetCoordinateInfo(current)
		if squareStatus == board.FREESPOT && moveRule != MUSTCAPTURE {
			coordinates = append(coordinates, current)
			continue
		} else if squareStatus != board.FREESPOT && squareStatus != color && moveRule != MUSTNOTCAPTURE {
			coordinates = append(coordinates, current)
			break
		}
		break
	}

	return
}

func (p Piece) leftMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)

	moveRange := int(math.Min(float64(pieceRange), float64(boardRange.Left)))
	return p.checkMoves(board, LEFTMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) rightMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)

	moveRange := int(math.Min(float64(pieceRange), float64(boardRange.Right)))
	return p.checkMoves(board, RIGHTMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) upMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)

	moveRange := int(math.Min(float64(pieceRange), float64(boardRange.Up)))
	return p.checkMoves(board, UPMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) downMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)

	moveRange := int(math.Min(float64(pieceRange), float64(boardRange.Down)))
	return p.checkMoves(board, DOWNMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) upLeftMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)
	minimalDiagonalMove := math.Min(float64(boardRange.Up), float64(boardRange.Left))
	moveRange := int(math.Min(float64(pieceRange), minimalDiagonalMove))
	return p.checkMoves(board, UPLEFTMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) upRightMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)
	minimalDiagonalMove := math.Min(float64(boardRange.Up), float64(boardRange.Right))
	moveRange := int(math.Min(float64(pieceRange), minimalDiagonalMove))
	return p.checkMoves(board, UPRIGHTMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) downLeftMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)
	minimalDiagonalMove := math.Min(float64(boardRange.Down), float64(boardRange.Left))
	moveRange := int(math.Min(float64(pieceRange), minimalDiagonalMove))
	return p.checkMoves(board, DOWNLEFTMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) downRightMoves(board board.Board, pieceRange int, moveRules int, coordinate int) []int {
	pieceColor := board.GetCoordinateInfo(coordinate)
	boardRange := board.GetPlaceInBoard(coordinate)
	minimalDiagonalMove := math.Min(float64(boardRange.Down), float64(boardRange.Right))
	moveRange := int(math.Min(float64(pieceRange), minimalDiagonalMove))
	return p.checkMoves(board, DOWNRIGHTMOVE, coordinate, pieceColor, moveRange, moveRules)
}

func (p Piece) CreateBoardStates(coordinates []int, chessboard board.Board, configs board.Configuration, name string, origin int) (states []board.BoardState) {
	for _, coordinate := range coordinates {

		state := board.BoardState{}
		state.CreateBoardState(origin, coordinate, chessboard, configs, name, p.IsWhite)
		states = append(states, state)
	}

	return states
}
