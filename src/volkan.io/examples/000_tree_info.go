package main

import (
	"fmt"
	"reflect"
	"volkan.io/ds"
)

func main() {
	tree := ds.Tree{nil, 42, nil}

	fmt.Println("Tree:")
	fmt.Printf("Typeof tree: “%T”.\n", tree)
	fmt.Printf("Typeof tree: “%s”.\n", reflect.TypeOf(tree))
	fmt.Println("")
}
