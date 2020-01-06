package units

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func FToM(f Feet) Meter {
	return Meter(f/3.2808)
}

func MToF(m Meter) Feet {
	return Feet(m*3.2808)
}

func LbToKg(lb Pound) Kilogram {
	return Kilogram(lb/2.20462)
}

func KgToLb(kg Kilogram) Pound {
	return Pound(kg*2.20462)
}