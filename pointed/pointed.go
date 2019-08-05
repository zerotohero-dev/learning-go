package main

import "fmt"

func pointed(pItems *struct{v string}) {
	fmt.Printf("pItems(pointed): %p %v\n", pItems, pItems)
	fmt.Printf("*pItems(pointed): %v %v\n", *pItems, *pItems)
	fmt.Printf("&pItems(pointed): %p %v\n", &pItems, &pItems)
	fmt.Println("--------")

	assigned := pItems
	fmt.Printf("assigned: %p %v\n", assigned, assigned)
	fmt.Printf("pItems(pointed): %p %v\n", pItems, pItems)
	fmt.Printf("&assigned: %p %v\n", &assigned, &assigned)
	fmt.Printf("&pItems(pointed): %p %v\n", &pItems, &pItems)
	fmt.Println("--------")

	// Change the value of the original object.
	pItems.v = "stars"

	// The internal pItems is now pointing to a totally new object.
	// This will NOT have impact outside this function.
	pItems = &struct{v string}{"world"}
}

func main() {
	items := struct{v string}{"hello"}

	assigned := items

	assigned.v = "dongi"

	fmt.Printf("assigned: %v %v\n", assigned, assigned)
	fmt.Printf("items: %v %v\n", items, items)
	fmt.Printf("&assigned: %p %v\n", &assigned, &assigned)
	fmt.Printf("&items: %p %v\n", &items, &items)

	fmt.Println("--------")

	pItems := &items

	fmt.Printf("pItems: %p %v\n", pItems, pItems)
	fmt.Printf("&pItems: %p %v\n", &pItems, &pItems)
	fmt.Println("--------")

	pointed(pItems)

	fmt.Println(items)

	// Output:
	//
	// assigned: {dongi} {dongi}
	// items: {hello} {hello}
	// &assigned: 0xc00008e040 &{dongi}
	// &items: 0xc00008e030 &{hello}
	// --------
	// pItems: 0xc00008e030 &{hello}
	// &pItems: 0xc00008a010 0xc00008a010
	// --------
	// pItems(pointed): 0xc00008e030 &{hello}
	// *pItems(pointed): {hello} {hello}
	// &pItems(pointed): 0xc00008a018 0xc00008a018
	// --------
	// assigned: 0xc00008e030 &{hello}
	// pItems(pointed): 0xc00008e030 &{hello}
	// &assigned: 0xc00008a020 0xc00008a020
	// &pItems(pointed): 0xc00008a018 0xc00008a018
	// --------
	// {stars}
	//
	// 1. Assigning to a variable creates a variable with its own address.
	// 2. Assignment is essentially a copy operation (for, e.g., assigning var
	//    that refers to a struct to a new var copies the struct into the var)
	//
	// The two pointers hold the same memory address (0xc00008e030)
	// but they are two separate pointers at 0xc00008a018 and 0xc00008a010
	//
	// Similarly, a pointer receiver has a *copy* of a pointer.
	// When you make it point to something else, the original item will
	// NOT be affected.
	//
	// Essentially “everything is passed by value” in Go.
}