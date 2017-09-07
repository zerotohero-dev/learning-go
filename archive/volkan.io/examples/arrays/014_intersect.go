package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	integers := make(map[int]int)
	result := make([]int, 0)

	for _, num := range nums1 {
		_, ok := integers[num]

		if ok {
			integers[num] += 1
		} else {
			integers[num] = 1
		}
	}

	for _, num := range nums2 {
		val, ok := integers[num]

		if ok && val > 0 {
			result = append(result, num)

			integers[num] -= 1
		}
	}

	return result
}

func main() {
	fmt.Println(
		intersect([]int{2,2,4,4,2,3,5,6}, []int{1,2,3,4,5,5,6,6,6,2,2,1}),
	)
}