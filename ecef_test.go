package geodesy

import (
	"math"
	"testing"
)

func TestToECEFValidation(t *testing.T) {
	testCases := []struct {
		name         string
		latitudeDeg  float64
		longitudeDeg float64
		heightM      float64
		ell          Ellipsoid
		wantErr      bool
	}{
		{"valid Nairobi", -1.2921, 36.8219, 1795, WGS84(), false},
		{"valid north pole", 90, 0, 0, WGS84(), false},
		{"valid south pole", -90, 0, 0, WGS84(), false},
		{"valid lon 180", 0, 180, 0, WGS84(), false},
		{"valid lon -180", 0, -180, 0, WGS84(), false},
		{"latitude too high", 91, 0, 0, WGS84(), true},
		{"latitude too low", -91, 0, 0, WGS84(), true},
		{"longitude too high", 0, 181, 0, WGS84(), true},
		{"longitude too low", 0, -181, 0, WGS84(), true},
		{"NaN latitude", math.NaN(), 0, 0, WGS84(), true},
		{"NaN longitude", 0, math.NaN(), 0, WGS84(), true},
		{"NaN height", 0, 0, math.NaN(), WGS84(), true},
		{"infinite height", 0, 0, math.Inf(1), WGS84(), true},
		{"invalid ellipsoid", 0, 0, 0, Ellipsoid{}, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ToECEF(GeodeticCoord{tc.latitudeDeg, tc.longitudeDeg, tc.heightM}, tc.ell)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("ToECEF(%v, %v, %v): err = %v, wantErr = %v",
					tc.latitudeDeg, tc.longitudeDeg, tc.heightM, err, tc.wantErr)
			}
		})

	}
}
