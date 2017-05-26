package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`(?i)[qwertyuiop]*|[asdfghjkl]*|[zxcvbnm]*`)

func findWords(words []string) []string {
	result := re.MatchString("flag")

	fmt.Println(result)

	// // 	if match {
	// // 		return ["a"]
	// // 	}

	// 	return string[]{"b"}

	return words
}

func main() {
	words := []string{"babel", "flag", "you", "toast"}

	fmt.Println(
		findWords(words),
	)
}
