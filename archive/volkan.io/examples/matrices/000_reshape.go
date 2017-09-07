package main

import (
	"fmt"
)

func matrixReshape(nums [][]int, r int, c int) [][]int {
	n := len(nums)
	m := len(nums[0])

	if r*c != n*m {
		fmt.Println("Oh poop!")

		return nums
	}

	res := make([][]int, r)

	for i := 0; i < r; i++ {
		res[i] = make([]int, c)
	}

	for i := 0; i < r*c; i++ {
		res[i/c][i%c] = nums[i/m][i%m]
	}

	return res
}

func main() {
	nums := make([][]int, 2)
	nums[0] = []int{1, 2, 3, 4, 5, 6}
	nums[1] = []int{7, 8, 9, 10, 11, 12}

	fmt.Println(
		matrixReshape(nums, 3, 4),
	)
}
