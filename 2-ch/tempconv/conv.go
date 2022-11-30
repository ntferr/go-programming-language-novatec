package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }
func FToC(f Fahrenheit) Celsius  { return Celsius((f - 32) * 5 / 9) }
func CToK(k kelvin) kelvin       { return kelvin(k - 273.15) }
