package server

import (
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/configs"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
	"github.com/mohammadne/bookman/auth/internal/network/grpc"
	"github.com/mohammadne/bookman/auth/internal/network/rest_api"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	"github.com/mohammadne/bookman/auth/pkg/tracer"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	return cmd
}

func main(cmd *cobra.Command, args []string) {
	env := cmd.Flag("env").Value.String()
	config := configs.Server(env)

	lg := logger.NewZap(config.Logger)
	tracer, err := tracer.New(config.Tracer)
	if err != nil {
		lg.Panic("error getting tracer object", logger.Error(err))
	}

	cache := cache.NewRedis(config.Cache, lg, tracer)
	jwt := jwt.New(config.Jwt, lg, tracer)
	userGrpc, err := grpc.NewUserClient(config.UserGrpc, lg, tracer)
	if err != nil {
		lg.Panic("error getting auth grpc connection", logger.Error(err))
	}

	servers := []network.Server{
		rest_api.New(config.RestApi, lg, tracer, cache, jwt, userGrpc),
		grpc.NewServer(config.AuthGrpc, lg, tracer, cache, jwt),
	}

	for _, server := range servers {
		go server.Serve(nil)
	}
}
