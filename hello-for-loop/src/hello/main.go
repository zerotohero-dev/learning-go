package main

import "fmt"

func main() {
	// Infinite loop:
	//
	// for {
	// 	fmt.Println("Hello")
	// }

	counter := 0

	for counter < 10 {
		fmt.Printf("Hello World\n")

		counter++
	}

	for c, d := 0, 1; c < 10; c, d = c+1, d+1 {
		fmt.Printf("Hello %d %d\n", c, d)
	}

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n %2 == 0 {
			continue
		}

		fmt.Println(n)
	}
}
