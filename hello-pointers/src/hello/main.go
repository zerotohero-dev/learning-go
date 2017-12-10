package main

import "fmt"

func zeroval(val int) {
	val = 0
}

func zeroptr(ptr *int) {
	*ptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)
}
