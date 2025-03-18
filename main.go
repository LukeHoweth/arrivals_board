package main

import (
	"fmt"
	"time"

	"golang.makers.tech/arrivals/flights" // name of module + package name (which is "flights")
	// golang.makers.tech/arrivals + /flights
)

func main() {
	fmt.Println("Welcome to the Arrivals Board!")

	// create 3 flights
	flight1 := &flights.Flight{
		Code:    "BA123",
		Origin:  "Madrid",
		DueTime: time.Now().Add(2 * time.Hour),
	}

	flight2 := &flights.Flight{
		Code:    "EJ456",
		Origin:  "Paris",
		DueTime: time.Now().Add(3 * time.Hour),
	}

	flight3 := &flights.Flight{
		Code:    "RY789",
		Origin:  "Berlin",
		DueTime: time.Now().Add(4 * time.Hour),
	}

	// Print flight information
	fmt.Println(flight1)
	fmt.Println(flight2)
	fmt.Println(flight3)
}
