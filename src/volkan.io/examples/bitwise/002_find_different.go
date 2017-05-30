package main

import "fmt"

func singleNumber(nums []int) int {
	result := 0

	for _, num := range nums {
		result ^= num
	}

	return result
}

func main() {
	fmt.Println(
		singleNumber([]int{1,2,3,3,2,1,4,6,5,5,4,6,42,7,7,9,9,13,13}),
	)
}