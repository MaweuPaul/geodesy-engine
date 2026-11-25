package geodesy

import (
	"fmt"
	"math"
)

// Ellipsoid represents a reference ellipsoid defined by its semi-major axis and flattening.
type Ellipsoid struct {
	// SemiMajorAxis is the equatorial radius (a) of the ellipsoid in metres.
	SemiMajorAxis float64
	// Flattening is the dimensionless flattening factor (f = (a - b) / a), not inverse flattening (1/f).
	Flattening float64
}

// NewEllipsoid creates a new Ellipsoid after validating that semiMajorAxis is positive
// and flattening is within the range (0, 1).
func NewEllipsoid(semiMajorAxis, flattening float64) (Ellipsoid, error) {
	ellipsoid := Ellipsoid{
		SemiMajorAxis: semiMajorAxis,
		Flattening:    flattening,
	}
	if err := ValidateEllipsoid(ellipsoid); err != nil {
		return Ellipsoid{}, err
	}
	return ellipsoid, nil
}

// WGS84 returns the WGS 84 reference ellipsoid parameters (NGA TR8350.2),
// with a semi-major axis in metres and dimensionless flattening.
func WGS84() Ellipsoid {
	return Ellipsoid{
		SemiMajorAxis: 6378137,
		Flattening:    1 / 298.257223563,
	}
}

// SemiMinorAxis returns the semi-minor axis (b) of the ellipsoid in metres.
func (e Ellipsoid) SemiMinorAxis() float64 {
	return e.SemiMajorAxis * (1 - e.Flattening)
}

// EccentricitySquared returns the dimensionless first eccentricity squared (e²) of the ellipsoid.
func (e Ellipsoid) EccentricitySquared() float64 {
	return 2*e.Flattening - e.Flattening*e.Flattening
}

// Eccentricity returns the dimensionless first eccentricity (e) of the ellipsoid.
func (e Ellipsoid) Eccentricity() float64 {
	return math.Sqrt(e.EccentricitySquared())
}

func ValidateEllipsoid(e Ellipsoid) error {
	if math.IsNaN(e.SemiMajorAxis) || math.IsInf(e.SemiMajorAxis, 0) {
		return fmt.Errorf("semi-major axis must be a finite positive number, got %v", e.SemiMajorAxis)
	}
	if e.SemiMajorAxis <= 0 {
		return fmt.Errorf("semi-major axis must be positive, got %v", e.SemiMajorAxis)
	}
	if math.IsNaN(e.Flattening) || math.IsInf(e.Flattening, 0) {
		return fmt.Errorf("flattening must be a finite number, got %v", e.Flattening)
	}
	if e.Flattening <= 0 || e.Flattening >= 1 {
		return fmt.Errorf("flattening must be in the range (0, 1), got %v", e.Flattening)
	}
	return nil
}
