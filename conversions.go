// Package conversions performs a variety of unit translations
package conversions

import "fmt"

// Metre - metric
type Metre float64

// Foot - imperial
type Foot float64

// Centimetre - metric
type Centimetre float64

// Inch - imperial
type Inch float64

// Constants
const (
	CentimetresPerMetre Metre = 100
	InchesPerFoot       Foot  = 12
)

func (m Metre) String() string       { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string        { return fmt.Sprintf("%gft", f) }
func (cm Centimetre) String() string { return fmt.Sprintf("%gcm", cm) }
func (in Inch) String() string       { return fmt.Sprintf("%gin", in) }
