package conversions_test

import (
	"testing"

	"github.com/davidjpeacock/conversions"
)

func TestMetresToFeet(t *testing.T) {
	t.Log("Converting metres to feet... (expected feet: 9.843)")
	answer := conversions.MetreToFoot(3)

	if answer != 9.843 {
		t.Errorf("Expected 9.843, but got %g", answer)
	}
}

func TestFeetToMetres(t *testing.T) {
	t.Log("Converting feet to metres... (expected metres: 2.7430661383724475)")
	answer := conversions.FootToMetre(9)

	if answer != 2.7430661383724475 {
		t.Errorf("Expected 2.7430661383724475, but got %g", answer)
	}
}
