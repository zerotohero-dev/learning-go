package main

import (
	"fmt"
	"reflect"
	"volkan.io/ds/tree"
)

func main() {
	tree := tree.BTree{
		&tree.BTree{nil, 42, nil},

		43,

		&tree.BTree{
			&tree.BTree{
				&tree.BTree{nil, 44, nil},
				45,
				nil,
			},
			46,
			&tree.BTree{nil, 99, nil},
		},
	}

	fmt.Printf("Typeof tree: “%T”.\n", tree)
	fmt.Printf("Typeof tree: “%s”.\n", reflect.TypeOf(tree))
	fmt.Println("")
	fmt.Println(&tree)

	/*
		Output:

		Typeof tree: “tree.BTree”.
		Typeof tree: “tree.BTree”.

		\-43
		  |-42
		  | |-<nada>
		  | \-<nada>
		  \-46
		    |-45
		    | |-44
		    | | |-<nada>
		    | | \-<nada>
		    | \-<nada>
		    \-99
		      |-<nada>
		      \-<nada>
	*/
}
