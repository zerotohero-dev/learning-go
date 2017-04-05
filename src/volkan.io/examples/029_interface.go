package main

import (
	"fmt"
	"math"
	"volkan.io/math/geo/vertex"
	. "volkan.io/math/interfaces"
	"volkan.io/reflect"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var a Abser

	v := vertex.New(2, 3)

	// a = v
	a = &v

	fmt.Println(v.Abs3())
	fmt.Println(a.Abs3())

	var i I = T{"Hello"}
	reflect.Describe(i)
	i.M()

	i = F(math.Pi)
	reflect.Describe(i)
	i.M()
}
