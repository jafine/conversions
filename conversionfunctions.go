package conversions

// MetreToFoot converts metres to feet
func MetreToFoot(m Metre) Foot { return Foot(m * 3.281) }

// FootToMetre converts feet to Metres
func FootToMetre(f Foot) Metre { return Metre(f / 3.281) }

// CentimetreToInch converts centimetres to inches
func CentimetreToInch(cm Centimetre) Inch { return Inch(cm * 0.394) }

// InchToCentimetre converts inches to centimetres
func InchToCentimetre(in Inch) Centimetre { return Centimetre(in / 2.54) }
