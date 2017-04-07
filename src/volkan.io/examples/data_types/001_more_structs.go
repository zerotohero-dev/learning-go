package main

import (
	"fmt"
)

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	pv = &Vertex{X: 1, Y: 2}
)

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{2, 4}

	fmt.Println(v.X)

	v.X = 102

	// p is a struct pointer:
	p := &v

	(*p).X = 99

	// This one is a syntactic sugar for the above expression:
	p.X = 42

	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(v)

	fmt.Println(v1, v2, v3, pv)
}
