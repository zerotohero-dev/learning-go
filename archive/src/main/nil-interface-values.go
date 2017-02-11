package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I;

	fmt.Printf("(%v, %T)", i, i)

	// error: nil pointer dereference:
	// i.M()
}
