package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func findDisappearedNumbers(nums []int) []int {
	collected := make([]int, 0)

	for _, item := range nums {
		nums[abs(item)-1] = -1 * abs(nums[abs(item)-1])
	}

	for index, item := range nums {
		if item > 0 {
			collected = append(collected, index+1)
		}
	}

	return collected
}

func main() {
	fmt.Println(
		findDisappearedNumbers([]int{9,8,8,6,5,4,4,2,1}),
	)
}