package cities_server_test

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"testing"
)

func TestPostCity(t *testing.T) {
	// Define the struct for the request payload
	type City struct {
		Name string `json:"name"`
	}

	// Create a City instance with the name 'Bremen'
	city := City{Name: "Bremen"}

	// Convert the city instance to JSON
	jsonData, err := json.Marshal(city)
	if err != nil {
		t.Fatalf("Error marshalling JSON: %v", err)
	}

	// Make a POST request to localhost:8080/addCity
	resp, err := http.Post("http://localhost:8080/addCity", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Fatalf("Error making POST request: %v", err)
	}
	defer resp.Body.Close()

	// Check if the response status code is 200 OK
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status 200 OK, got %v", resp.StatusCode)
	}

	log.Println("City added successfully!")
}

/*


> gobase> go test -v .\cities_server\cities_test.go
=== RUN   TestPostCity
2024/12/23 00:42:10 City added successfully!
--- PASS: TestPostCity (0.01s)
PASS
ok      command-line-arguments  0.364s

*/
