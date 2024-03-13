package app

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/realPointer/chat-server/config"
	desc "github.com/realPointer/chat-server/pkg/chat_v1"
	"github.com/realPointer/chat-server/pkg/postgres"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	desc.UnimplementedChatV1Server
}

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

	err = pg.Pool.Ping(context.Background())
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - pg.Pool.Ping: %w", err))
	}

	// GRPC
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", cfg.GRPC.Port, err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, server{})
	log.Printf("server listening at: %v", cfg.GRPC.Port)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
