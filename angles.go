package geodesy

import "math"

// degToRad converts an angle from degrees to radians.
func degToRad(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

// radToDeg converts an angle from radians to degrees.
func radToDeg(radians float64) float64 {
	return radians * 180 / math.Pi
}
