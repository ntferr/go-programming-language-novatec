package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit  { return Fahrenheit(c*9/5 + 32) }
func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }
func FToC(f Fahrenheit) Celsius  { return Celsius((f - 32) * 5 / 9) }

func main() {
	var c Celsius
	result := CToF(c)
	fmt.Println(result)
	fmt.Printf("%T", result)
	//---
	var f Fahrenheit
	result2 := FToC(f)
	fmt.Println(result2)
	fmt.Printf("%T", result2)
	//---
	fmt.Printf("%g\n", BoilingC-FreezingC)
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC))
	// fmt.Printf("%g\n", boilingF-FreezingC) incompatibilidade de tipos
	//---
	fmt.Println(c == 0)
	fmt.Println(f >= 0)
	//fmt.Println(c == f) incompatibilidade de tipos
	fmt.Println(c == Celsius(f))
	//---
}
