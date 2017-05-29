package main

import (
	"fmt"
	"strings"
)

func generatePossibleNextMoves(s string) []string {
	result := make([]string, 0)

	i := -1

	for {
		i++

		index := strings.Index(s[i:], "++")

		if index == -1 {
			break
		}

		i += index

		result = append(result, s[0:i] + "--" + s[i+2:])
	}

	return result
}

func main() {
	fmt.Println(generatePossibleNextMoves("++--+++++---+"))
}
