package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	maxSoFar := 0
	globalMax := 0

	for _, item := range nums {
		if item == 1 {
			maxSoFar++
		} else {
			maxSoFar = 0
		}

		if globalMax < maxSoFar {
			globalMax = maxSoFar
		}
	}

	return globalMax
}

func main() {
	nums := []int{1,1,1,1,0,0,0,1,1,0,1,1,1,0,0,0,1,1,1,1,1,0,1}

	fmt.Println(
		findMaxConsecutiveOnes(nums),
	)
}