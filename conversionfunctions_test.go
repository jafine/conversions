package conversions

import (
	"testing"

	"github.com/davidjpeacock/conversions"
)

func TestMetresToFeet(t *testing.T) {
	t.Log("Converting metres to feet... (expected feet: 9.843)")
	answer := conversions.MetresToFeet(3)

	if answer != "9.843ft" {
		t.Errorf("Expected 9.843ft, but got %g", answer)
	}
}

func TestFeetToMetres(t *testing.T) {
	t.Log("Converting feet to metres... (expected metres: 2.743)")
	answer := conversions.FeetToMetres(9)

	if answer != "2.743m" {
		t.Errorf("Expected 2.743m, but got %g", answer)
	}
}
