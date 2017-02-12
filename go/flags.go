package main

// go run flags.go -h

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit traling newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()

	fmt.Print(strings.Join(flag.Args(), *sep))

	if !*n {
		fmt.Println()
	}
}
