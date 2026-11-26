package geodesy

import "math"

// degToRad converts an angle from degrees to radians.
func degToRad(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}
