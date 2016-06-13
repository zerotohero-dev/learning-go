package main

import (
	"fmt"
)

func sumOfFibonacciNumbers(number int64) int64 {
	if number <= 0 {
		return 0
	}

	var results = make([]int64, number + 1, number + 1)

	sum := int64(1)

	results[0] = 0
	results[1] = 1

	for i := 2; int64(i) <= number; i++ {
		results[i] = results[i-1] + results[i-2]
		sum += results[i]
	}

	return sum
}

func fibonacciRecursive(number int64) int64 {
	if (number <= 0) {
		return 0
	}

	if (number == 1) {
		return 1
	}

	return fibonacciRecursive(number-1) + fibonacciRecursive(number-2)
}

func fibonacciTailRecursiveHelper(number int64, old int64, current int64) int64 {
	if number == 1 {
		return current
	}

	return fibonacciTailRecursiveHelper(number - 1, current, old + current)
}

func fibonacciTailRecursive(number int64) int64 {
	if (number <= 0) {
		return 0
	}

	if (number == 1) {
		return 1
	}


	return fibonacciTailRecursiveHelper(number, 0, 1)
}

func main() {
	fmt.Printf("%d\n", sumOfFibonacciNumbers(10))
	fmt.Printf("%d\n", fibonacciRecursive(5))
	fmt.Printf("%d\n", fibonacciTailRecursive(5))
}
