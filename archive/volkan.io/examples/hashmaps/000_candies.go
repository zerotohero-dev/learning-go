package main

import (
	"fmt"
)

func distributeCandies(candies []int) int {
	m := make(map[int]int)

	for _, i := range candies {
		_, ok := m[i]

		if ok {
			m[i]++
		} else {
			m[i] = 1
		}
	}

	if len(m) >= len(candies)/2 {
		return len(candies) / 2
	}

	return len(m)
}

func main() {
	candies := []int{1, 2, 2, 2, 2, 3, 3, 4, 4, 6, 6, 9}

	fmt.Println(
		distributeCandies(candies),
	)
}
