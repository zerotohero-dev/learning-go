package main

import (
	"fmt"
	"os"
)

func main() {
	if n, err := fmt.Printf("Hello, World!\n"); err != nil {
		fmt.Println(n)
		os.Exit(1)
	} else {
		fmt.Printf("Printed %d bytes\n", n)
	}

	// fmt.Printf("Printed %d bytes\n", n)
}