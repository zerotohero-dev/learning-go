package main

import (
	"fmt"
)

// Will return `2`.
func test() (i int) {
	defer func() { i++ }()

	return 1
}

func main() {

	// Deferred call’s arguments are immediately evaluated, but
	// the call is not executed until the function returns.
	defer fmt.Println("Hello")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("World")

	fmt.Println("test", test())
}
