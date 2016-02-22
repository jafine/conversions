package conversions

// MetreToFoot converts metres to feet
func MetreToFoot(m Metre) Foot { return Foot(m * 3.281) }

// FootToMetre converts feet to Metres
func FootToMetre(f Foot) Metre { return Metre(f / 3.281) }

// CentimetreToInch converts centimetres to inches
func CentimetreToInch(cm Centimetre) Inch { return Inch(cm * 0.394) }

// InchToCentimetre converts inches to centimetres
func InchToCentimetre(in Inch) Centimetre { return Centimetre(in / 2.54) }

// PsiToBar converts pounds per square inch to bar
func PsiToBar(psi Psi) Bar { return Bar(psi * 0.0689476) }

// BarToPsi converts bar to pounds per square inch
func BarToPsi(bar Bar) Psi { return Psi(bar * 14.5038) }

// CelsiusToFahrenheit converts a Celsius temperature to Fahrenheit
func CelsiusToFahrenheit(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FahrenheitToCelsius converts a Fahrenheit temperature to Celsius
func FahrenheitToCelsius(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// KelvinToCelsius converts a Kelvin temperature to Celsius
func KelvinToCelsius(k Kelvin) Celsius { return Celsius(k - 273.15) }

// CelsiusToKelvin converts a Celsius temperature to Kelvin
func CelsiusToKelvin(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// KelvinToFahrenheit converts a Kelvin temperature to Fahrenheit
func KelvinToFahrenheit(k Kelvin) Fahrenheit { return Fahrenheit(k*9/5 - 459.67) }

// FahrenheitToKelvin converts a Fahrenheit temperature to Kelvin
func FahrenheitToKelvin(f Fahrenheit) Kelvin { return Kelvin((f + 459.67) * 5 / 9) }
