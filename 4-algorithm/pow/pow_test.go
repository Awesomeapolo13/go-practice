package pow_test

import (
	"algorithm/pow"
	"testing"
)

var testCases = []struct {
	name     string
	number   int
	power    int
	expected int
}{
	{"0 to 0", 0, 0, 1},
	{"0 to 3", 0, 3, 0},
	{"3 to 0", 3, 0, 1},
	{"2 to 2", 2, 2, 4},
	{"2 to 3", 2, 3, 8},
	{"3 to 4", 2, 4, 16},
}

func TestRaiseIntToIntPositivePower(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := pow.RaiseIntToIntPositivePower(tc.number, tc.power)
			if result != tc.expected {
				t.Errorf("expected %d, got %d", tc.expected, result)
			}
		})
	}
}
