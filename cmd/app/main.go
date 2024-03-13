package main

import (
	"fmt"
	"log"
	"net"

	desc "github.com/realPointer/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	grpcPort = 50052
)

type server struct {
	desc.UnimplementedChatV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen port %d: %v", grpcPort, err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, server{})
	log.Printf("server listening at: %v", grpcPort)

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve grpc: %v", err)
	}
}
