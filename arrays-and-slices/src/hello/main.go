package main

import (
	"fmt"
)

func main() {
	// arrays are passed by copy by default; not by ref.
	// [4]string{}
	words := [...]string{"the", "quick", "brown", "fox"}

	fmt.Printf("%s\n", words[2])

	// slices are passed by value, but their ref is a pointer to this array.
	// so for practical purposes you can think like slices are passed by ref.
	foxes := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}

	fmt.Println(len(foxes))

	// Also make([]string, 4, 10) to give capacity along with length.
	words2 := make([]string, 4)

	words2[3] = "Fox"

	words2 = append(words2, "Dongi")

	fmt.Println(words2, len(words2), cap(words2))

	newWords := make([]string, 4)
	copy(newWords, words2)

	fmt.Println(newWords)
}
