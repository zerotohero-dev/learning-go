package main

import (
	"fmt"
	"math"
)

func findComplement2(num int) int {
	i := 0
	j := 0

	for i < num {
		i += int(math.Pow(2, float64(j)))
		j++
	}

	return i - num
}

func findComplement(num int) int {
	mask := ^1
	for num&mask > 0 {
		mask <<= 1
	}

	return (^mask) & (^num)
}

func main() {
	fmt.Println(findComplement(64))
	fmt.Println(findComplement2(64))
}
