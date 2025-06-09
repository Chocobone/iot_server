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
	// Send request to Home Assistant API
	haURL := fmt.Sprintf("%s/api/services/vacuum/start", vs.Config.HomeAssistantURL())
	payload := []byte(`{"vacuum_id": "vacuum.robosceongsogi"}`)
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

/*
func main() {
    http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, World!"))
    })
	http.HandlerFunc("/vacuum/start", func(w http.ResponseWriter, r *http.Request){
	    url := "https://localhost:8123/api/services/vacuum/start"
		token := "my_token"
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
