package main

import (
	"fmt"
	"time"
	"volkan.io/counter/safe"
)

func main() {
	c := safe.NewCounter()

	for i := 0; i < 100001; i++ {
		go c.Inc("volkan.io")
	}

	fmt.Println()
	fmt.Println(c.Value("volkan.io"))

	time.Sleep(time.Second)

	fmt.Println()
	fmt.Println(c.Value("volkan.io"))
	fmt.Println()
}
