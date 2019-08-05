package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("!") // This deferred won’t execute.

	os.Exit(3)

	// Execute `echo $?` after the program exits.
}
