package main

import (
	"fmt"
)

type IntBox struct {
	Value int
}

func (b IntBox) increment() IntBox {
	b.Value++

	return b
}

func (b *IntBox) incrementP() IntBox {
	b.Value++

	return *b
}

func main() {
	theBox := IntBox{42}

	box1 := theBox.increment()

	fmt.Println("theBox", theBox.Value)
	fmt.Println("box1 (val)", box1.Value)

	fmt.Println("________")

	box2 := theBox.incrementP()

	fmt.Println("theBox", theBox.Value)
	fmt.Println("box2 (ptr)", box2.Value)

}
