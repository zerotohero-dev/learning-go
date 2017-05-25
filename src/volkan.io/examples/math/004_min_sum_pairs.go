package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {

	// Sorting is in-place; it changes the given slice, it does not return a new one.
	sort.Ints(nums)

	result := 0

	for index, number := range nums {
		if index%2 == 0 {
			result += number
		}
	}

	return result
}

func main() {
	nums := []int{8, 2, 3, 4, 5, 6, 7, 1}

	fmt.Println(arrayPairSum(nums))
}
