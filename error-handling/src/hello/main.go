package main

import (
	"fmt"
	"os"
	"errors"
)

var (
	errorEmptyString = errors.New("Unwilling to print an empty string!")
)


func printer(msg string) error {
	if msg == "" {
		// return fmt.Errorf("I am unwilling to print an empty string!")
		// panic(errorEmptyString)
		return errorEmptyString
	}

	_, err := fmt.Printf("%s\n", msg)

	return err
}

func main() {
	// Idiomatic go says that errors are returned from the functions
	// and the calling function has to deal with the errors.
	// There is also an exception-like construct using `panic` and `recover`;
	// however, that is extremely rarely used.
	// Go programs are not build around exceptions, they are built around
	// handling errors returned by functions.

	if err := printer("Hello, world."); err != nil {
		os.Exit(1)
	}

	err2 := printer("")

	if err2 == errorEmptyString {
		fmt.Println("You tried to print an empty string!")
	}

	fmt.Println(err2)
}
