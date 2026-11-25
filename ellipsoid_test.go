package geodesy

import (
	"math"
	"testing"
)

func TestWGS84DerivedParameters(t *testing.T) {
	// declare ell

	ell := WGS84()
	testCases := []struct {
		name      string
		got       float64
		want      float64
		tolerance float64
	}{
		{
			name:      "semi-minor axis b",
			got:       ell.SemiMinorAxis(),
			want:      6356752.3142,
			tolerance: 1e-4,
		},
		{
			name:      "eccentricity squared e²",
			got:       ell.EccentricitySquared(),
			want:      0.00669437999014,
			tolerance: 1e-14,
		},
		{
			name:      "eccentricity e",
			got:       ell.Eccentricity(),
			want:      0.0818191908426,
			tolerance: 1e-13,
		},
		{
			name:      "flattening (sanity check)",
			got:       ell.Flattening,
			want:      0.0033528106647474805,
			tolerance: 1e-16,
		},
	}

	// check if got is close enough to want, within the specified tolerance
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if math.Abs(tc.got-tc.want) > tc.tolerance {
				t.Errorf("%s: got %v, want %v (tolerance %v)", tc.name, tc.got, tc.want, tc.tolerance)
			}
		})
	}

}

func TestNewEllipsoidValidation(t *testing.T) {
	testCases := []struct {
		name          string
		semiMajorAxis float64
		flattening    float64
		wantErr       bool
	}{
		{"valid WGS84", 6378137, 1 / 298.257223563, false},
		{"zero a", 0, 1 / 298.257223563, true},
		{"negative a", -6378137, 1 / 298.257223563, true},
		{"NaN a", math.NaN(), 1 / 298.257223563, true},
		{"infinite a", math.Inf(1), 1 / 298.257223563, true},
		{"zero f (sphere)", 6378137, 0, true},
		{"inverse f mistake", 6378137, 298.257223563, true},
		{"NaN f", 6378137, math.NaN(), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := NewEllipsoid(tc.semiMajorAxis, tc.flattening)

			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("NewEllipsoid(%v, %v): err = %v, wantErr = %v",
					tc.semiMajorAxis, tc.flattening, err, tc.wantErr)
			}
		})
	}
}
