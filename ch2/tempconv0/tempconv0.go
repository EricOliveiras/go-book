package main

import "fmt"

type Celsius float64
type Fahrenheit float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BollingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {
	fmt.Printf("%g°\n", BollingC-FreezingC)
	bollingF := CToF(BollingC)
	fmt.Printf("%g°\n", bollingF-CToF(FreezingC))

	c := FToC(212.0)
	fmt.Println(c.String())
	fmt.Printf("%v\n", c)
	fmt.Printf("%s\n", c)
	fmt.Printf("%g\n", c)
	fmt.Println(float64(c))
	// fmt.Printf("%g\n", bollingF-FreezingC) erro de compilaçao de tipos
}
