package main

import (
	"fmt"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune {
	return r + 1
}

type intIntFnType func(n int) int

func callFunc(f func(n int) int) int {
	return f(44) + f(66)
}

func callFunc2(f intIntFnType) int {
	return f(44) + f(66)
}

func main() {
	f := square

	fmt.Println(f(3))

	f = negative

	fmt.Println(f(3))

	fmt.Printf("%T\n", f)

	// f = product // Compile error due to signature mismatch

	fmt.Println(strings.Map(add1, "Dolorem"))
	fmt.Println(callFunc(square))
	fmt.Println(callFunc2(negative))
}
