package main

import (
	"fmt"
	. "volkan.io/person"
)

type Stringer interface {
	String() string
}

func main() {
	p := Person{"Volkan", 412}

	fmt.Println(p)

	var s Stringer

	s = &p

	// s = &p
	// func(p Person) String() string
	// String() string
	// … works
	//
	// s = p
	// func(p Person) String() string
	// String() string
	// … works
	//
	// s = p
	// func(p *Person) String() string
	// String() string
	// … gives:
	//      cannot use p (type person.Person) as type Stringer in assignment:
	//      person.Person does not implement Stringer
	//          (String method has pointer receiver)
	//
	// So if the concrete method has a pointer receiver, the reference needs to be a pointer type too.
	// But if the concrete method has a non-pointer receiver, the reference can either be a
	// pointer type or a regular type.
	//
	// See <https://golang.org/ref/spec#Method_sets> for why that is how it is.
	//
	//      “The method set of the corresponding pointer type *T is the set of all methods
	//      declared with receiver *T or T (that is, it also contains the method set of T)”

	fmt.Println(s)
}
