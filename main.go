package main

import (
	"log"

	"github.com/j4ng5y/comp-time-tracker/cmd"
	"github.com/j4ng5y/comp-time-tracker/tracker"
)

func init() {
	err := tracker.InitDB()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cmd.Execute()
}
