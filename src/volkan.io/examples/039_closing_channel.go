package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		c <- x

		x, y = y, x+y
	}

	// **Notes**
	//
	// * Only the sender should close a channel, never the receiver.
	// * Sending on a closed channel will cause a panic.

	close(c)
}

func main() {
	c := make(chan int, 10)

	go fibonacci(cap(c), c)

	for i := range c {
		fmt.Println(i)
	}
}
