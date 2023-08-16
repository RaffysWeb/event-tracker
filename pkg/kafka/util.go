package kafka

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchSchemaFromRegistry(subject string) (string, error) {
	// Create an HTTP client
	client := &http.Client{}

	// Construct the Schema Registry API URL
	schemaRegistryAPI := fmt.Sprintf("%s/subjects/%s/versions/latest", registryURL, subject)

	// Send a GET request to the Schema Registry API
	resp, err := client.Get(schemaRegistryAPI)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Parse the response JSON to extract the schema
	var responseData struct {
		Schema string `json:"schema"`
	}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return "", err
	}

	return responseData.Schema, nil
}

// func main() {
// 	registryURL := "http://localhost:8081" // Your local Schema Registry URL
// 	subject := "your-topic-name-value"     // Replace with your topic name

// 	schema, err := fetchSchemaFromRegistry(registryURL, subject)
// 	if err != nil {
// 		log.Fatalf("Error fetching schema from registry: %v", err)
// 	}

// 	fmt.Println("Fetched schema:")
// 	fmt.Println(schema)
// }
