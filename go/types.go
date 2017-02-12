package main

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -237.15
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func CToF(c Celcius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celcius {
	return Celcius((f - 32) * 5 / 9)
}

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Println(CToF(AbsoluteZeroC))

	c := Celcius(100)

	fmt.Println(c.String())
}
