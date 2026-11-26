package geodesy

import (
	"fmt"
	"math"
)

// GeodeticCoord represents a point using geodetic latitude, longitude, and ellipsoidal height.
type GeodeticCoord struct {
	LatitudeDeg  float64
	LongitudeDeg float64
	HeightM      float64
}

// ECEFCoord represents a 3D Cartesian coordinate in the Earth-Centered, Earth-Fixed frame (in meters).
type ECEFCoord struct {
	X float64
	Y float64
	Z float64
}

func ToECEF(coord GeodeticCoord, ell Ellipsoid) (ECEFCoord, error) {
	// 1. Check if the ellipsoid is valid
	if err := ValidateEllipsoid(ell); err != nil {
		return ECEFCoord{}, err
	}

	if err := ValidateGeodeticCoord(coord); err != nil {
		return ECEFCoord{}, fmt.Errorf("Error %v", err)
	}

	lat := degToRad(coord.LatitudeDeg)
	lon := degToRad(coord.LongitudeDeg)

	n := ell.PrimeVerticalRadius(lat)
	e2 := ell.EccentricitySquared()

	sinLat := math.Sin(lat)
	cosLat := math.Cos(lat)
	sinLon := math.Sin(lon)
	cosLon := math.Cos(lon)

	x := (n + coord.HeightM) * cosLat * cosLon
	y := (n + coord.HeightM) * cosLat * sinLon
	z := (n*(1-e2) + coord.HeightM) * sinLat

	return ECEFCoord{X: x, Y: y, Z: z}, nil
}

func ValidateGeodeticCoord(coord GeodeticCoord) error {

	// 2. Check if the latitude is in the range [-90, 90] (also catches NaN)
	if !(coord.LatitudeDeg >= -90 && coord.LatitudeDeg <= 90) {
		return fmt.Errorf("latitude must be between [-90, 90], got %v", coord.LatitudeDeg)
	}

	// 3. Check if the longitude is in the range [-180, 180] (also catches NaN)
	if !(coord.LongitudeDeg >= -180 && coord.LongitudeDeg <= 180) {
		return fmt.Errorf("longitude must be between [-180, 180], got %v", coord.LongitudeDeg)
	}

	// 4. Check if the height is a finite number
	if math.IsNaN(coord.HeightM) || math.IsInf(coord.HeightM, 0) {
		return fmt.Errorf("height must be a finite number, got %v", coord.HeightM)
	}

	return nil

}
