package main

import (
	"fmt"
	"time"

	f "golang.makers.tech/arrivals/flights" // Alias the import
)

func main() {
	fmt.Println("Welcome to the Arrivals Board!")

	// create a slice of Flights
	flightsList := []f.Flight{
		{
			Code:    "BA123",
			Origin:  "Madrid",
			DueTime: time.Now().Add(2 * time.Hour),
		},
		{
			Code:    "EJ456",
			Origin:  "Paris",
			DueTime: time.Now().Add(3 * time.Hour),
		},
		{
			Code:    "RY789",
			Origin:  "Berlin",
			DueTime: time.Now().Add(4 * time.Hour),
		},
	}

	arrivalBoard := f.NewBoard(flightsList)

	arrivalBoard.Display()
}
