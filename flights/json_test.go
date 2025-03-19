package flights

import (
	"testing"
)

func TestGetFlightsFromJSON(t *testing.T) {
	// Use the existing flights.json file
	filePath := "/Users/lukehoweth/Projects/arrivals_board/flights.json"

	// Test successful case
	flights, err := GetFlightsFromJSON(filePath)
	if err != nil {
		t.Fatalf("GetFlightsFromJSON returned error: %v", err)
	}

	// Check if we got the expected number of flights (16 based on the provided flights.json)
	if len(flights) != 16 {
		t.Errorf("Expected 16 flights, got %d", len(flights))
	}

	// Test file not found case
	_, err = GetFlightsFromJSON("non_existent_file.json")
	if err == nil {
		t.Error("Expected error for non-existent file, got nil")
	}
}
