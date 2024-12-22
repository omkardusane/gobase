package cities_server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type City struct {
	Name string `json:"name"`
}
type Distance struct {
	From     string  `json:"from"`
	To       string  `json:"to"`
	Distance float64 `json:"distance"` // in kilometers
	Duration int     `json:"duration"` // in minutes
}

var cities = []City{
	{Name: "Berlin"},
	{Name: "Munich"},
	{Name: "Hamburg"},
	{Name: "Frankfurt"},
	{Name: "Cologne"},
}
var distances = []Distance{
	{From: "Berlin", To: "Munich", Distance: 584},
	{From: "Berlin", To: "Hamburg", Distance: 289},
	{From: "Munich", To: "Frankfurt", Distance: 394},
	{From: "Frankfurt", To: "Cologne", Distance: 191},
	{From: "Hamburg", To: "Cologne", Distance: 360},
	{From: "Cologne", To: "Berlin", Distance: 575},
}

func InitServer() {
	http.HandleFunc("/cities", getCitiesHandler)
	http.HandleFunc("/distance", getDistanceHandler)
	http.HandleFunc("/addCity", addCityHandler) // New handler for POST requests

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

func getCitiesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cities)
}

func getDistanceHandler(w http.ResponseWriter, r *http.Request) {
	from := r.URL.Query().Get("from")
	to := r.URL.Query().Get("to")

	if from == "" || to == "" {
		http.Error(w, "Both 'from' and 'to' parameters are required", http.StatusBadRequest)
		return
	}

	var found *Distance
	for _, d := range distances {
		if (d.From == from && d.To == to) || (d.From == to && d.To == from) {
			found = &d
			break
		}
	}

	if found == nil {
		http.Error(w, fmt.Sprintf("Distance between %s and %s not found", from, to), http.StatusNotFound)
		return
	}

	found.Duration = (int)(found.Distance / 2.5)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(found)
}

// http://localhost:8080/distance?from=Munich&to=Berlin
// http://localhost:8080/cities

func addCityHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method. Use POST", http.StatusMethodNotAllowed)
		return
	}

	// Decode the incoming city JSON data
	var newCity City
	err := json.NewDecoder(r.Body).Decode(&newCity)
	if err != nil {
		http.Error(w, "Failed to parse city data", http.StatusBadRequest)
		return
	}

	// Add the city to the list in memory
	cities = append(cities, newCity)

	// Save the updated list of cities to a file
	err = saveCitiesToDisk()
	if err != nil {
		http.Error(w, "Failed to save cities to disk", http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newCity)
}

func saveCitiesToDisk() error {
	// Marshal cities array to JSON
	data, err := json.MarshalIndent(cities, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling cities to JSON: %v", err)
	}

	// Write the JSON data to a file using os.WriteFile
	err = os.WriteFile("cities.json", data, 0644) // Save it to `cities.json`
	if err != nil {
		return fmt.Errorf("error writing cities to file: %v", err)
	}

	return nil
}
