package main

import "fmt"

var romans = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sum := int(romans[s[len(s) - 1]])

	for i := len(s) - 2; i >= 0; i-- {
		if romans[s[i]] < romans[s[i+1]] {
			sum -= romans[s[i]]
		} else {
			sum += romans[s[i]]
		}
	}

	return sum
}

func main() {
	fmt.Println(
		romanToInt("MDCCCLXXXIV"),
	)
}
