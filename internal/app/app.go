package app

import (
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

	// GRPC
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Println(fmt.Errorf("failed to listen port %d: %w", cfg.GRPC.Port, err))

		return
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, server{})
	log.Printf("server listening at: %v", cfg.GRPC.Port)

	if err = s.Serve(lis); err != nil {
		log.Println(fmt.Errorf("failed to serve grpc: %w", err))

		return
	}
}
