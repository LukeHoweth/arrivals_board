package flights

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type Board struct {
	Flights []Flight
}

func (b Board) Display() {

	// Setting up header/column formatting
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	// Creating a new table
	tbl := table.New("Time", "From", "Code")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	// Adding rows to the table
	for _, flight := range b.Flights {
		tbl.AddRow(flight.DueTime.Format("15:04"), flight.Origin, flight.Code)
	}
	tbl.Print()
}
