package main

import "fmt"

const boilingF = 212.0

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

// slice pointer map channel function

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9

	fmt.Printf("boling point = %g°F or %g°C\n", f, c)
	fmt.Printf("boling point = %g°F or %g°C\n", f, fToC(f))
}
