package main

import "fmt"

func swap(n *int, m *int) {
	mv := *m
	nv := *n

	*n = mv
	*m = nv
}

// log(N)
// n      m
// F5 (=) F4 | + F3
// F3     F4
// F4 (=) F3 | + F2
// F2     F3
// F3 (=) F2 | + F1
// F2     F1
// F2 (=) F1 | + F0
func gcd(n int, m int) int {
	if (n%m == 0) {
		return m;
	}

	if (n < m) {swap(&n, &m)}

	for m > 0 {
		n = n%m
		swap(&n, &m)
	}

	return n
}

func main() {
	a := 2
	b := 4

	swap(&a, &b)

	fmt.Println(a, b)

	a, b = b, a

	fmt.Println(a, b)

	fmt.Println(gcd(152, 32))
}
