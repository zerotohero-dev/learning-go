package main

import "fmt"

func islandPerimeter(grid [][]int) int {
	cells := 0
	neighbors := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				cells++

				if i < len(grid) - 1 && grid[i+1][j] == 1 { neighbors++ }

				if j < len(grid[i]) - 1 && grid[i][j+1] == 1  { neighbors++ }
			}
		}
	}

	return 4 * cells - 2 * neighbors
}

func main() {
	area := [][]int{
		{0, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 1, 1},
		{0, 1, 0, 0},
		{1, 1, 1, 0},
	}

	fmt.Println(islandPerimeter(area))
}