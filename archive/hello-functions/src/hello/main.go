package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Println(nums)

	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func intSeq() func() int {
	i := 0

	return func() int {
		i++

		return i
	}
}

func fact(n int) int {
	if n <= 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	_, c := vals()
	fmt.Println(c)

	sum(1, 2, 3, 4, 5)

	numbers := []int{1, 2, 3, 4, 5}
	sum(numbers...)

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()

	fmt.Println(newInts())

	fmt.Println(fact(7))
}
