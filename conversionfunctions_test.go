package conversions_test

import (
	"testing"

	"github.com/davidjpeacock/conversions"
)

func TestMetreToFoot(t *testing.T) {
	t.Log("Converting metres to feet... (expected feet: 9.843)")
	answer := conversions.MetreToFoot(3)

	if answer != 9.843 {
		t.Errorf("Expected 9.843, but got %g", answer)
	}
}

func TestFootToMetre(t *testing.T) {
	t.Log("Converting feet to metres... (expected metres: 2.7430661383724475)")
	answer := conversions.FootToMetre(9)

	if answer != 2.7430661383724475 {
		t.Errorf("Expected 2.7430661383724475, but got %g", answer)
	}
}

func TestCentimetreToInch(t *testing.T) {
	t.Log("Converting centimetres to inches... (expected inches: 3.9400000000000004)")
	answer := conversions.CentimetreToInch(10)

	if answer != 3.9400000000000004 {
		t.Errorf("Expected 3.9400000000000004, but got %g", answer)
	}
}

func TestInchToCentimetre(t *testing.T) {
	t.Log("Converting inches to centimetres... (expected centimetres: 3.937007874015748)")
	answer := conversions.InchToCentimetre(10)

	if answer != 3.937007874015748 {
		t.Errorf("Expected 3.937007874015748, but got %g", answer)
	}
}

func TestPsiToBar(t *testing.T) {
	t.Log("Converting psi to bar... (expected bar: 0.0689476)")
	answer := conversions.PsiToBar(1)

	if answer != 0.0689476 {
		t.Errorf("Expected 0.0689476, but got %g", answer)
	}
}

func TestBarToPsi(t *testing.T) {
	t.Log("Converting bar to psi... (expected psi: 14.5038)")
	answer := conversions.BarToPsi(1)

	if answer != 14.5038 {
		t.Errorf("Expected 14.5038, but got %g", answer)
	}
}

func TestCelsiusToFahrenheit(t *testing.T) {
	t.Log("Converting celsius to fahrenheit... (expected fahrenheit: 32)")
	answer := conversions.CelsiusToFahrenheit(0)

	if answer != 32 {
		t.Errorf("Expected 32, but got %g", answer)
	}
}

func TestFahrenheitToCelsius(t *testing.T) {
	t.Log("Converting fahrenheit to celsius... (expected celsius: -40")
	answer := conversions.FahrenheitToCelsius(-40)

	if answer != -40 {
		t.Errorf("Expected -40, but got %g", answer)
	}
}

func TestKelvinToCelsius(t *testing.T) {
	t.Log("Converting kelvin to celsius... (expected celsius: 0)")
	answer := conversions.KelvinToCelsius(273.15)

	if answer != 0 {
		t.Errorf("Expected 0, but got %g", answer)
	}
}

func TestCelsiusToKelvin(t *testing.T) {
	t.Log("Converting celsius to kelvin... (expected kelvin: 0)")
	answer := conversions.CelsiusToKelvin(-273.15)

	if answer != 0 {
		t.Errorf("Expected 0, but got %g", answer)
	}
}

func TestKelvinToFahrenheit(t *testing.T) {
	t.Log("Converting kelvin to fahrenheit... (expected fahrenheit: -459.67)")
	answer := conversions.KelvinToFahrenheit(0)

	if answer != -459.67 {
		t.Errorf("Expected -459.67, but got %g", answer)
	}
}

func TestFahrenheitToKelvin(t *testing.T) {
	t.Log("Converting fahrenheit to kelvin... (expected kelvin: 285.9277777777778)")
	answer := conversions.FahrenheitToKelvin(55)

	if answer != 285.9277777777778 {
		t.Errorf("Expected 285.9277777777778, but got %g", answer)
	}
}
