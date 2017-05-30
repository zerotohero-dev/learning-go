package main

import "fmt"

func detectCapitalUse(word string) bool {
	count := 0

	for _, c := range word {
		if int('Z') - int(c) >= 0 {
			count++
		}
	}

	if count == 0 {
		return true
	}

	if count == len(word) {
		return true
	}

	if count == 1 && int('Z') - int(word[0]) >= 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(
		detectCapitalUse("ffffffffffffffffffffF"),
	)
}