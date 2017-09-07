package main

import "fmt"

func main() {
	// Infinite loop:
	//
	// for {
	// 	fmt.Println("Hello")
	// }

	counter := 0

	for counter < 10 {
		fmt.Printf("Hello World\n")

		counter++
	}

	for c, d := 0, 1; c < 10; c, d = c+1, d+1 {
		fmt.Printf("Hello %d %d\n", c, d)
	}
}