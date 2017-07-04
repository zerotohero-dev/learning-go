package main

import "fmt"

func longestPalindrome(s string) int {
	set := make(map[rune]bool)
	count := 0

	for _, c := range s {
		_, ok := set[c]

		if ok {
			delete(set, c)
			count++
		} else {
			set[c] = true
		}
	}

	if len(set) > 0 {
		return count*2 + 1
	}

	return count*2
}

func main() {
	fmt.Println(
		longestPalindrome("doloremipsum"),
	)
}