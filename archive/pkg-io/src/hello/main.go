package main

import (
	"fmt"
	"poetry"
)

func main() {
	fmt.Println("Hello, World")

	p, err := poetry.LoadPoem("words")

	if err != nil {
		fmt.Printf("Error loading poem %s\n", err)
	}

	fmt.Printf("%s\n", p)
}
