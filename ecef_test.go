package geodesy

import (
	"math"
	"testing"
)

func almostEqual(got, want, tol float64) bool {
	return math.Abs(got-want) <= tol
}

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
			_, err := ToECEF(GeodeticCoord{LatitudeDeg: tc.latitudeDeg, LongitudeDeg: tc.longitudeDeg, HeightM: tc.heightM}, tc.ell)
			gotErr := err != nil
			if gotErr != tc.wantErr {
				t.Errorf("ToECEF(%v, %v, %v): err = %v, wantErr = %v",
					tc.latitudeDeg, tc.longitudeDeg, tc.heightM, err, tc.wantErr)
			}
		})
	}
}

func TestToECEF(t *testing.T) {
	tests := []struct {
		name  string
		lat   float64
		lon   float64
		h     float64
		wantX float64
		wantY float64
		wantZ float64
	}{
		// hand-checkable: exact values from a and b
		{"equator prime meridian", 0, 0, 0, 6378137, 0, 0},
		{"equator 90 E", 0, 90, 0, 0, 6378137, 0},
		{"equator 90 W", 0, -90, 0, 0, -6378137, 0},
		{"equator 180", 0, 180, 0, -6378137, 0, 0},
		{"north pole", 90, 0, 0, 0, 0, 6356752.314245},
		{"south pole", -90, 0, 0, 0, 0, -6356752.314245},
		{"equator 1000 m up", 0, 0, 1000, 6379137, 0, 0},

		// general points (computed independently; verify with an online converter)
		{"45 N 45 E", 45, 45, 0, 3194419.1451, 3194419.1451, 4487348.4089},
		{"Nairobi", -1.2921, 36.8219, 1795, 5105860.6371, 3822716.4445, -142901.6315},
		{"Sydney", -33.8688, 151.2093, 58, -4646093.4773, 2553229.5358, -3534404.7109},
	}

	const tolerance = 1e-3 // 1 mm: expected values are rounded to 0.1 mm
	ell := WGS84()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := ToECEF(GeodeticCoord{LatitudeDeg: tc.lat, LongitudeDeg: tc.lon, HeightM: tc.h}, ell)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !almostEqual(got.X, tc.wantX, tolerance) ||
				!almostEqual(got.Y, tc.wantY, tolerance) ||
				!almostEqual(got.Z, tc.wantZ, tolerance) {
				t.Errorf("ToECEF(%v, %v, %v) = (%.6f, %.6f, %.6f), want (%.6f, %.6f, %.6f) within tolerance %g",
					tc.lat, tc.lon, tc.h,
					got.X, got.Y, got.Z,
					tc.wantX, tc.wantY, tc.wantZ,
					tolerance)
			}
		})
	}
}
