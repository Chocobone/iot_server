package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Chocobone/iot_server/entity"
	"github.com/go-playground/validator/v10"
)

type VacuumStart struct {
	validator *validator.Validate
}

func (vs *VacuumStart) ServeHTTP(w http.ResponseWriter, r http.Request) {
	ctx := r.Context()
	
}