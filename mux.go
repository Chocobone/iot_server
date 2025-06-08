package main

import (
	"fmt"
	"net/http"

	"github.com/Chocobone/iot_server/config"
	"github.com/Chocobone/iot_server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	cfg, err := config.New()
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %v", err))
	}

	v := validator.New()
	token := config.GetSecret()
	vacuumID := config.GetVacuumID()
	vs := handler.NewVacuumStart(v, token, vacuumID, cfg)
	mux.Post("/vacuum/start", vs.ServeHTTP)
	vp := &handler.VacuumPause{Validator: v}
	mux.Post("/vacuum/pause", vp.ServeHTTP)
	vr := &handler.VacuumReturn{Validator: v}
	mux.Post("/vacuum/return", vr.ServeHTTP)

	return mux
}
