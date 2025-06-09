package handler

import (
	"bytes"
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
	haURL := fmt.Sprintf("%s/api/services/vacuum/Return", vs.Config.HomeAssistantURL())
	payload := []byte(`{"vacuum_id": "vacuum.robosceongsogi"}`)
	haReq, err := http.NewRequest("POST", haURL, bytes.NewBuffer(payload))
	if err != nil {
		http.Error(w, "Failed to create request", http.StatusInternalServerError)
		return
	}
	haReq.Header.Set("Authorization", "Bearer "+vs.Token)
	haReq.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(haReq)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	w.Write(body)

}
