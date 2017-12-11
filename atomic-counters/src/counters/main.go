package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	// Since there is a chance that the counter might still be
	// updated, we atomically copy its current value.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println(opsFinal)
}
