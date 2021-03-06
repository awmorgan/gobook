package tempconv

import (
	"flag"
	"fmt"
)

// Celsius is a temperature in degrees Celsius.
type Celsius float64

// Fahrenheit is a temperature in degrees Fahrenheit.
type Fahrenheit float64

// CToF converts from degrees Celsius to degrees Fahrenheit.
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9.0/5.0 + 32.0)
}

// FToC converts from degrees Fahrenheit to degrees Celsius.
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32.0) * 5.0 / 9.0)
}

// String returns a string representing the Celsius degrees of temperature.
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

// *celsiusFlag satisfies the flag.Value interface.
type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit) // no error check needed
	switch unit {
	case "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

// CelsiusFlag defines a Celsius flag with the specified name,
// default value, and usage, and returns the address of the flag variable.
// The flag argument must have a quantity and a unit, e.g., "100C"
func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}
