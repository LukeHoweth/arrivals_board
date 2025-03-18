package flights

import (
	"testing"
	"time"

	"golang.makers.tech/arrivals/test_utils"
)

func TestBoardDisplayWithMultipleFlights(t *testing.T) {
	flights := []Flight{
		{Code: "BA123", Origin: "Madrid", DueTime: time.Date(2024, 1, 1, 16, 23, 0, 0, time.Local)},
		{Code: "EJ456", Origin: "Paris", DueTime: time.Date(2024, 1, 1, 17, 24, 0, 0, time.Local)},
		{Code: "RY789", Origin: "Berlin", DueTime: time.Date(2024, 1, 1, 18, 20, 0, 0, time.Local)},
	}

	expected := "Time:\tFrom:\tCode:\n" +
		"16:23\tMadrid\tBA123\n" +
		"17:24\tParis\tEJ456\n" +
		"18:20\tBerlin\tRY789"

	board := Board{Flights: flights}

	// Test and record/capture termainl output
	recording := test_utils.StartRecording()
	board.Display()
	result := test_utils.EndRecording(recording)

	if result != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}
func TestBoardDisplayWithSingleFlight(t *testing.T) {
	flights := []Flight{
		{Code: "BA123", Origin: "London", DueTime: time.Date(2024, 1, 1, 12, 30, 0, 0, time.Local)},
	}

	expected := "Time:\tFrom:\tCode:\n" +
		"12:30\tLondon\tBA123"

	board := Board{Flights: flights}

	recording := test_utils.StartRecording()
	board.Display()
	result := test_utils.EndRecording(recording)

	if result != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}
func TestBoardDisplayWithNoFlights(t *testing.T) {
	flights := []Flight{}
	expected := "Time:\tFrom:\tCode:"

	board := Board{Flights: flights}

	recording := test_utils.StartRecording()
	board.Display()
	result := test_utils.EndRecording(recording)

	if result != expected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}
