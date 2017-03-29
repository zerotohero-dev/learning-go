package main

import (
	"fmt"
	"volkan.io/math/peano"
)

func main() {
	fmt.Println()

	for i := 0; i <= 9; i++ {
		f := peano.ToInt(
			peano.Fact(
				peano.ToPeano(i),
			),
		)

		fmt.Printf("%d! = %d\n", i, f)
	}
}
