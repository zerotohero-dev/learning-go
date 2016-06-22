package main

import "fmt"

func main() {
	const N = 512

	a := 0;
	i := N;

	counter := 0

	// log(N)
	for i > 0 {
		a += i
		i /= 2
		counter++

		fmt.Println(counter, a, i)
	}
}
