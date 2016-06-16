package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	a := 2
	b := 4

	c := (a*b + a*b)

	v := Vertex{2, 3}

	fmt.Println(c)
	fmt.Println(v.Abs())
}
