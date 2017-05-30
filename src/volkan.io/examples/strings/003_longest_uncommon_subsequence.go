package main

import "fmt"

func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}

	lenA := len(a)
	lenB := len(b)

	if lenA > lenB {
		return lenA
	}

	return lenB
}

func main() {
	fmt.Println(
		findLUSlength("loremdolorem", "ipsumdolipsum"),
	)
}
