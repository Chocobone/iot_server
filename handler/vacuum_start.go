package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/Chocobone/iot_server/config"
	"github.com/Chocobone/iot_server/entity"
	"github.com/go-playground/validator/v10"
)

type VacuumStart struct {
	Validator *validator.Validate
	Token     string // Home Assistant API token
	VacuumID  string // Vacuum entity ID
	Config    *config.Config
}

func (vs *VacuumStart) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Log request body for debugging
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	// Restore the request body for later use
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// Log the request body
	fmt.Printf("Received request body: %s\n", string(bodyBytes))

	var request struct {
		VacuumToken entity.VacuumToken `json:"vacuum_token" validate:"required"`
	}

	if err := json.NewDecoder(bytes.NewBuffer(bodyBytes)).Decode(&request); err != nil {
		fmt.Printf("JSON decode error: %v\n", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := vs.Validator.Struct(request); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Prepare payload for Home Assistant API
	payload := map[string]string{
		"entity_id": vs.VacuumID,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Failed to marshal payload: %v\n", err)
		http.Error(w, "Failed to prepare request payload", http.StatusInternalServerError)
		return
	}

	// Send request to Home Assistant API
	baseURL := "http://127.0.0.1:8123"
	haURL := fmt.Sprintf("%s/api/services/vacuum/start", baseURL)
	fmt.Printf("Home Assistant base URL: %s\n", baseURL)
	fmt.Printf("Full Home Assistant URL: %s\n", haURL)
	fmt.Printf("Request payload: %s\n", string(jsonPayload))
	fmt.Printf("Token being used: %s\n", vs.Token)

	// Validate URL format
	if _, err := url.Parse(haURL); err != nil {
		fmt.Printf("Invalid URL format: %v\n", err)
		http.Error(w, fmt.Sprintf("Invalid URL format: %v", err), http.StatusInternalServerError)
		return
	}

	haReq, err := http.NewRequest("POST", haURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		fmt.Printf("Failed to create request: %v\nURL: %s\nPayload: %s\n", err, haURL, string(jsonPayload))
		http.Error(w, fmt.Sprintf("Failed to create request: %v", err), http.StatusInternalServerError)
		return
	}

	// Add required headers
	haReq.Header.Set("Authorization", "Bearer "+vs.Token)
	haReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(haReq)
	if err != nil {
		fmt.Printf("Failed to send request to Home Assistant: %v\n", err)
		http.Error(w, fmt.Sprintf("Failed to send request to Home Assistant: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read and log the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Failed to read response body: %v\n", err)
	}
	fmt.Printf("Home Assistant response status: %d\n", resp.StatusCode)
	fmt.Printf("Home Assistant response body: %s\n", string(body))

	if resp.StatusCode != http.StatusOK {
		http.Error(w, fmt.Sprintf("Home Assistant API error: %s", string(body)), resp.StatusCode)
		return
	}

	// Return success response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Vacuum cleaning started",
	})
}

/*
func main() {
    http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, World!"))
    })
	http.HandlerFunc("/vacuum/start", func(w http.ResponseWriter, r *http.Request){
	    url := "https://localhost:8123/api/services/vacuum/start"
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJkYjYwYWJkOTRlN2M0YTZjODkyMzQ3Y2JjOTgzZWUxYSIsImlhdCI6MTc0NzAyMTI5NCwiZXhwIjoyMDYyMzgxMjk0fQ.7mybkEqIh7coIRrVxkno8I1iTXCDz5wipB9rpomVUB0"
		payload := []byte(`{"vacuum_id": "vacuum.robosceongsogi"}`)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		req.Header.Set("Authorization", "Bearer "+token)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
		    panic(err)
		}
	    defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		w.Write(body)
	})
    http.ListenAndServe(":8080", nil)
} */
