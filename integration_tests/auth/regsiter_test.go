package auth

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestRegisterEndToEnd(t *testing.T) {
	t.Run("Expects to sucessfully onboard a customer with valid request params", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":     "customer1@example.com",
			"password":  "123456789",
			"user_type": "customer",
		}

		payloadBytes, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "http://192.168.49.2:31001/auth/register", bytes.NewBuffer(payloadBytes))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := http.DefaultClient
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			t.Errorf("Expected status code %d but got %d", http.StatusCreated, res.StatusCode)
		}

		var response map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		if _, ok := response["status"]; !ok {
			t.Error("Response does not contain the 'status' field")
		}
	})

	t.Run("Expects to return server conflict for existing user trying to signup again", func(t *testing.T) {
		payload := map[string]interface{}{
			"email":     "customer1@example.com",
			"password":  "123456789",
			"user_type": "customer",
		}

		payloadBytes, _ := json.Marshal(payload)

		req, err := http.NewRequest("POST", "http://192.168.49.2:31001/auth/register", bytes.NewBuffer(payloadBytes))
		if err != nil {
			t.Fatalf("Failed to create request: %v", err)
		}

		req.Header.Set("Content-Type", "application/json")

		client := http.DefaultClient
		res, err := client.Do(req)
		if err != nil {
			t.Fatalf("Failed to send request: %v", err)
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusConflict {
			t.Errorf("Expected status code %d but got %d", http.StatusConflict, res.StatusCode)
		}

		var response map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&response)
		if err != nil {
			t.Fatalf("Failed to parse response body: %v", err)
		}

		if _, ok := response["status"]; !ok {
			t.Error("Response does not contain the 'status' field")
		}
		if _, ok := response["error"]; !ok {
			t.Error("Response does not contain the 'error' field")
		}
		if response["error"] != "email already registered" {
			t.Errorf("Expected error message 'email already registered' but got '%v'", response["error"])
		}
	})
}
