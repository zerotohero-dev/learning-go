package main

import (
	"sort"
	"fmt"
)

type ByFirstItem [][]int

func (a ByFirstItem) Len() int {
	return len(a)
}

func (a ByFirstItem) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByFirstItem) Less(i, j int) bool {
	return a[i][0] > a[j][0]
}

func findRelativeRanks(nums []int) []string {
	pairs := make([][]int, len(nums))
	for i := 0; i < len(nums); i++ {
		pairs[i] = make([]int, 2)
		pairs[i][0] = nums[i]
		pairs[i][1] = i
	}

	if len(pairs) > 1 {
		sort.Sort(ByFirstItem(pairs))
	}

	result := make([]string, len(nums))

	fmt.Println(pairs)

	for i := 0; i < len(nums); i++ {
		if i == 0 {
			result[pairs[i][1]] = "Gold Medal"
		} else if i == 1 {
			result[pairs[i][1]] = "Silver Medal"
		} else if  i == 2 {
			result[pairs[i][1]] = "Bronze Medal"
		} else {
			result[pairs[i][1]] = fmt.Sprintf("%v", i+1)
		}
	}

	return result
}

func main() {
	fmt.Println(
		findRelativeRanks([]int{10,3,8,9,4}),
	)
}