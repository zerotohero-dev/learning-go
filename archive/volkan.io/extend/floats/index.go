package floats

type FloatX float64

func (f FloatX) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}
