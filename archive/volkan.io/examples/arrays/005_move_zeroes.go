package main

import "fmt"

func moveZeroes(nums []int) {
	if nums == nil {return}
	if len(nums) == 0 {return}

	insertPos := 0

	for _, num := range nums {
		if num != 0 {
			nums[insertPos] = num
			insertPos++
		}
	}

	for insertPos < len(nums) {
		nums[insertPos] = 0
		insertPos++
	}
}

func main() {
	nums := []int{0,0,1,3,0,4,0,5,12,44,0,25,33}
	moveZeroes(nums)
	fmt.Println(nums)
}