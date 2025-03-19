package flights

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Board struct {
	flights  []Flight
	airports map[string]string // Map of airport codes to names
}

// NewBoard creates a new Board with the provided flights and airport mapping
func NewBoard(flights []Flight, airports map[string]string) Board {
	return Board{
		flights:  flights,
		airports: airports,
	}
}

// NewBoardFromJSON creates a new Board by reading both flights and airports from a JSON file
func NewBoardFromJSON(jsonPath string) (Board, error) {
	flights, err := GetFlightsFromJSON(jsonPath)
	if err != nil {
		return Board{}, err
	}

	airports, err := GetAirportsFromJSON(jsonPath)
	if err != nil {
		return Board{}, err
	}

	return NewBoard(flights, airports), nil
}

// AddFlight adds a flight to the board
func (b *Board) AddFlight(flight Flight) {
	b.flights = append(b.flights, flight)
}

// GetFlights returns a copy of the flights slice
func (b Board) GetFlights() []Flight {
	result := make([]Flight, len(b.flights))
	copy(result, b.flights)
	return result
}

func (b Board) Display() {
	// Setting up header/column formatting
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	// Creating a new table
	tbl := table.New("Time", "From", "Code", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Adding rows to the table
	for _, flight := range b.flights {
		tbl.AddRow(flight.DueTime.Format("15:04"), flight.Origin, flight.Code, flight.GetStatus())
	}
	tbl.Print()
}

func (b Board) DisplayArrivalsForAirport(airportCode string) {
	// Setting up header/column formatting
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	// Creating a new table
	tbl := table.New("Time", "From", "Code", "Status")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Convert airport code to full name if it exists in the mapping
	airportName, exists := b.airports[airportCode]
	if !exists {
		fmt.Print("Error: Airport code not found")
		return
	}

	// Adding rows to the table for flights to this airport
	for _, flight := range b.flights {
		if flight.Destination == airportName {
			tbl.AddRow(flight.DueTime.Format("15:04"), flight.Origin, flight.Code, flight.GetStatus())
		}
	}
	tbl.Print()
}
