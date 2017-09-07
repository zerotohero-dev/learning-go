package main

import (
	"sort"
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maximumProduct(nums []int) int {
	sort.Ints(nums)

	length := len(nums)

	return max(nums[length-1] * nums[length-2] * nums[length-3], nums[0] * nums[1] * nums[length-1])
}

func main() {
	fmt.Println(
		maximumProduct([]int{2,11,6,3}),
	)
}