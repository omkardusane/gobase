package cities_server

import (
	"encoding/json"
	"fmt"
	"net/http"
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
