package main

import (
	"fmt"
)

func reverse(begin int, end int, in []rune) []rune {
	for begin < end {
		in[begin] = in[begin] ^ in[end]
		in[end] = in[begin] ^ in[end]
		in[begin] = in[begin] ^ in[end]

		begin++
		end--
	}

	return in
}

func reverseWordsFromRune(in []rune) string {
	index := 0

	for i := 0; i < len(in); i++ {
		if in[i] == ' ' {
			in = reverse(index, i-1, in)
			index = i + 1
		}
	}

	return string(in)
}

func reverseWords(s string) string {
	result := reverseWordsFromRune([]rune(s + " "))

	return result[:len(result)-1]
}

func main() {
	const phrase = "Go is rocking my world!"

	fmt.Println("+" + reverseWords(phrase) + "+")
}
