package main

import (
	"fmt"
	"time"
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

	i := 2
	fmt.Print("Write", i , " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend.")
	default:
		fmt.Println("It is a weekday.")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It’s before noon.")
	default:
		fmt.Println("It’s afternoon.")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I’m a bool.")
		case int:
			fmt.Println("I’m an int")
		default:
			fmt.Printf("Don’t know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
