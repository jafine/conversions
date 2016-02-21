// Package conversions performs a variety of unit translations in
// a type-safe manner.
// Copyright (c) 2015 David J Peacock - david@davidjpeacock.ca
package conversions

import "fmt"

/*
 * Lengths
 */

// Metre - metric
type Metre float64

// Foot - imperial
type Foot float64

// Centimetre - metric
type Centimetre float64

// Inch - imperial
type Inch float64

/*
 * Pressures
 */

// Psi - imperial
type Psi float64

// Bar - metric
type Bar float64

/*
 * Temperatures
 */

// Celsius - metric
type Celsius float64

// Fahrenheit - imperial
type Fahrenheit float64

// Kelvin - scientific
type Kelvin float64

// Constants
const (
	CentimetresPerMetre Metre   = 100
	InchesPerFoot       Foot    = 12
	AbsoluteZeroC       Celsius = -273.15
	FreezingC           Celsius = 0
	BoilingC            Celsius = 100
	KelvinZeroC         Celsius = -273.15
)

/*
 * Lengths
 */

func (m Metre) String() string       { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string        { return fmt.Sprintf("%gft", f) }
func (cm Centimetre) String() string { return fmt.Sprintf("%gcm", cm) }
func (in Inch) String() string       { return fmt.Sprintf("%gin", in) }

/*
 * Pressures
 */

func (psi Psi) String() string { return fmt.Sprintf("%gpsi", psi) }
func (bar Bar) String() string { return fmt.Sprintf("%gbar", bar) }

/*
 * Temperatures
 */

func (c Celsius) String() string    { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%gºK", k) }
