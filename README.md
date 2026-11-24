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

> Status: stage 1 in progress (geodetic → ECEF on WGS84)

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

-

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

Expected values come from pyproj (EPSG:4326 → EPSG:4978).

```
go test ./...
```

## What I learned

<!-- Fill in at the end of each stage. -->

## Limits

<!-- Fill in honestly: accuracy, what it doesn't handle. -->
