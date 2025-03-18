package flights

import "fmt"

type Board struct {
	Flights []Flight
}

func (b Board) Display() {
	fmt.Print("Time:\tFrom:\tCode:")

	for _, flight := range b.Flights {
		fmt.Printf("\n%v\t%v\t%v", flight.DueTime.Format("15:04"), flight.Origin, flight.Code)
	}
}
