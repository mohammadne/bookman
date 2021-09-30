package server

import (
	"os"
	"os/signal"
	"syscall"

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

func main(cmd *cobra.Command, _ []string) {
	env := cmd.Flag("env").Value.String()
	configs := configs.Server(env)

	lg := logger.NewZap(configs.Logger)
	tracer, err := tracer.New(configs.Tracer)
	if err != nil {
		lg.Panic("error getting tracer object", logger.Error(err))
	}

	cache := cache.NewRedis(configs.Cache, lg, tracer)
	jwt := jwt.New(configs.Jwt, lg, tracer)
	userGrpc, err := grpc.NewUserClient(configs.UserGrpc, lg, tracer)
	if err != nil {
		lg.Panic("error getting auth grpc connection", logger.Error(err))
	}

	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGINT, syscall.SIGTERM)

	servers := []network.Server{
		rest_api.New(configs.RestApi, lg, tracer, cache, jwt, userGrpc),
		grpc.NewServer(configs.AuthGrpc, lg, tracer, cache, jwt),
	}

	for _, server := range servers {
		go server.Serve()
	}

	field := logger.String("signal", (<-signalChannel).String())
	lg.Info("exiting by recieving a unix signal", field)
}
