package main

import (
	"fmt"
)

type Point struct{ X, Y int }

type Path []Point

func main() {
	fmt.Printf("%T\n", make(Path, 10, 10))
	//fmt.Printf("%T\n", Path)

	p := make(Path, 5, 10)

	p[0] = Point{3, 4}

	fmt.Printf("%T\n", p)
	fmt.Printf("%v\n", p)
}
