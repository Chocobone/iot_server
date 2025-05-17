package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	v := validator.New()
	mux.Handle("/vacuum/start", &handler.AddTask{Validator: v})
	mux.Handle("/vacuum/status", &handler.GetStatus{Validator: v})
	mux.Handle("/vacuum/pause", &handler.Pause{Validator: v})
	mux.Handle("/vacuum/return", &handler.Return{Validator: v})

	return mux
}
