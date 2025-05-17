package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Chocobone/iot_server/config"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Printf("failed to terminate server: %v", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	cfg, err := config.New()
	if err != nil {
		return err
	}
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		return err
	}
	url := fmt.Sprintf("http://%s", l.Addr().String())
	log.Printf("start with: %s", url)
	mux := NewMux()
	s := NewServer(l, mux)
	return s.Run(ctx)
}

/*
func main() {
    http.HandlerFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("Hello, World!"))
    })
	http.HandlerFunc("/vacuum/start", func(w http.ResponseWriter, r *http.Request){
	    url := "https://localhost:8123/api/services/vacuum/start"
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJkYjYwYWJkOTRlN2M0YTZjODkyMzQ3Y2JjOTgzZWUxYSIsImlhdCI6MTc0NzAyMTI5NCwiZXhwIjoyMDYyMzgxMjk0fQ.7mybkEqIh7coIRrVxkno8I1iTXCDz5wipB9rpomVUB0"
		payload := []byte(`{"entity_id": "vacuum.robosceongsogi"}`)

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
