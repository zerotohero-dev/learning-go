package main

import "fmt"

var keyExists = struct{}{}

func add(set map[interface{}]struct{}, item interface{}) {
	set[item] = keyExists
}

func remove(set map[interface{}]struct{}, item interface{}) {
	_, ok := set[item]

	if !ok {return}

	delete(set, item)
}

func contains(set map[interface{}]struct{}, item interface{}) bool {
	_, ok := set[item]

	return ok
}

func canPermutePalindrome(s string) bool {
	set := make(map[interface{}]struct{})

	for _, c := range s {
		if contains(set, c) {
			remove(set, c)
		} else {
			add(set, c)
		}
	}

	length := len(set)

	return length == 0 || length == 1
}

func main() {
	fmt.Println(canPermutePalindrome("doldolboox"))
}