package main

import (
	"fmt"
)

func doReverse(begin int, end int, in []rune) []rune {
	for begin < end {
		in[begin] = in[begin] ^ in[end]
		in[end] = in[begin] ^ in[end]
		in[begin] = in[begin] ^ in[end]

		begin++
		end--
	}

	return in
}

func reverseRune(in []rune) string {
	in = doReverse(0, len(in)-1, in)

	return string(in)
}

func reverseString(s string) string {
	return reverseRune([]rune(s))
}

func main() {
	const phrase = "Go is rocking my world!"

	fmt.Println("+" + reverseString(phrase) + "+")
}
