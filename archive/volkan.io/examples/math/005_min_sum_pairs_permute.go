package main

import (
	"fmt"
)

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func permute(nums []int, l int, maxSum int) int {
	if l == len(nums)-1 {
		sum := 0

		for i := 0; i < len(nums)/2; i++ {
			sum += Min(nums[i], nums[(len(nums)/2)+i])
		}

		return Max(maxSum, sum)
	}

	result := 0
	for i := l; i < len(nums); i++ {
		swap(nums, i, l)
		result = permute(nums, l+1, result)
		swap(nums, i, l)
	}

	return result
}

func swap(nums []int, i int, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

func main() {
	fmt.Println(
		permute([]int{1, 2, 3, 4}, 0, 0),
	)
}
