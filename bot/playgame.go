package bot

import (
	"fmt"
	"gochess/board"
	"reflect"
	"sync"
	"time"
)

func getMoves(state board.BoardState, depth int) MoveTree {
	configs := state.Configuration
	chessboard := state.Board
	whiteTurn := configs.WhiteTurn
	var states []board.BoardState

	for coordinate, piece := range chessboard {
		if piece == nil {
			continue
		}
		if piece.IsWhitePiece() != whiteTurn {
			continue
		}
		moves := piece.MovesAvailable(chessboard, configs, coordinate)
		for _, m := range moves {
			if len(m.Board) != 64 {
				fmt.Println(reflect.TypeOf(piece).Name())
			}
		}
		states = append(states, moves...)
	}
	origin := MoveTree{
		current: state,
		depth:   depth,
	}

	for _, state := range states {
		nextState := MoveTree{
			current: state,
			depth:   depth + 1,
		}
		origin.next = append(origin.next, &nextState)
	}
	return origin
}

func putMovesOnQueue(moves chan *MoveTree, moveList []*MoveTree) {
	for _, move := range moveList {
		moves <- move
	}
}

func iterateMoves(moves chan *MoveTree, depth int, wg *sync.WaitGroup) {
	go func(moves chan *MoveTree, depth int) {
		defer wg.Done()
		for {
			var move *MoveTree
			select {
			case res := <-moves:
				move = res

			case <-time.After(100 * time.Millisecond):
				return
			}
			if move.depth >= depth {
				continue
			}
			replacement := getMoves(move.current, move.depth)
			move.next = replacement.next

			go putMovesOnQueue(moves, replacement.next)
		}

	}(moves, depth)
}

func calculateNextMoves(states *MoveTree, depth int, threads int) {
	wg := sync.WaitGroup{}
	wg.Add(threads)
	moves := make(chan *MoveTree, len(states.next))
	for _, move := range states.next {
		moves <- move
	}

	for i := 0; i < threads; i++ {
		iterateMoves(moves, depth, &wg)
	}

	wg.Wait()
}

func Play(state board.BoardState, depth int, threads int) board.BoardState {

	states := getMoves(state, 0)
	calculateNextMoves(&states, depth, threads)
	return states.Calculate().current
}
