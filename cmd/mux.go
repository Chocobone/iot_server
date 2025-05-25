package main

import (
	"net/http"

	"github.com/Chocobone/iot_server/config"
	"github.com/Chocobone/iot_server/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux(cfg *config.Config) http.Handler {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})

	v := validator.New()
	mux.Handle("/vacuum/start", handler.NewVacuumStart(cfg, v))
	mux.Handle("/vacuum/pause", handler.NewVacuumPause(cfg, v))
	mux.Handle("/vacuum/return", handler.NewVacuumReturn(cfg, v))

	return mux
}
