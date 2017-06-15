package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	indices := make([]int, 2)

	if numbers == nil {
		return indices
	}

	if len(numbers) < 2 {
		return indices
	}

	left := 0
	right := len(numbers) - 1

	for left < right {
		v := numbers[left] + numbers[right]

		if v == target {
			indices[0] = left + 1
			indices[1] = right + 1

			break
		} else if v > target {
			right--
		} else {
			left++
		}
	}

	return indices
}

func main() {
	fmt.Println(
		twoSum([]int{1,2,3,4,5,6,7,8,10,22}, 27),
	)
}
