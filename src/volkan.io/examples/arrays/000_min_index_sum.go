package main

import (
	"fmt"
	"math"
)

// const MaxUint = ^uint(0)
// const MinUint = 0
// const MaxInt = int(MaxUint >> 1)
// const MinInt = -MaxInt - 1

func findMinIndexSum(list1 []string, list2 []string) []string {
	minSum := math.MaxInt32

	indexes := make(map[string]int)

	for index, item := range list1 {
		indexes[item] = index
	}

	result := make([]string, 0)

	for index, item := range list2 {
		j, ok := indexes[item]

		if ok {
			if index + j <= minSum {
				if index + j < minSum {
					result = result[0:0]
					minSum = index + j
				}

				result = append(result, item)
			}
		}
	}

	return result
}

func main() {
	list1 := []string{"lorem", "dolor", "ahmet", "ipsum"}
	list2 := []string{"sit", "ahmet","dolor", "bingo"}

	fmt.Println(findMinIndexSum(list1, list2))
}