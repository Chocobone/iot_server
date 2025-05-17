package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		_, _ = w.Write([]byte(`{"status": "ok"}`))
	})
	mux.Handle("/vacuum/start", &handler.AddTask)
	return mux

}
