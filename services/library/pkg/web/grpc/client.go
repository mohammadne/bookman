package grpc

import (
	"google.golang.org/grpc"
)

type client struct {
	config *Config
}

func NewClient(cfg *Config) *client {
	return &client{config: cfg}
}

func (g *client) Setup() (*grpc.ClientConn, error) {
	return grpc.Dial(g.config.Url, grpc.WithInsecure())
}
