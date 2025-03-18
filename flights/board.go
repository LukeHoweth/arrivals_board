package flights

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Board struct {
	flights []Flight // Private field (lowercase)
}

// NewBoard creates a new Board with the provided flights
func NewBoard(flights []Flight) Board {
	return Board{
		flights: flights,
	}
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
	tbl := table.New("Time", "From", "Code")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Adding rows to the table
	for _, flight := range b.flights {
		tbl.AddRow(flight.DueTime.Format("15:04"), flight.Origin, flight.Code)
	}
	tbl.Print()
}
