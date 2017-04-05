package main

import (
	"fmt"
)

func main() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "World"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Println(primes)

	s := primes[1:4]

	fmt.Println(s)

	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Println(names)

	as := names[0:2]
	bs := names[1:3]

	fmt.Println(as, bs)

	bs[0] = "XXX"

	fmt.Println(as, bs)
	fmt.Println(names)

	q := []int{2, 3, 5, 7, 11, 13}

	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}

	fmt.Println(r)

	sts := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}

	fmt.Println(sts)

	var nilSlice []int

	if nilSlice == nil {
		fmt.Println("There is a nil slice!")
		fmt.Println(len(nilSlice), cap(nilSlice))
	}

	dynamic := make([]int, 5)
	dynamic2 := make([]int, 0, 5)

	fmt.Println(len(dynamic), cap(dynamic), len(dynamic2), cap(dynamic2))
}
