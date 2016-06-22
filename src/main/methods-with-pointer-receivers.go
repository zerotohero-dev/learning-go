package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale6(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs6() float64 {
	return math.Sqrt(v.X * v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}

	fmt.Printf("Before %+v, Abs: %v\n", v, v.Abs6())

	v.Scale6(50)

	fmt.Printf("After %+v, Abs: %v\n", v, v.Abs6())
}
