package main

import (
	"fmt"
)

func hamming(x int, y int) int {
	tmp := x ^ y
	dis := 0

	for tmp != 0 {
		tmp1 := ((tmp >> 1) << 1)

		if tmp1 != tmp {
			dis++
		}

		tmp >>= 1
	}

	return dis
}

func main() {
	fmt.Println(hamming(12, 224))
}
