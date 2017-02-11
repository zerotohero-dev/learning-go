package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	if x <= 0 {return 0.0}
	if x == 1 {return 1.0}

	result := x

	for i := 0; i < 10; i++ {
		result = result - (math.Pow(result, 2)  - x)/(2 * result)
	}

	return result
}

func main() {
	fmt.Println(Sqrt(2), math.Sqrt(2))
}
