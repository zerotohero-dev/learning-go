package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}

	return true
}

func appending() {
	var runes []rune

	for _, r := range "Hello, world" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)
}

func appendInt(x []int, y int) []int {
	var z []int

	zlen := len(x) + 1

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen

		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)

		copy(z, x)
	}

	z[len(x)] = y

	return z
}

func appendIntVariadic(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)

	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen

		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		z = make([]int, zlen, zcap)

		copy(z, x)
	}

	copy(z[len(x):], y)

	return z
}

func nonempty(strings []string) []string {
	i := 0

	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}

	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]

	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}

	return out
}

func remove(slice []int, i int) []int {
	copy(slice[i:], slice[i+1:])

	return slice[:len(slice)-1]
}

func removeNoOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]

	return slice[:len(slice)-1]
}

func main() {
	a := [...]int{0, 1, 2, 3, 4, 5}

	reverse(a[:])

	fmt.Println(a)

	b := []string{"a", "b", "c"}

	s1 := b[0:1]
	s2 := b[1:2]

	fmt.Println(equal(s1, s2))

	appending()

	appendInt(a[:], 42)
	appendIntVariadic(a[:], 1, 2, 3)

	var x, y []int

	for i := 0; i < 10; i++ {
		y = appendInt(x, i)

		x = y

		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)

		x = y
	}

	var f []int

	f = append(f, 1)
	f = append(f, 2, 3)
	f = append(f, 4, 5, 6)
	f = append(f, f...)

	fmt.Println(f)

	nonempty(b[:])
	nonempty2(b[:])

	remove(a[:], 1)
	removeNoOrder(a[:], 1)
}
