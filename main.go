package main

import (
	"fmt"

	f "golang.makers.tech/arrivals/flights"
)

func main() {
	flights, err := f.GetFlightsFromJSON("flights.json")
	if err != nil {
		fmt.Printf("Error getting flights from JSON: %v\n", err)
		return
	}

	board := f.NewBoard(flights)
	board.Display()
}
