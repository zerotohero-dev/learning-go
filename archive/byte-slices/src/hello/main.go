package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.OpenFile("test.txt", os.O_RDWR, os.ModeAppend)

	someString := "foobar"

	_, error := f.Write([]byte(someString))

	if error != nil {
		fmt.Println(error)
	}

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	defer f.Close()

	b := make([]byte, 100)

	n, err := f.Read(b)

	fmt.Printf("%d %t \n", n, b)

	stringed := string(b)

	fmt.Println(stringed)
}
