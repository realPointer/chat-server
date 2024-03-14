package grpcserver

type Option func(*Server)

// WithPort sets the port for the gRPC server.
func WithPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}
