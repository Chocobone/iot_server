package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Chocobone/iot_server/entity"
	"github.com/go-playground/validator/v10"
)

type VacuumPause struct {
	Validator *validator.Validate
}

func (vs *VacuumPause) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var request struct {
		VacuumToken entity.VacuumToken `json:"vacuum_token" validate:"required"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := vs.Validator.Struct(request); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	// Send request to Home Assistant API using the service name
	haURL := "http://127.0.0.1:8123/api/services/vacuum/pause"
	haReq, err := http.NewRequest("POST", haURL, nil)
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}

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
