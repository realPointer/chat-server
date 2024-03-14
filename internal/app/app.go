package app

import (
	"fmt"
	"log"

	"github.com/realPointer/chat-server/config"

	grpcserver "github.com/realPointer/chat-server/pkg/grpc_server"
	"github.com/realPointer/chat-server/pkg/postgres"
)

func Run() {
	// Config
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("config error: %v", err)
	}

	// Postgres
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// GRPC
	grpcServer, err := grpcserver.New(grpcserver.WithPort(cfg.GRPC.Port))
	if err != nil {
		log.Printf("app - Run - grpcserver.New: %v", err)
	}

	err = grpcServer.Start()
	if err != nil {
		log.Printf("app - Run - grpcServer.Start: %v", err)
	}
}
