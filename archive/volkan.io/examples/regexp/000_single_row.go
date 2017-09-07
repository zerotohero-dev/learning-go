package main

import (
	"fmt"
	"regexp"
)

var re = regexp.MustCompile(`(?i)^[qwertyuiop]*$|^[asdfghjkl]*$|^[zxcvbnm]*$`)

func findWords(words []string) []string {
	results := make([]string, 0)

	for _, str := range words {
		if re.MatchString(str) {
			results = append(results, str)
		}
	}

	return results
}

func main() {
	words := []string{"babel", "flag", "you", "toast"}

	fmt.Println(
		findWords(words),
	)
}
