package conversions

// MetreToFoot converts metres to feet
func MetreToFoot(m Metre) Foot { return Foot(m * 3.281) }

// FootToMetre converts feet to Metres
func FootToMetre(f Foot) Metre { return Metre(f / 3.281) }
