package utils

import (
	"github.com/pawlh/gfly/internal/models"
)

// MinDistance meters
const MinDistance = 100

func FindCollisions(setA map[int64]models.Location, setB map[int64]models.Location) []models.Location {
	var collisions []models.Location
	for k := range setA {
		if _, ok := setB[k]; ok {
			if CheckDistance(k, setA, setB) {
				collisions = append(collisions, setA[k])
			}
		}
	}
	return collisions
}

// CheckDistance check distance if key exists
func CheckDistance(key int64, setA map[int64]models.Location, setB map[int64]models.Location) bool {
	distance := Distance(setA[key].Latitude, setA[key].Longitude, setB[key].Latitude, setB[key].Longitude)
	if distance < MinDistance {
		return true
	}

	return false
}
