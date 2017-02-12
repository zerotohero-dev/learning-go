package temperature

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}
