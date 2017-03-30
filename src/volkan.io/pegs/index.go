package pegs

import (
    "fmt"
)

const N = 11 + 1

var board = []rune(`
...........
...........
....●●●....
....●●●....
..●●●●●●●..
..●●●○●●●..
..●●●●●●●..
....●●●....
....●●●....
...........
...........
`)

var center int

func init() {
    n := 0

    for pos, field := range board {
        if field == '○' {
            center = pos
            n++
        }
    }

    if n != 1 {
        center = -1
    }
}

var moves int

func move(pos, dir int) bool {
    moves++

    if board[pos] == '●' && board[pos+dir] == '●' && board[pos+2*dir] == '○' {
		board[pos] = '○'
		board[pos+dir] = '○'
		board[pos+2*dir] = '●'
		return true
	}

	return false
}

func unmove(pos, dir int) {
	board[pos] = '●'
	board[pos+dir] = '●'
	board[pos+2*dir] = '○'
}

func Solve() bool {
    var last, n int

    for pos, field := range board {
        if field == '●' {
            for _, dir := range [...]int{-1, -N, +1, +N} {
                if move(pos, dir) {
                    if Solve() {
                        unmove(pos, dir)
                        fmt.Println(string(board))
                        return true
                    }
                    unmove(pos, dir)
                }
            }
            last = pos
            n++
        }
    }

    if n == 1 && last == center {
        fmt.Println(string(board))
        fmt.Println("moves tried so far:", moves)
        return true
    }

    return false
}