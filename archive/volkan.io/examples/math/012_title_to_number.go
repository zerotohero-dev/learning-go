package main

import "fmt"

func titleToNumber(s string) int {
	result := 0

	for i := 0; i < len(s); i++ {
		result = result * 26 + (int(s[i] - 'A') + 1)
	}

	return result
}

func main() {
	fmt.Println(
		titleToNumber("ABC"),
	)
}