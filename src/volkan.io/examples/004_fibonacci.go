package main

import (
	"fmt"
	"volkan.io/math/fibonacci"
)

func main() {
	fib := fibonacci.Closure()

	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())
	fmt.Println(fib())

	/*
		    Output:

			1
			1
			2
			3
			5
			8
			13
			21
			34
			55
			89
			144
			233
	*/
}
