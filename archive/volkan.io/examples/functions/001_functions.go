package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	// naked return is required.
	return
}

func main() {
	a, b := swap("hello", "world")

	fmt.Println(a, b)
	fmt.Println(add(44, 22))
	fmt.Println(split(42))
}
