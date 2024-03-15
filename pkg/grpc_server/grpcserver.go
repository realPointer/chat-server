package grpcserver

import (
	"fmt"
	"log"
	"net"

	"github.com/realPointer/chat-server/internal/api/chat"
	pb "github.com/realPointer/chat-server/pkg/chat_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	_defaultPort = 50051
)

type Server struct {
	pb.UnimplementedChatV1Server
	GrpcServer *grpc.Server
	port       int
}

func New(opts ...Option) (*Server, error) {
	server := &Server{
		GrpcServer: grpc.NewServer(),
		port:       _defaultPort,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server, nil
}

func (s *Server) Start(chat *chat.Implementation) error {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	log.Printf("server listening at: %v", s.port)
	reflection.Register(s.GrpcServer)
	pb.RegisterChatV1Server(s.GrpcServer, chat)

	if err := s.GrpcServer.Serve(lis); err != nil {
		return fmt.Errorf("failed to serve: %w", err)
	}

	return nil
}
