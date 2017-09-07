package main

import (
	"fmt"
	"sort"
)

func min(a int, b int) int {
	if a < b {
		return a
	}

	return b
}

func minMoves(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	m := nums[0]

	for _, n := range nums {
		m = min(m, n)
	}

	result := 0

	for _, n := range nums {
		result += n - m
	}

	return result
}

func minMovesDp(nums[] int) int {
	sort.Ints(nums)

	moves := 0

	for i, num := range nums {
		if i > 0 {
			diff := (moves + num) - nums[i-1]
			nums[i] += moves
			moves += diff
		}
	}

	return moves
}

func main() {
	fmt.Println(
		minMoves([]int{}),
		minMovesDp([]int{}),
	)
}