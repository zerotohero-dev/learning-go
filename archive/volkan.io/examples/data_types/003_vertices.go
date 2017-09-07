package main

import (
	"fmt"
	"volkan.io/math/geo/vertex"
)

func main() {
	v := vertex.New(7, 24)

	fmt.Println(v.Abs())

	v.Scale(6)

	fmt.Println(v)

	v1 := &v

	v1.Scale(0.2)

	// In general, all methods on a given type should have either value or pointer receivers,
	// but not a mixture of both.

	fmt.Println(v)
	fmt.Println(v.Abs())
}
