package main

import (
	"fmt"
)

// var (
//	message = "The answer to life is %d\n"
//	answer = 42
// )

const (
	message = "The answer to life is %d %d %d\n"
	answer  = 42
	answer1 = iota * 2
	answer2
)

func main() {
	// var message string
	// message = "Hello World!\n"

	// message := "Hello, World!\n"

	fmt.Printf(message, answer, answer1, answer2)

	var pi float64 = 3.14

	pi32 := float32(3.15)

	fmt.Printf("Value %f %f\n", pi, pi32)

	bello := true

	// %t is for the “truth” value
	fmt.Printf("\nBello %t\n", bello)
}
