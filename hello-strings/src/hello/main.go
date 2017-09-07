package main

import (
	"fmt"
)

func main() {
	atoz := "the quick brown fox jumps over the lazy dog\n"

	fmt.Printf("%s\n", atoz[0:9])
	fmt.Printf("%s\n", atoz[:9])
	fmt.Printf("%s\n", atoz[15:])
	fmt.Printf("%s\n", atoz[15:19])
	fmt.Printf("%s\n", atoz[:])

	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	fmt.Printf("len %d\n", len(atoz))

	backquotes := `
		this is taken
		completely literally"";''\n
	`

	fmt.Println(backquotes)
}
