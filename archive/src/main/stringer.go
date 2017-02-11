package main

import "fmt"

// type Stringer interface { String() string }

type Person struct {
	Name string
	Age int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beblerox", 9001}

	fmt.Println(a, z)
}
