package main

import (
	"flag"
	"fmt"
	"github.com/mrunalpagnis/WayToGo/exercises/units"
)

var weight = flag.Float64("w", 100, "input weight / mass")
var length = flag.Float64("l", 5, "input length / height")
var temperature = flag.Float64("t", 32, "input temperature")

func main() {
	flag.Parse()

	//temperature calculation
	c := units.Celsius(*temperature)
	f := units.CToF(c)
	fmt.Printf("%s = %s\n", c.String(), f.String())
	f = units.Fahrenheit(*temperature)
	c = units.FToC(f)
	fmt.Printf("%s = %s\n", f.String(), c.String())

	//weight calculation
	lb := units.Pound(*weight)
	kg := units.LbToKg(lb)
	fmt.Printf("%s = %s\n", lb.String(), kg.String())
	kg = units.Kilogram(*weight)
	lb = units.KgToLb(kg)
	fmt.Printf("%s = %s\n", kg.String(), lb.String())

	//length calculation
	ft := units.Feet(*length)
	m := units.FToM(ft)
	fmt.Printf("%s = %s\n", ft.String(), m.String())
	m = units.Meter(*length)
	ft = units.MToF(m)
	fmt.Printf("%s = %s\n", m.String(), ft.String())
}
