package flights

import (
	"testing"
	"time"
)

func TestFlightStringifiesCorrectly(t *testing.T) {
	parsedTime, err := time.Parse("15:04", "14:00") // "15:04" is Go's time format for HH:MM, how strange?
	if err != nil {
		t.Fatal("Failed to parse time:", err)
	}

	flight := Flight{
		Code:    "BA 341",
		Origin:  "London",
		DueTime: parsedTime,
	}

	result := flight.String()
	expected := "Flight BA 341 from London is expected at 14:00"

	if result != expected {
		t.Errorf("\nExpected: %s\nGot: %s", expected, result)
	}
}
