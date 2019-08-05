package main

import (
	"fmt"
	"shuffler"
)

type intSlice []int

func (is intSlice) Len() int {
	return len(is)
}
func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}

func main() {
	fmt.Println("Hello Packages")

	is := intSlice{1, 2, 4, 3, 5, 6, 1, 3, 11, 44}

	shuffler.Shuffle(is)

	fmt.Printf("%v\n", is)
}
