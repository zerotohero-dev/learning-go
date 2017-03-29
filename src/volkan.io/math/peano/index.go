package peano

// Number is a pointer to Number.
type Number *Number

func zero() *Number {
	return nil
}

func isZero(x *Number) bool {
	return x == nil
}

func addOne(x *Number) *Number {
	e := new(Number)

	*e = x

	return e
}

func subOne(x *Number) *Number {
	return *x
}

func add(x, y *Number) *Number {
	if isZero(y) {
		return x
	}

	return add(addOne(x), subOne(y))
}

func mul(x, y *Number) *Number {
	if isZero(x) || isZero(y) {
		return zero()
	}

	return add(mul(x, subOne(y)), x)
}

func Fact(n *Number) *Number {
	if isZero(n) {
		return addOne(zero())
	}

	return mul(Fact(subOne(n)), n)
}

func ToPeano(n int) *Number {
	if n > 0 {
		return addOne(ToPeano(n - 1))
	}

	return zero()
}

func ToInt(x *Number) int {
	if isZero(x) {
		return 0
	}

	return ToInt(subOne(x)) + 1
}
