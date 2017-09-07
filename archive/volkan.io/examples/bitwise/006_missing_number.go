package main

import "fmt"

func missingNumber(nums []int) int {
	x := 0
	index := 0

	for i, num := range nums {
		x = x ^ i ^ num
		index = i
	}

	return x ^ (index + 1)
}

func main() {
	fmt.Println(
		missingNumber([]int{0,1,3}),
	)
}