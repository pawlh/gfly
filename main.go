package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("Bad arguments. Usage: ./gfly export1.json export2.json")
		os.Exit(1)
	}

	setA, _ := process(args[0], args[1])
	save(setA, "Paul1")
	// save(setB, "Paul2")
}
