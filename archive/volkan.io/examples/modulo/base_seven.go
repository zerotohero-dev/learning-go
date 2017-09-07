package main

import "fmt"

func convertToBase7(num int) string {
	if num < 0 {
		return fmt.Sprintf("-%v", convertToBase7(-num))
	}

	if num < 7 {
		return fmt.Sprintf("%v", num)
	}

	return fmt.Sprintf("%v%v", convertToBase7(num/7), num % 7)
}

func main() {
	fmt.Println(
		convertToBase7(8),
	)
}
