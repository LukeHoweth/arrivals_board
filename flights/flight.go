package flights

import (
	"fmt"
	"time"
)

type Flight struct {
	Code        string       `json:"code"`
	Origin      string       `json:"from"`
	Destination string       `json:"to"`
	DueTime     time.Time    `json:"scheduled_arrival"`
	Status      FlightStatus `json:"status"`
}

type FlightStatus struct {
	Arrived    string `json:"arrived"`
	Cancelled  bool   `json:"cancelled"`
	ExpectedAt string `json:"expected_at"`
}

func (f Flight) String() string {
	return fmt.Sprintf("Flight %s from %s is expected at %s", f.Code, f.Origin, f.DueTime.Format("15:04"))
}

func (f Flight) GetStatus() string {
	if f.Status.Cancelled {
		return "Cancelled"
	}

	// Parse the time from the Arrived string
	arrivedTime, err := time.Parse(time.RFC3339, f.Status.Arrived)
	if err == nil {
		return fmt.Sprintf("Landed %s", arrivedTime.Format("15:04"))
	}

	// Parse the time from the ExpectedAt string
	expectedTime, err := time.Parse(time.RFC3339, f.Status.ExpectedAt)
	if err == nil {
		return fmt.Sprintf("Expected %s", expectedTime.Format("15:04"))
	}
	return "Unknown"
}
