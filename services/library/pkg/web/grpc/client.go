package grpc

import (
	"google.golang.org/grpc"
)

func NewClient(cfg *Config) (*grpc.ClientConn, error) {
	return grpc.Dial(cfg.Url, grpc.WithInsecure())
}
