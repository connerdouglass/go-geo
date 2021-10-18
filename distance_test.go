package geo_test

import (
	"testing"

	"github.com/connerdouglas/go-geo"
)

func TestDistanceCalculation(t *testing.T) {
	type testCase [3]geo.Distance
	testCases := []testCase{
		{geo.Meter, 100 * geo.Centimeter},
		{geo.Meter, 1000 * geo.Millimeter},
		{geo.Meter, geo.Kilometer / 1000},
		{15 * geo.Mile, 24.14015969842181 * geo.Kilometer, geo.Centimeter},
	}
	for _, tc := range testCases {
		if !tc[0].Equals(tc[1], tc[2]) {
			t.Fatalf("distances %s and %s should be equal, got unequal", tc[0], tc[1])
		}
	}
}
