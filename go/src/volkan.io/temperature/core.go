package temperature

import "fmt"

type Celcius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celcius = -273.5
	FreezingC     Celcius = 0
	BoilingC      Celcius = 100
)

func (c Celcius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
