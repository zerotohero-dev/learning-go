package main

import (
	"fmt"
)

func GreetNames(names []string, prefix string) {
	for _, n := range names {
		fmt.Printf("<%s>: %s\n", prefix, n)
	}
}

func main() {
	names := []string{
		"Anakin",
		"Obi-Wan",
		"Leia",
		"Han Solo",
		"Windu",
		"Vader",
		"Jar Jar",
	}

	pipe := make(chan string)

	go func() {
		GreetNames(names, "C")
		pipe <- "Done!"
	}()

	GreetNames(names, "M")

	fmt.Println(<-pipe)
}
