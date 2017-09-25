package main

import (
	"fmt"
	"math/rand"
)

type shuffler interface {
	Len() int
	Swap(i, j int)
}

func shuffle(s shuffler) {
	for i := 0; i < s.Len(); i++ {
		j := rand.Intn(s.Len()-i)
		s.Swap(i, j)
	}	
}

type intSlice []int
// type stringSlice []string

func (i intSlice) Len() int {
	return len(i)
}

func (is intSlice) Swap(i, j int) {
	is[i], is[j] = is[j], is[i]
}


func main() {
	fmt.Println("Hello Interfaces")

	is := intSlice{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	shuffle(is)

	fmt.Printf("%d\n", is)
}
