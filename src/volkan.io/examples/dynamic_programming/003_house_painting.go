package main

import "fmt"

func min(a int, b int) int {
	if a > b {return b}

	return a
}

func minCost(costs [][]int) int {
	if costs == nil {return 0}
	if len(costs) == 0 {return 0}

	for i := 1; i < len(costs); i++ {
		costs[i][0] += min(costs[i-1][1], costs[i-1][2])
		costs[i][1] += min(costs[i-1][0], costs[i-1][2])
		costs[i][2] += min(costs[i-1][1], costs[i-1][0])
	}

	n := len(costs) - 1

	return min(min(costs[n][0], costs[n][1]), costs[n][2])
}

func main() {
	values := [][]int{}

	// These are the first two rows.
	row1 := []int{20, 18, 4}
	row2 := []int{9, 9, 10}

	// Append each row to the two-dimensional slice.
	values = append(values, row1)
	values = append(values, row2)

	fmt.Println(
		minCost(values),
	)
}
