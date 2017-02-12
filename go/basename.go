package main

import (
	"fmt"
	"strings"
)

func baseName(s string) string {
	var temp string

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			temp = s[i+1:]

			break
		}
	}

	for i := len(temp) - 1; i >= 0; i-- {
		if temp[i] == '.' {
			temp = temp[:i]

			break
		}
	}

	return temp
}

func baseName2(s string) string {
	slash := strings.LastIndex(s, "/")
	temp := s[slash+1:]

	if dot := strings.LastIndex(temp, "."); dot >= 0 {
		temp = temp[:dot]
	}

	return temp
}

func main() {
	path := "/a/b/c/d/f.go"

	fmt.Println(baseName(path))
	fmt.Println(baseName2(path))
}
