package main

import "fmt"

const MaxUint = ^uint(0)
//const MinUint = 0
const MaxInt = int(MaxUint >> 1)
//const MinInt = -MaxInt - 1

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxCount(m int, n int, ops [][]int) int {
	if ops == nil {
		return m*n
	}

	if len(ops) == 0 {
		return m*n
	}

	row := MaxInt
	col := MaxInt

	for _, op := range ops {
		row = min(row, op[0])
		col = min(col, op[1])
	}

	return row * col
}

func main() {
	fmt.Println(
		maxCount(10, 10, nil),
	)
}