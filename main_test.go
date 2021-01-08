package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	tests := []struct {
		NumberA  int
		NumberB  int
		Expected int
	}{
		{1, 2, 3},
		{100, 100, 200},
		{5, 3, 8},
	}

	for _, test := range tests {
		total := Sum(test.NumberA, test.NumberB)

		if total != test.Expected {
			t.Errorf("Returned result was incorrect, got: %d want: %d", total, test.Expected)
		}
	}
}
