package main

import "sort"

func containsDuplicateNaive(nums []int) bool {
	length := len(nums)

	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}

func containsDuplicateSort(nums []int) bool {
	sort.Ints(nums)

	for i := range nums {
		if i > 0 {
			if nums[i-1] == nums[i] {
				return true
			}
		}
	}

	return false
}

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)

	for _, num := range nums {
		if set[num] == true {
			return true
		}

		set[num] = true
	}

	return false
}


func main() {
	containsDuplicateNaive([]int{42})
	containsDuplicateSort([]int{42})
	containsDuplicate([]int{42})
}


