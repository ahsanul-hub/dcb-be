package lib

import (
	"app/dto/model"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendPaymentRequest(url string, data model.InputPaymentRequest, appKey string, bodySign string) error {
	// Encode request data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error marshalling data: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("appkey", appKey)
	req.Header.Set("bodysign", bodySign)

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
		return fmt.Errorf("received non-200 response status: %s", resp.Status)
	}

	return nil
}

func generateSMSCode() string {
	// Implement your SMS code generation logic here.  This is a placeholder.
	return "12345"
}
