package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale3(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v Vertex) Abs4() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale4(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale4(2)
	(&v).Scale4(12)

	var v1 = &v

	(&v).Scale4(44)
	(v1).Scale4(12)
	v.Scale4(11)

	v3 := Vertex{3, 4}
	fmt.Println(v3.Abs4())

	p3 := &Vertex{3, 4}
	fmt.Println(p3.Abs4())
	fmt.Println((*p3).Abs4())

}
