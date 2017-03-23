package main

import (
	"fmt"
	"time"
	"volkan.io/life"
)

func main() {
	l := life.New(149, 19)

	for i := 0; i < 300; i++ {
		l.Step()

		// CLRSCR
		fmt.Print("\x0c", l)

		time.Sleep(time.Second / 30)
	}
}
