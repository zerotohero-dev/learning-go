package main

import "fmt"

// See also
// https://en.wikipedia.org/wiki/Boyer–Moore_majority_vote_algorithm
// https://www.quora.com/What-is-the-proof-of-correctness-of-Moores-voting-algorithm
func majorityElement(nums []int) int {
	major := nums[0]
	count := 1

	for i, num := range nums {
		if i > 0 {
			if count == 0 {
				count++
				major = num
			} else if major == num {
				count++
			} else {
				count--
			}
		}
	}

	return major
}

func main() {
	fmt.Println(
		majorityElement([]int{2,4,2,2,2,3,5,6,2,9,2,11}),
	)
}