package gocalculation

import (
	"testing"
)

func TestCal(t *testing.T) {

	formula := "cal0 + (cal1 - cal2)/(cal3 * cal2)"
	vals := map[string]float64{
		"cal0": 12,
		"cal1": 1,
		"cal2": 4,
		"cal3": 5,
	}

	ret := Cal(formula, vals)
	if ret != 11.85 {
		t.Errorf("calculation error")
	}
}
