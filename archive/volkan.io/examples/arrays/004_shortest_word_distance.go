package main

import "fmt"

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func minimum(x int, y int) int {
	if x < y {
		return x
	}

	return y
}

func shortestDistance(words []string, word1 string, word2 string) int {
	min := -1
	ptr1 := -1
	ptr2 := -1

	for index, item := range words {
		if item == word1 {
			ptr1 = index
		}

		if item == word2 {
			ptr2 = index
		}

		if ptr1 != -1 && ptr2 != -1 {
			if min == -1 {
				min = abs(ptr2-ptr1)
			} else {
				min = minimum(min, abs(ptr2-ptr1))
			}
		}
	}

	return min
}

func main() {
	fmt.Println(
		shortestDistance(
			[]string{"lorem", "dolor"}, "lorem", "dolor",
		),
	)
}
