package server

import (
	"github.com/pawlh/gfly/internal/models"
	"github.com/pawlh/gfly/internal/utils"
)

type point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func getPoints() []point {
	datasetA, datasetB := utils.Open("Paul.json", "Chase.json")
	collisions := utils.FindCollisions(datasetA, datasetB)

	return locationsToPoints(collisions)
}

func locationsToPoints(locations []models.Location) []point {
	points := make([]point, len(locations))
	for i, location := range locations {
		points[i] = point{Latitude: float64(location.Latitude), Longitude: float64(location.Longitude)}
	}
	return points
}
