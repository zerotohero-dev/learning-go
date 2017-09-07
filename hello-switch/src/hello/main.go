package main

import (
	"fmt"
	"os"
)

func main() {
	n, err := fmt.Printf("Hello, World!")

	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("\nNo bytes output\n")
		fallthrough
	case n != 13:
		fmt.Printf("\nWrong number of characters %d\n", n)
	default:
		fmt.Printf("\nOK\n")
	}

	atoz := "the quick brown fox jumps over the lazy dog"

	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels++
		case 'z':
			zeds++
			fallthrough
		default:
			consonants++
		}
	}

	fmt.Printf("\nVowels: %d; Consonants: %d; Zeds: %d\n", vowels, consonants, zeds)
}
