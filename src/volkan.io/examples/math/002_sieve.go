package main

import (
	"fmt"
	"volkan.io/math/sieve"
)

func main() {
	ch := make(chan int)

	go sieve.Generate(ch)

	fmt.Println()

	for i := 0; i < 1000; i++ {
		prime := <-ch

		fmt.Println(prime)

		ch1 := make(chan int)

		go sieve.Filter(ch, ch1, prime)

		ch = ch1
	}
}
