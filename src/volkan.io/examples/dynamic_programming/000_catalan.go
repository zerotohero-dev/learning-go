package main

import (
	"fmt"
)

func main() {
	n := 10
	a := [11]int{}

	a[0] = 1
	a[1] = 1

	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			a[i] += a[j] * a[i-1-j]
		}
	}

	fmt.Println(a)
}
