package main

import (
	"fmt"
)

func whatIsThis(i interface{}) {
	fmt.Printf("%T\n", i)

	switch v := i.(type) {
	case int:
		fmt.Println(i.(int))
		fmt.Printf("It is an int: %d\n", v)
	default:
		fmt.Printf("Unknown type: %v\n", v)	
	}
}

func main() {
	whatIsThis(42)
}
