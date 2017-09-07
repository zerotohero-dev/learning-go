package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	fmt.Println()
	for i := 0; i < len(board); i++ {
		fmt.Printf("\t%s\n", strings.Join(board[i], " "))
	}
	fmt.Println()
}
