package gcd_test

import (
	"algorithm/gcd"
	"testing"
)

var testCases = []struct {
	name     string
	a        int
	b        int
	expected int
}{
	{"Value a more than b", 48, 18, 6},
	{"Value b more than a", 101, 103, 1},
	{"Value a equals b", 100, 100, 100},
	{"Value a is zero", 0, 100, 100},
	{"Value b is zero", 100, 0, 100},
	{"Both values are zero", 0, 0, 0},
}

func TestGCDByEuclid(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			GCD := gcd.GCDByEuclid(tc.a, tc.b)
			if GCD != tc.expected {
				t.Errorf("Got wrong GCD. Expected %d, got %d", tc.expected, GCD)
			}
		})
	}
}
