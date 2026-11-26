# geodesy-engine

A geodesy library written from scratch in Go: the core mathematics of geodesy, implemented by hand and checked against pyproj.

This is a learning project. PROJ already solves these problems; the goal here is to implement the mathematics myself, understand every step, and verify the results against pyproj to the millimetre.

## What it covers

Geodesy is the foundation of geomatics: every survey, map and GIS layer depends on knowing the shape of the Earth and which reference frame a coordinate belongs to. This library builds that foundation piece by piece:

- **Reference ellipsoids**: WGS84, Clarke 1880, and others
- **Coordinate conversions**: geodetic ↔ Earth-Centred, Earth-Fixed (ECEF)
- **Datum transformations**: 7-parameter Helmert, Arc 1960 → WGS84 (Kenya)
- **Map projections**: Transverse Mercator / UTM
- **Geodetic computations** (later): geodesic distance and bearing, ellipsoidal area, geoid heights

> Status: stage 1 in progress (geodetic → ECEF on WGS84). Geodetic → ECEF is done and tested to 1 mm; ECEF → geodetic is next.

## Stage 1 design: geodetic → ECEF

### Problem

GPS positions come as geodetic coordinates (latitude, longitude, ellipsoidal height), which are angles on a curved surface. Distances, datum shifts (Helmert) and satellite geometry need straight-line Cartesian coordinates instead: Earth-Centred, Earth-Fixed (ECEF) X, Y, Z in metres. This stage converts geodetic coordinates to ECEF on a given ellipsoid.

### First idea

<!-- How you plan to solve it, in one statement. -->

### Information needed

<!-- Bullet list. Not just lat/lon/h: what do you need to know about the Earth? -->

-

### Generalised

<!-- Where does "WGS84" become something broader? -->

### Functions and types

<!-- Names, inputs, outputs. No bodies. -->

**Ellipsoid** ([ellipsoid.go](ellipsoid.go))

- `Ellipsoid{SemiMajorAxis, Flattening}`: `a` in metres, `f` dimensionless (not `1/f`)
- `WGS84() Ellipsoid`: WGS 84 parameters (NGA TR8350.2)
- `NewEllipsoid(a, f float64) (Ellipsoid, error)`: builds and validates an ellipsoid
- `ValidateEllipsoid(e Ellipsoid) error`: `a` finite and positive, `f` finite and in (0, 1)
- `(e Ellipsoid) SemiMinorAxis() float64`: `b = a(1 − f)`
- `(e Ellipsoid) EccentricitySquared() float64`: `e² = 2f − f²`
- `(e Ellipsoid) Eccentricity() float64`: `e = √e²`
- `(e Ellipsoid) PrimeVerticalRadius(latRad float64) float64`: `N = a / √(1 − e²·sin²φ)`, latitude in radians, result in metres

**Coordinates** ([ecef.go](ecef.go))

- `GeodeticCoord{LatitudeDeg, LongitudeDeg, HeightM}`: degrees, degrees, metres above the ellipsoid
- `ECEFCoord{X, Y, Z}`: metres
- `ValidateGeodeticCoord(coord GeodeticCoord) error`: latitude in [−90, 90], longitude in [−180, 180], finite height; NaN fails
- `ToECEF(coord GeodeticCoord, ell Ellipsoid) (ECEFCoord, error)`: validates the ellipsoid and the coordinate, then computes
  - `X = (N + h)·cosφ·cosλ`
  - `Y = (N + h)·cosφ·sinλ`
  - `Z = (N(1 − e²) + h)·sinφ`

**Angles** ([angles.go](angles.go))

- `degToRad(degrees float64) float64`: unexported helper

### Invariants

<!-- Things that must always be true. Each one becomes a test. -->

1.
2.
3.

### Out of scope (for now)

-
-

## Stages

- [ ] 1. Geodetic ↔ ECEF on WGS84
- [ ] 2. Clarke 1880 ellipsoid
- [ ] 3. 7-parameter Helmert transformation
- [ ] 4. Arc 1960 → WGS84
- [ ] 5. Transverse Mercator / UTM
- [ ] 6. HTTP service + benchmark against pyproj

## Later (not until stages 1–6 are done)

- Geodesic distance and bearing (Vincenty)
- Ellipsoidal area
- Geoid height lookup
- Traverse adjustment (Bowditch), COGO

## Testing

Expected values come from pyproj (EPSG:4326 → EPSG:4978) and from points that can be worked out by hand (for example, lat 0 / lon 0 gives X = a; the north pole gives Z = b). Coordinates are compared with a tolerance of 1 mm.

```
go test -count=1 -v ./...
```

| Test | Checks |
|---|---|
| `TestWGS84DerivedParameters` | `b`, `e²`, `e`, `f`, and `N` at the equator, 45° and the pole for WGS84 |
| `TestNewEllipsoidValidation` | rejects zero, negative, NaN and infinite `a`; `f` of 0, NaN, or inverse flattening by mistake |
| `TestToECEFValidation` | rejects out-of-range and NaN lat/lon, non-finite height, and an invalid ellipsoid |
| `TestToECEF` | equator (4 longitudes), both poles, a height offset, 45°/45°, Nairobi and Sydney, to 1 mm |
| `TestDegToRad` | cardinal, negative, fractional and multi-revolution angles |

## What I learned

<!-- Fill in at the end of each stage. -->

## Limits

<!-- Fill in honestly: accuracy, what it doesn't handle. -->
