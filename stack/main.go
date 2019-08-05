package main

import (
	"fmt"
	"github.com/v0lkan/go-stack/stack"
)

func main() {
	haystack := make(stack.Stack, 0)

	haystack.Push("hay")
	haystack.Push(42)
	haystack.Push([]string{"a", "b", "c"})

	fmt.Println(haystack)

	for {
		item, err := haystack.Pop()

		if err != nil {
			break
		}

		fmt.Println(item)
	}
}
