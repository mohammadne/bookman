package grpc

import (
	"net"

	"google.golang.org/grpc"
)

type server struct {
	// injected parameters
	config *Config

	// internal dependencies
	instance *grpc.Server
}

func NewServer(cfg *Config) *server {
	return &server{config: cfg, instance: grpc.NewServer()}
}

func (s *server) Serve(<-chan struct{}) error {
	listener, err := net.Listen("tcp", s.config.Url)
	if err != nil {
		return err
	}

	return s.instance.Serve(listener)
}

func (s *server) Instance() *grpc.Server {
	return s.instance
}
