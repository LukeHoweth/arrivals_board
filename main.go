package main

import (
	"fmt"
	"os"

	f "golang.makers.tech/arrivals/flights"
)

func main() {
	// =============== collecting CLI argument ===============
	// os.Args[0] is the program name (main.go)
	// os.Args[1], os.Args[2] etc are then any passed arguments

	board, err := f.NewBoardFromJSON("flights.json")
	if err != nil {
		fmt.Println("Error creating board:", err)
		return
	}

	if len(os.Args) > 2 {
		fmt.Println("Error: Too many arguments. Please provide only one airport code (e.g., LHR)")
		return
	} else if len(os.Args) == 2 {
		// If an airport code is provided
		selectedAirport := os.Args[1]
		board.DisplayArrivalsForAirport(selectedAirport)
	} else {
		// No airport code provided, show all arrivals
		board.Display()
	}
}
