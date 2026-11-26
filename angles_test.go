package geodesy

import (
	"math"
	"testing"
)

func TestDegToRad(t *testing.T) {
	const epsilon = 1e-9

	tests := []struct {
		name     string
		degrees  float64
		expected float64
	}{

		// Standard / cardinal angles
		{name: "zero degrees", degrees: 0, expected: 0},
		{name: "30 degrees", degrees: 30, expected: math.Pi / 6},
		{name: "45 degrees", degrees: 45, expected: math.Pi / 4},
		{name: "60 degrees", degrees: 60, expected: math.Pi / 3},
		{name: "90 degrees", degrees: 90, expected: math.Pi / 2},
		{name: "180 degrees", degrees: 180, expected: math.Pi},
		{name: "270 degrees", degrees: 270, expected: 3 * math.Pi / 2},
		{name: "360 degrees", degrees: 360, expected: 2 * math.Pi},

		// Negative angles
		{name: "negative 90 degrees", degrees: -90, expected: -math.Pi / 2},
		{name: "negative 180 degrees", degrees: -180, expected: -math.Pi},

		// Angles beyond a full circle
		{name: "720 degrees (two revolutions)", degrees: 720, expected: 4 * math.Pi},

		// Fractional / decimal degrees
		{name: "fractional 22.5 degrees", degrees: 22.5, expected: math.Pi / 8},
		{name: "decimal 12.345 degrees", degrees: 12.345, expected: 0.21546089615869998}, // from Python math.radians(12.345)
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual := degToRad(tc.degrees)
			if math.Abs(actual-tc.expected) > epsilon {
				t.Errorf("degToRad(%v) = %v; want %v (diff > %v)",
					tc.degrees, actual, tc.expected, epsilon)
			}
		})
	}

}
