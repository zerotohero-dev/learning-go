package main

import (
	"fmt"
)

type Board [][]int

func (b *Board) visited(c Coordinate) bool {
	return (*b)[c.X][c.Y] != -1
}

func initializeBoard(n int) Board {
	board := make(Board, n)

	for i := 0; i < len(board); i++ {
		board[i] = make([]int, n)
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			board[i][j] = -1
		}
	}

	return board
}

type Coordinate struct {
	X, Y int
}

func (c *Coordinate) stillInTheBoard(edgeSize int) bool {
	return c.X >= 0 && c.Y >= 0 && c.X < edgeSize && c.Y < edgeSize
}

type Next struct {
	X, Y  [8]int
	Count int
}

var nextCoords = Next{
	X:     [8]int{2, 1, -1, -2, -2, -1, 1, 2},
	Y:     [8]int{1, 2, 2, 1, -1, -2, -2, -1},
	Count: 8,
}

func (*Next) point(c Coordinate, i int) Coordinate {
	return Coordinate{X: nextCoords.X[i], Y: nextCoords.Y[i]}
}

func tour(coord Coordinate, moveNumber, edgeSize int, board *Board) bool {
	boardSize := edgeSize * edgeSize

	if moveNumber == boardSize {
		fmt.Println("Move number reached the board size", moveNumber, boardSize)

		return true
	}

	for i := 0; i < nextCoords.Count; i++ {
		nextCoord := Coordinate{X: coord.X + nextCoords.X[i], Y: coord.Y + nextCoords.Y[i]}

		if nextCoord.stillInTheBoard(edgeSize) && !board.visited(nextCoord) {
			(*board)[nextCoord.X][nextCoord.Y] = moveNumber

			nextMoveNumber := moveNumber + 1

			if tour(nextCoord, nextMoveNumber, edgeSize, board) {
				return true
			} else {

				// backtrack:
				(*board)[nextCoord.X][nextCoord.Y] = -1
			}

		}
	}

	return false
}

func main() {
	edgeSize := 8

	board := initializeBoard(edgeSize)

	board[0][0] = 0

	result := tour(Coordinate{0, 0}, 1, edgeSize, &board)

	if result {
		fmt.Println("Solution found.")
		fmt.Println(board)

		return
	}

	fmt.Println("Solution does not exist.")
}
