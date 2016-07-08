package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	// sends to a buffered channel block only when the buffer is full.
	// receives from a buffered channel block only when the buffer is empty.
	ch <- 1
	fmt.Println(<-ch)
	ch <- 2
	ch <- 2

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)
}
