package main

import (
	"fmt"
	"volkan.io/geometry"
)

func main() {
	p := geometry.Point{1, 2}
	q := geometry.Point{3, 4}

	fmt.Println(geometry.Distance(q, p))
	fmt.Println(p.Distance(q))
}
