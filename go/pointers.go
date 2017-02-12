package main

import "fmt"

func main() {
	x := 1

	// `p` points to `x`
	p := &x

	// "1"
	fmt.Println(*p)

	*p = 2

	// "2"
	fmt.Println(x)

	{
		var x, y int

		fmt.Println(&x == &x, &x == &y, &x == nil)
	}
}

func f() *int {
	v := 1

	// It is perfectly safe for a function to return
	// an address of a local variable.
	// This local variable will not be `gc`ed even after
	// the function exits.
	return &v
}
