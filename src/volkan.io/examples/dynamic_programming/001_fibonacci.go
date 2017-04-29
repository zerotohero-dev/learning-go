package main

import (
	"fmt"
)

func main() {
	index := 15
	memory := make([]int, 2)
	memory[0] = 1
	memory[1] = 1

	for i := 2; i < index; i++ {
		memory = append(memory, 0)
	}

	for i := 2; i < index; i++ {
		memory[i] = memory[i-2] + memory[i-1]
	}

	fmt.Println(memory)
}
