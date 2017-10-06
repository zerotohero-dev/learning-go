package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make(chan string)

	var w sync.WaitGroup

	w.add(2)

	go func(queue chan string, wg *sync.WaitGroup) {
		defer wg.Done()

		fmt.Println("func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("func 1")

		queue <- "done func 1"
	}(queue, &w)

	go func(queue chan string, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("func 1")
		time.Sleep(2 * time.Second)
		fmt.Println("func 1")

		queue <- "done func 2"
	}(queue, &w)

	// fmt.Println(<-queue)
	// fmt.Println(<-queue)

	//for text := range queue {
	//	fmt.Println(text)
	//}

	w.Wait()
}
