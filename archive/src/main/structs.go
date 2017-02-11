package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{1, 2}
	v.X = 4

	fmt.Println(v.X)

	// Pointers to structs

	v = Vertex{1, 2}
	p := &v
	(*p).X = 1e9
	p.X = 11

	fmt.Println(v)
}
