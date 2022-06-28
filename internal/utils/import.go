package utils

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "sort"
    "sync"
    "time"

    "github.com/pawlh/gfly/internal/models"
)

const GroupSize = int64((time.Hour * 2) / time.Millisecond)

func Open(file1 string, file2 string) (datasetA map[int64]models.Location, datasetB map[int64]models.Location) {
    var wg sync.WaitGroup

    wg.Add(2)
    go func() {
        datasetA = process(file1)
        wg.Done()
    }()
    go func() {
        datasetB = process(file2)
        wg.Done()
    }()

    wg.Wait()
    return
}

func process(file string) map[int64]models.Location {
    contents, err := os.Open(file)
    if err != nil {
        log.Fatal("Unable to open " + fmt.Sprintf(file))
    }
    defer func(contents *os.File) {
        err := contents.Close()
        if err != nil {

        }
    }(contents)

    bytes, err := ioutil.ReadAll(contents)
    if err != nil {
        log.Fatal("Unable to read " + fmt.Sprintf(file))
    }

    var locationHistory models.GoogleLocationHistory

    err = json.Unmarshal(bytes, &locationHistory)
    if err != nil {
        log.Fatal("Json is invalid in " + fmt.Sprintf(file))
    }

    categorized := categorize(locationHistory)
    reduced := reduce(categorized)

    return reduced
}

// group locations by time blocks of duration GroupSize
func categorize(locationHistory models.GoogleLocationHistory) map[int64][]models.Location {
    groups := make(map[int64][]models.Location)

    for _, location := range locationHistory.Locations {
        groupId := location.Timestamp.Ms / GroupSize
        groups[groupId] = append(groups[groupId], location)
    }

    return groups
}

// reduce slice of locations to median of latitude and longitude
func reduce(groups map[int64][]models.Location) map[int64]models.Location {
    reduced := make(map[int64]models.Location)

    for groupId, locations := range groups {
        latitude, longitude := median(locations)
        reduced[groupId] = models.Location{
            Timestamp:        locations[0].Timestamp,
            Latitude:         latitude,
            Longitude:        longitude,
            Accuracy:         locations[0].Accuracy,
            Altitude:         locations[0].Altitude,
            VerticalAccuracy: locations[0].VerticalAccuracy,
            Activity:         locations[0].Activity,
        }
    }

    return reduced
}

// find median of latitude and longitude
func median(locations []models.Location) (models.GPSCoordinate, models.GPSCoordinate) {
    latitudes := make([]float64, len(locations))
    longitudes := make([]float64, len(locations))

    for i, location := range locations {
        latitudes[i] = float64(location.Latitude)
        longitudes[i] = float64(location.Longitude)
    }

    latitude := medianFloat(latitudes)
    longitude := medianFloat(longitudes)

    return models.GPSCoordinate(latitude), models.GPSCoordinate(longitude)
}

// determine median of float64 slice
func medianFloat(values []float64) float64 {
    sort.Float64s(values)
    n := len(values)
    if n == 0 {
        return 0
    }
    if n%2 == 0 {
        return (values[n/2-1] + values[n/2]) / 2
    }
    return values[n/2]
}
