package tempconv

import "math"

// All conversion functions with a celsius value as input
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func CToK(c Celsius) Kelvin     { return Kelvin(c + 273.15) }

// All conversion functions with a fahrenheit value as input
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
func FToK(f Fahrenheit) Kelvin  { return CToK(FToC(f)) }

// All conversion functions with a kelvin value as input
func KToC(k Kelvin) Celsius    { return Celsius(math.Round(float64((k-273.15)*100)) / 100) }
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }
