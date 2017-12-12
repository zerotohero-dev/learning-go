package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

//  ./bin/hello --word=optimus lorem ipsum dolor sit ahmet

func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	wordPtr := flag.String("word", "foo", "a string")

	flag.Parse()

	fmt.Println("word:", *wordPtr)

	os.Setenv("FOO", "1")

	fmt.Println(os.Getenv("FOO"))

	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0], pair[1])
	}
}
