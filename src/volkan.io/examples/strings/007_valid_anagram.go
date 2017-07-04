package main

import "fmt"

func isAnagram(s string, t string) bool {
	var alphabet [26]int

	for _, c := range s {
		alphabet[c - 'a']++
	}

	for _, c := range t {
		alphabet[c - 'a']--
	}

	for _, c := range alphabet {
		if c != 0 {
			return false
		}
	}

	return true
}


func main() {
	fmt.Println(
		isAnagram("congo", "ongco"),
	)
}

