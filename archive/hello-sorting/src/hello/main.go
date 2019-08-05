package main

import (
	"fmt"
	"sort"
)

// ByLength is an alias type to sort strings by length
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	ints := []int{7, 3, 17}
	sort.Ints(ints)
	fmt.Println("Ints", ints)

	fmt.Println("sorted?", sort.IntsAreSorted(ints))

	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))

	fmt.Println(fruits)
}
