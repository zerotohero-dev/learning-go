package main

import (
	"fmt"
	"sort"
)

func intersectionWithSets(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	intersect := make(map[int]bool)

	for _, num := range nums1 {
		set[num] = true
	}

	for _, num := range nums2 {
		if set[num] {
			intersect[num] = true
		}
	}

	result := make([]int, len(intersect))

	i := 0
	for num := range intersect {
		result[i] = num
		i++
	}

	return result
}

func intersectionTwoPointers(nums1 []int, nums2 []int) []int {
	sort.Ints(nums1)
	sort.Ints(nums2)

	i := 0
	j := 0
	set := make(map[int]bool)

	for i < len(nums1) && j < len(nums2) {
		if nums1[i] < nums2[j] {
			i++
		} else if nums1[i] > nums2[j] {
			j++
		} else {
			set[nums1[i]] = true
			i++
			j++
		}
	}

	result := make([]int, len(set))

	i = 0
	for num := range set {
		result[i] = num
		i++
	}

	return result
}

func search(haystack []int, needle int) bool {
	low := 0
	high := len(haystack) - 1

	for low <= high {
		mid := low + (high-low)/2

		if haystack[mid] == needle {
			return true
		}

		if haystack[mid] > needle {
			high = mid -1
		} else if haystack[mid] < needle {
			low = mid + 1
		}
	}

	return false
}

func intersectUsingBinarySearch(nums1 []int, nums2 []int) []int {
	set := make(map[int]bool)
	sort.Ints(nums2)

	for _, num := range nums1 {
		if search(nums2, num) {
			set[num] = true
		}
	}

	result := make([]int, len(set))
	i := 0
	for num := range set {
		result[i] = num
		i++
	}

	return result
}

func main() {
	nums1 := []int{1,2,4,6,1,33,42}
	nums2 := []int{4,4,23,62,52,42,12,1}
	fmt.Println(
		intersectionWithSets(nums1, nums2),
		intersectionTwoPointers(nums1, nums2),
		intersectUsingBinarySearch(nums1, nums2),
	)
}