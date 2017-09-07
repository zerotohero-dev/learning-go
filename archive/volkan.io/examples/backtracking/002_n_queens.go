package main

import (
	"fmt"
	"math"
)

type Queen struct {
    Row int
    Col int
}

func valid(row int, currentCol int, currentQueen int, queens *[8]Queen) bool {
    for q := 0; q < currentQueen; q++ {
        currentQ := (*queens)[q]

        if row == currentQ.Row ||
            currentCol == currentQ.Col ||
            math.Abs(float64(currentCol-currentQ.Col)) == math.Abs(float64(row - currentQ.Row)) {
            return false
        }
    }

    return true
}

func nQ(currentQueen int, currentCol int, rows int, queens *[8]Queen) bool {
    if currentQueen >= rows {
        return true
    }

    found := false
    row := 0

    for row < rows && !found {
        if valid(row, currentCol, currentQueen, queens) {
            (*queens)[currentQueen].Row = row
            (*queens)[currentQueen].Col = currentCol

            found = nQ(currentQueen+1, currentCol+1, rows, queens)
        }

        row++
    }

    return found
}

func main() {
    rows := 8
    queens := [8]Queen{}

	nQ(0, 0, rows, &queens)

	fmt.Println(queens)
}
