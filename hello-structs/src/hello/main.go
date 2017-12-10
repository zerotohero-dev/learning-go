package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 42})
	fmt.Println(person{name: "Alice", age: 20})
	fmt.Println(&person{name: "Nina", age: 62})

	s := person{name: "Sean", age: 16}
	fmt.Println(s.name, s.age)

	// You can use dot reference with pointers to structs too.
	// They are automatically dereferenced.
	sp := &s
	fmt.Println(sp.name, sp.age)

	sp.age = 52

	fmt.Println(sp)
}
