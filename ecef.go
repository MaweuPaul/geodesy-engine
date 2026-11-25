package geodesy

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

	// first check the errors to return
	// 1. check if the ellipsoid is valid
	// 2. check if the latitude is in the range [-90, 90]
	// 3. check if the longitude is in the range [-180, 180]
	// 4. check if the height is a finite number
	return ECEFCoord{}, nil
}
