package main

import (
	"time"
	"math/rand"
	"fmt"
	"sync/atomic"
)

var (
	running int64
	test int64
)

func work(i int) {

	// No operation in go is atomic; so if you want to
	// update a shared resource you either use channels, or locks
	// or the atomic package.
	atomic.AddInt64(&running, 1)

	// This may or may not give what you expect.
	// Don’t rely on it!
	//
	// The only things which are guaranteed to be atomic in go are
	// the operations in sync.atomic.
	test++

	fmt.Printf("[%d %d", test, running)
	time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	fmt.Printf("]")

	test--

	atomic.AddInt64(&running, -1)
}

func worker(sema chan bool, i int) {
	<-sema
	work(i)
	sema <- true
}


func main() {

	// up to ten routines can be running at the same time.
	sema := make(chan bool, 10)

	for i := 0; i < 1000; i++ {
		go worker(sema, i)
	}

	for i := 0; i < cap(sema); i++ {
		fmt.Println("Sending true to channel…")
		sema <- true
	}

	time.Sleep(30 * time.Second)
}
