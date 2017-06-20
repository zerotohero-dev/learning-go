package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	var counts [26]int

	for _, letter := range magazine {
		counts[letter-'a']++
	}

	for _, letter := range ransomNote {
		counts[letter-'a']--

		if counts[letter-'a'] < 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(
		canConstruct("dolorem", "idpsolorsimemsumorem"),
	)
}