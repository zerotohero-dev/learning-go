package main

import "fmt"

func main() {
	a := 0
	b := 0
	N := 100

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			a += j
		}
	}

	for k := 0; k < N; k++ {
		b += k
	}

	fmt.Println(a, b, N)
}
