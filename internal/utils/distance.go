package utils

import (
	"math"

	"github.com/pawlh/gfly/internal/models"
)

/* credit to cdipaolo - https://gist.github.com/cdipaolo/d3f8db3848278b49db68 */

func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 models.GPSCoordinate) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = float64(lat1 * math.Pi / 180)
	lo1 = float64(lon1 * math.Pi / 180)
	la2 = float64(lat2 * math.Pi / 180)
	lo2 = float64(lon2 * math.Pi / 180)

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}
