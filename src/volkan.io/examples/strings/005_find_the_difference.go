package main

import "fmt"

func findTheDifference(s string, t string) rune {
	rs := []rune(s)
	rt := []rune(t)
	theRune := rt[len(s)]

	for i, _ := range rs {
		theRune -= rs[i]
		theRune += rt[i]
	}

	return theRune
}

func main() {
	fmt.Println(
		string(findTheDifference("abcde", "abcfde")),
	)
}