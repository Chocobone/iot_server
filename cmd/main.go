package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/Chocobone/iot_server/internal/config"
)

func main() { //main에서 run 함수만을 수행, 실행불가 시 에러 반환
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
	
}