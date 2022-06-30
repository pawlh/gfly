package main

import (
	"github.com/pawlh/gfly/internal/server"
)

func main() {
	// datasetA, datasetB := utils.Open("Paul.json", "Chase.json")
	// _ = utils.FindCollisions(datasetA, datasetB)
	server.Start()
}
