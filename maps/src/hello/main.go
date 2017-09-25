package main

import (
	"fmt"
)

func main() {
	dayMonths := make(map[string]int)

	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in February: %d\n", dayMonths["Feb"])
	fmt.Printf("Days in February: %d\n", dayMonths["Februray"])

	days, ok := dayMonths["January"]

	if !ok {
		fmt.Println("Can’t get days of January!")
	} else {
		fmt.Printf("There are %d days in January.\n", days)
	}

	for month, days := range dayMonths {
		fmt.Printf("%s has %d days.\n", month, days)
	}

	delete(dayMonths, "Feb")
	delete(dayMonths, "Feb")
	delete(dayMonths, "Feb")

	dayMonths2 := map[string]int {
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	fmt.Println(dayMonths2)
}
