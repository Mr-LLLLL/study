package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZereC Celsius = -273.15
	FreezingC Celsius = 0
	BoilingC Celsius = 100
)

func (c Celsius) String() string {
	return fmt.Sprintf("%gºC", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%gºC", f)
}
