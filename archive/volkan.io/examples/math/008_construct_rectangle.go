package main

import (
	"math"
	"fmt"
)

func constructRectangle(area int) []int {
	width := int(math.Sqrt(float64(area)))

	for area % width != 0 {
		width--
	}

	return []int{area/width, width};
}

func main() {
	fmt.Println(
		constructRectangle(62),
	)
}