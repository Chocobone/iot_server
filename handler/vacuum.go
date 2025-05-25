package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Chocobone/iot_server/config"
	"github.com/Chocobone/iot_server/entity"
	"github.com/go-playground/validator/v10"
)

type VacuumHandler struct {
	validator *validator.Validate
	config    *config.Config
}

func NewVacuumHandler(cfg *config.Config, v *validator.Validate) *VacuumHandler {
	return &VacuumHandler{
		validator: v,
		config:    cfg,
	}
}

func (h *VacuumHandler) sendToHomeAssistant(service string) ([]byte, error) {
	url := fmt.Sprintf("%s/api/services/vacuum/%s", h.config.Home_url, service)
	payload := entity.VacuumRequest{
		EntityID: h.config.Home_vacuum_id,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+h.config.Home_vacuum_token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %w", err)
	}

	return body, nil
}

func (h *VacuumHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var cmd entity.VacuumCommand
	if err := json.NewDecoder(r.Body).Decode(&cmd); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(cmd); err != nil {
		http.Error(w, "Validation failed", http.StatusBadRequest)
		return
	}

	body, err := h.sendToHomeAssistant(cmd.Command)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}

// VacuumStartHandler handles vacuum start commands
type VacuumStartHandler struct {
	handler *VacuumHandler
}

// VacuumPauseHandler handles vacuum pause commands
type VacuumPauseHandler struct {
	handler *VacuumHandler
}

// VacuumReturnHandler handles vacuum return commands
type VacuumReturnHandler struct {
	handler *VacuumHandler
}

func (h *VacuumStartHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func (h *VacuumPauseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func (h *VacuumReturnHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}

func NewVacuumStart(cfg *config.Config, v *validator.Validate) *VacuumStartHandler {
	return &VacuumStartHandler{
		handler: NewVacuumHandler(cfg, v),
	}
}

func NewVacuumPause(cfg *config.Config, v *validator.Validate) *VacuumPauseHandler {
	return &VacuumPauseHandler{
		handler: NewVacuumHandler(cfg, v),
	}
}

func NewVacuumReturn(cfg *config.Config, v *validator.Validate) *VacuumReturnHandler {
	return &VacuumReturnHandler{
		handler: NewVacuumHandler(cfg, v),
	}
}
