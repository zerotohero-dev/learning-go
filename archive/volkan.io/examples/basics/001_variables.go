package main

import (
	"fmt"
)

var c, python, java bool
var k, l int = 12, 44

// Constants can be character, string, boolean, or numeric values:
const Pi = 3.14

func main() {
	var i int

	// Shorthand notation:
	t := 3
	p, s := true, 44

	fmt.Println(i, c, python, java)
	fmt.Println(k, l)
	fmt.Println(t, p, s, Pi)
}

/*
Go’s basic types:
    bool
    string
    int  int8  int16  int32  int64
    uint uint8 uint16 uint32 uint64 uintptr

    byte // alias for uint8

    rune // alias for int32
         // represents a Unicode code point

    float32 float64

    complex64 complex128
*/
