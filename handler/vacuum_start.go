package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Chocobone/iot_server/config"
	//"github.com/Chocobone/iot_server/entity"
	"github.com/go-playground/validator/v10"
)

type VacuumStart struct {
	Validator *validator.Validate
	Token     string // Home Assistant API token
	VacuumID  string // Vacuum entity ID
	Config    *config.Config
}

func NewVacuumStart(v *validator.Validate, token, vacuumID string, cfg *config.Config) *VacuumStart {
	return &VacuumStart{
		Validator: v,
		Token:     token,
		VacuumID:  vacuumID,
		Config:    cfg,
	}
}

func (vs *VacuumStart) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	payload := map[string]string{
		"vacuum_id": vs.VacuumID,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to prepare request payload", http.StatusInternalServerError)
		return
	}

	// Send request to Home Assistant API
	haURL := fmt.Sprintf("%s/api/services/vacuum/start", vs.Config.HomeAssistantURL())
	haReq, err := http.NewRequest("POST", haURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	// Add required headers
	haReq.Header.Set("Authorization", "Bearer "+vs.Token)
	haReq.Header.Set("content-type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(haReq)
	if err != nil {
		http.Error(w, "Failed to send request to Home Assistant", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
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
		"vacuum_id": vs.VacuumID,
	}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "Failed to prepare request payload", http.StatusInternalServerError)
		return
	}

	// Send request to Home Assistant API
	haURL := fmt.Sprintf("%s/api/services/vacuum/start", vs.Config.HomeAssistantURL())
	haReq, err := http.NewRequest("POST", haURL, bytes.NewBuffer(jsonPayload))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	// Add required headers
	haReq.Header.Set("Authorization", "Bearer "+vs.Token)
	haReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(haReq)
	if err != nil {
		http.Error(w, "Failed to send request to Home Assistant", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
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
*/

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
