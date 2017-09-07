package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	stack := make([]int, 0)

	for _, number := range nums {
		for len(stack) > 0 && stack[len(stack)-1] < number {
			popped := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			m[popped] = number
		}
		stack = append(stack, number)
	}

	for i, foundNumber := range findNums {
		mapped, ok := m[foundNumber]

		if !ok {
			mapped = -1
		}


		findNums[i]	= mapped
	}

	return findNums
}

func main() {
	fmt.Println(
		nextGreaterElement([]int{2,3,4,56,22}, []int{11,22,33,44,55}),
	)
}