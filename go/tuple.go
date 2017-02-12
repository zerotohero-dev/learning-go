package main

func swap(x *int, y *int) {
	*x, *y = *y, *x
}

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func fib(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

func main() {
	fib(12)

	a := 22
	b := 44

	swap(&a, &b)
}
