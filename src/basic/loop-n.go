package main

import "fmt"

func main() {
	a := 0
	const N = 10
	counter := 0

	// (N(N+1))/2


	for i := 0; i < N; i++ {
		for j := N; j > i; j-- {
			a = a + i + j
			counter += 1

			fmt.Println(counter, a)
		}
	}
}
