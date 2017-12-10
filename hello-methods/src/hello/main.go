package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

// Value receiver
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	// Go automatically converts the value to pointer
	fmt.Println("Area:", r.area())
	fmt.Println("Perim:", r.perim())

	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Perim:", rp.perim())

	// Generally with structs you use methods with pointer receivers since
	// you generally intend to mutate the original struct.
	//
	// A rule of thumb is to NOT have mixed value/pointer receiver methods
	// (as in this example): Either all methods should have value receivers,
	// or all methods should have pointer receivers.
}
