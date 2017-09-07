package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var literal = map[string]Vertex{
	"Bell Labs": Vertex{44, 22},
	"Google":    Vertex{12, 24},
}

var literal2 = map[string]Vertex{
	"Bell Labs": {442, 13},
	"Google":    {12, 432},
}

func main() {
	m = make(map[string]Vertex)

	m["Bell Labs"] = Vertex{50, 44}

	fmt.Println(m["Bell Labs"])

	fmt.Println(literal)
	fmt.Println(literal2)

	osman := make(map[string]int)

	osman["Answer"] = 42
	fmt.Println("The value:", osman["Answer"])

	osman["Answer"] = 48
	fmt.Println("The value:", osman["Answer"])

	delete(osman, "Answer")
	fmt.Println("The value:", osman["Answer"])

	v, ok := osman["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
