package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs

			if more {
				fmt.Println("Received job:", j)
			} else {
				fmt.Println("Received all the jobs. ", j)
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sent job:", j)
	}
	close(jobs)

	// If you don’t wait for this, the program will exit without
	// consuming the jobs.
	<-done
}
