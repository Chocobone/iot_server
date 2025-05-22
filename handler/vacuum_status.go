package handler

import (
	"encoding/json"
	"net/http"

	//"time"

	"github.com/Chocobone/iot_server/entity"
	"github.com/go-playground/validator/v10"
)

type VacuumStatus struct {
	Validator *validator.Validate
}

func (v *VacuumStatus) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var req entity.VacuumStatusRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := v.Validator.Struct(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
