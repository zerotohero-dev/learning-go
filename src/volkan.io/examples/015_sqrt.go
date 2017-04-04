package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1000.0

	for i := 0; i < 1000; i++ {
		z = z - ((z*z - x) / (2 * z))
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}
