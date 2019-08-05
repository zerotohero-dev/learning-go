package main

import (
	"fmt"
	"math"
)

const s = "constant"

func main() {
	fmt.Println(s)

	const n = 5000000000

	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}
