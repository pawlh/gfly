package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func save(data map[int64][]Locations, owner string) {
	// Create the database handle, confirm driver is present
	db, err := sql.Open("mysql", "gfly:PASSWORD@tcp(localhost:3306)/GFLY")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	for binId, bin := range data {
		for _, location := range bin {
			_, err := db.Exec(
				"INSERT INTO locations(bin, owner, latitude, longitude, accuracy, timestamp) VALUES(?,?,?,?,?,?)",
				binId, owner, location.Latitude, location.Longitude, location.Accuracy, location.Timestamp.time)
			if err != nil {
				log.Fatalln(err)
			}
		}

	}
}
