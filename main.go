package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chocobone/iot_server/config"
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
