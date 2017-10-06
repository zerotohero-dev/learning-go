package main

import (
	"log"
	"runtime"
	"sync"
)

var (
	count int
	wg    sync.WaitGroup
	mutex sync.Mutex
)

func increment() {
	for i := 0; i < 10000; i++ {
		mutex.Lock()
		{
			c := count
			runtime.Gosched() // To generate a race condition.
			c++
			count = c
		}
		mutex.Unlock()
	}

	wg.Done()
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(2)

	go increment()
	go increment()

	wg.Wait()

	log.Printf("counter %d\n", count)
}
