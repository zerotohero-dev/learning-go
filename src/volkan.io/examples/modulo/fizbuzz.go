package main

import "fmt"

func fizzBuzz(n int) []string {
	result := make([]string, 0)

	for i, fizz, buzz := 1, 0, 0; i <= n; i++ {
		fizz++
		buzz++

		if fizz == 3 && buzz == 5 {
			result = append(result, "FizzBuzz")
			fizz = 0
			buzz = 0
		} else if fizz == 3 {
			result = append(result, "Fizz")
			fizz = 0
		} else if buzz == 5 {
			result = append(result, "Buzz")
			buzz = 0
		} else {
			result = append(result, fmt.Sprint(i))
		}
	}

	return result
}

func main() {
	fmt.Println(
		fizzBuzz(42),
	)
}