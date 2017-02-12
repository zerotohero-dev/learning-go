package main

import "fmt"

func main() {

	// The below is identical to:
	// var x int;p = &x;
	p := new(int)

	fmt.Println(*p)

	*p = 2

	fmt.Println(*p)
}
