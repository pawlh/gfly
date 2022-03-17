package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func save() {
    // Create the database handle, confirm driver is present
    db, err := sql.Open("mysql", "gfly:PASSWORD@tcp(localhost:3306)/GFLY")
    if err != nil {
        log.Fatalln(err)
    }
    defer db.Close()

    // Connect and check the server version
    var version string
    db.QueryRow("SELECT VERSION()").Scan(&version)
    fmt.Println("Connected to:", version)
}
