package main

import (
	"math/rand"
	"fmt"
)

func main() {
	var a int = 0
	var b int = 0

	const N = 100
	const M = 300

	for i := 0; i < N; i++ {
		a += rand.Int()
	}

	for j := 0; j < M; j++ {
		b += rand.Int()
	}

	fmt.Println(a, b)
}
