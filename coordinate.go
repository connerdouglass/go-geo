package geo

import (
	"fmt"
	"math"
	"strconv"
)

// Coordinate represents a specific location on Earth
type Coordinate struct {
	Latitude, Longitude float64
}

// Constants needed for distance calculations
const (
	EarthRadius       = 6371 * Kilometer
	DoubleEarthRadius = 2 * EarthRadius
	PiOver180         = math.Pi / 180
)

// DistanceBetween calculates the distance between two coordinates
func DistanceBetween(a, b Coordinate) Distance {
	value := 0.5 - math.Cos((b.Latitude-a.Latitude)*PiOver180)/2 + math.Cos(a.Latitude*PiOver180)*math.Cos(b.Latitude*PiOver180)*(1-math.Cos((b.Longitude-a.Longitude)*PiOver180))/2
	return DoubleEarthRadius * Distance(math.Asin(math.Sqrt(value)))
}

// DistanceTo calculates the distance from this coordinate to another coordinate
func (c Coordinate) DistanceTo(other Coordinate) Distance {
	return DistanceBetween(c, other)
}

// String implements Stringer, returns a string representation of the coordinate
func (c Coordinate) String() string {
	return fmt.Sprintf(
		"(%s, %s)",
		strconv.FormatFloat(c.Latitude, 'f', -1, 64),
		strconv.FormatFloat(c.Longitude, 'f', -1, 64),
	)
}
