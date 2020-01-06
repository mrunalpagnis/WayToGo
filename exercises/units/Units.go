package units

import "fmt"

type Celsius float64
type Fahrenheit float64
type Feet float64
type Meter float64
type Pound float64
type Kilogram float64

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func (f Fahrenheit) String() string {
	return fmt.Sprintf("%g°F", f)
}

func (f Feet) String() string {
	return fmt.Sprintf("%g ft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%g m", m)
}

func (lb Pound) String() string {
	return fmt.Sprintf("%g lb", lb)
}

func (kg Kilogram) String() string {
	return fmt.Sprintf("%g kg", kg)
}
