package main

import "fmt"

func bitCount(num int) int {
	count := 0

	for num > 0 {
		count += num & 1

		num >>= 1
	}

	return count
}

func readBinaryWatch(num int) []string {
	times := make([]string, 0)

	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bitCount(h * 64 + m) == num {
				times = append(times, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}

	return times
}

func main() {
	fmt.Println(
		readBinaryWatch(8),
	)
}