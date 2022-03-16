package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

const GROUP_SIZE = time.Hour * 2

func process(file1 string, file2 string) (datasetA map[int][]Locations, datasetB map[int][]Locations) {
	var wg sync.WaitGroup

	wg.Add(2)
	go func() {
		datasetA = open(file1)
		wg.Done()
	}()
	go func() {
		datasetB = open(file2)
		wg.Done()
	}()
	wg.Wait()

	return
}

func open(file string) map[int][]Locations {
	contents, err := os.Open(file)
	if err != nil {
		log.Fatal("Unable to open " + fmt.Sprintf(file))
	}
	defer contents.Close()

	bytes, err := ioutil.ReadAll(contents)
	if err != nil {
		log.Fatal("Unable to read " + fmt.Sprintf(file))
	}

	var locationHistory GoogleLocationHistory

	err = json.Unmarshal(bytes, &locationHistory)
	if err != nil {
		log.Fatal("Json is invalid in " + fmt.Sprintf(file))
	}

	return categorize(locationHistory)

}

func categorize(locationHistory GoogleLocationHistory) map[int][]Locations {
	groups := make(map[int][]Locations)

	for _, location := range locationHistory.Locations {
		groupId := int(location.Timestamp.Duration / GROUP_SIZE)
		groups[groupId] = append(groups[groupId], location)
	}

	return groups
}
