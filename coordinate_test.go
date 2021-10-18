package geo_test

import (
	"testing"

	"github.com/connerdouglass/go-geo"
)

func TestDistanceBetween(t *testing.T) {
	type testCase struct {
		a, b geo.Coordinate
		dist geo.Distance
	}
	testCases := []testCase{
		{geo.Coordinate{0, 0}, geo.Coordinate{0, 0}, 0},
		{geo.Coordinate{33.696239, -117.307427}, geo.Coordinate{36.614465, -110.369171}, 709079.58 * geo.Meter},
		{geo.Coordinate{47.368632, -106.153079}, geo.Coordinate{46.4053467, -105.8556375}, 109471.30 * geo.Meter},
		{geo.Coordinate{40.389707, -85.838675}, geo.Coordinate{40.380558, -85.826437}, 1452.36 * geo.Meter},
		{geo.Coordinate{40.234460, -85.560556}, geo.Coordinate{13.936424, -13.766256}, 7475474.67 * geo.Meter},
	}
	for _, tc := range testCases {
		dist := geo.DistanceBetween(tc.a, tc.b)
		// Check for correct distance calculation, with a tolerance of 1 meter
		if !dist.Equals(tc.dist, geo.Meter) {
			t.Fatalf("distance between %s and %s should be %s, got %s", tc.a, tc.b, tc.dist, dist)
		}
	}
}
