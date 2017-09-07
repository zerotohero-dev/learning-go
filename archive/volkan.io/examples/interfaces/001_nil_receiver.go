package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")

		return
	}

	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T

	i = t

	// nil value of concrete tpye *T
	i.M()

	var j I

	// nil value of nil concrete type.
	fmt.Println(j)

	var k interface{}

	k = 42

	fmt.Println(k)

	k = "Hello"

	fmt.Println(k)
}
