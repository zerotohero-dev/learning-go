package main

import (
	"fmt"
	"math"
	"volkan.io/extend/floats"
)

func main() {
	f := floats.FloatX(-math.Sqrt2)

	fmt.Println(f.Abs())
}
