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

type VacuumReturn struct {
	Validator *validator.Validate
	Token     string // Home Assistant API token
	VacuumID  string // Vacuum entity ID
	Config    *config.Config
}

func NewVacuumReturn(v *validator.Validate, token, vacuumID string, cfg *config.Config) *VacuumReturn {
	return &VacuumReturn{
		Validator: v,
		Token:     token,
		VacuumID:  vacuumID,
		Config:    cfg,
	}
}

func (vs *VacuumReturn) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Send request to Home Assistant API
	haURL := fmt.Sprintf("%s/api/services/vacuum/return_to_base", vs.Config.HomeAssistantURL())
	payloadData := map[string]string{
		"entity_id": vs.VacuumID,
	}
	payload, err := json.Marshal(payloadData)
	if err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: "Failed to create request payload",
		}, http.StatusInternalServerError)
		return
	}

	haReq, err := http.NewRequest("POST", haURL, bytes.NewBuffer(payload))
	if err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: "Failed to create request to Home Assistant",
		}, http.StatusInternalServerError)
		return
	}
	haReq.Header.Set("Authorization", "Bearer "+vs.Token)
	haReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(haReq)
	if err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: "Failed to send request to Home Assistant",
		}, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		RespondJSON(r.Context(), w, ErrResponse{
			Message: "Failed to read response from Home Assistant",
		}, http.StatusInternalServerError)
		return
	}

	// Forward the response status code and body from Home Assistant
	RespondJSON(r.Context(), w, json.RawMessage(body), resp.StatusCode)
}
