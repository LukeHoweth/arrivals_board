package flights

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// FlightsData represents the structure of the flights.json file
type FlightsData struct {
	Flights  []Flight          `json:"flights"`
	Airports map[string]string `json:"airports"`
}

func GetFlightsFromJSON(filePath string) ([]Flight, error) {
	// ================= Open ===================
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// ================= Read ===================
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// ======== Unmarshal JSON to struct ========
	var flightsData FlightsData
	err = json.Unmarshal(data, &flightsData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return flightsData.Flights, nil
}

// GetAirportsFromJSON reads the airports map from the JSON file
func GetAirportsFromJSON(filePath string) (map[string]string, error) {
	// Open file
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	// Read file
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}

	// Unmarshal JSON to struct
	var flightsData FlightsData
	err = json.Unmarshal(data, &flightsData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	return flightsData.Airports, nil
}
