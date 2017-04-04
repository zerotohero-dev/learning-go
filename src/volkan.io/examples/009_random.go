package main

import (
	"fmt"
	"volkan.io/math/rand"
)

func main() {
	// rand.GenerateInt(42) -> will panic because Seed() has not been called.

	rand.Seed()

	fmt.Println("Here is a random number ", rand.GenerateInt(34))
}
