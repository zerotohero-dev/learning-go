package main

import (
	"fmt"
	"sort"
)

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}

	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}

	return true
}

func main() {
	ages := make(map[string]int)

	otherAges := map[string]int{
		"alice": 31,
		"joe":   44,
	}

	fmt.Println(ages)
	fmt.Println(otherAges)

	delete(otherAges, "alice")

	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}

	//var names []string
	names := make([]string, 0, len(ages))

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)

	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}

	age, ok := ages["bob"]

	fmt.Println(age, ok)

	if !ok {
		fmt.Println("key bob does not exist.")
	}

	equal(ages, otherAges)
}
