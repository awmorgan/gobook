// main performs Celsius and Fahrenheit conversions.
package main

import "fmt"

func main() {
	c := celsius(23)
	fmt.Println(c)
}

type celsius float64
type fahrenheit float64

const (
	absoluteZeroC celsius = -273.15
	freezingC     celsius = 0
	boilingC              = 100
)

func (c celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}
