package main

import (
	"log"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var (
	count int
	wg    sync.WaitGroup
	mutex sync.RWMutex
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func read() {
	for i := 0; i < 1000; i++ {
		mutex.RLock()
		{
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			// log.Printf("count %d\n", count)
		}
		mutex.RUnlock()
	}
	wg.Done()
}

func write() {
	for {
		mutex.Lock()
		{
			c := count
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			c++
			count = c

			log.Printf("count %d\n", count)
		}
		mutex.Unlock()
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	wg.Add(1)

	for i := 0; i < 10000; i++ {
		go read()
	}

	for i := 0; i < 10000; i++ {
		go write()
	}

	wg.Wait()
}
