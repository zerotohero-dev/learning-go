package main

import (
	"fmt"
)

func do(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Integer:", i)
	case string:
		fmt.Println("String:", i)
	default:
		fmt.Println("Don’t know what this is:", i)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
