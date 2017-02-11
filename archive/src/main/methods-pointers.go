package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs5() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale5(f float64) {
	v.X = f * v.X
	v.Y = f * v.Y
}

func main() {
	v := Vertex{X: 3, Y: 4}
	v.Scale5(10)
	fmt.Println(v.Abs5())
}
