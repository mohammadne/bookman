package server

import (
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/configs"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
	"github.com/mohammadne/bookman/auth/internal/network/grpc"
	grpc_client "github.com/mohammadne/bookman/auth/internal/network/grpc/clients"
	"github.com/mohammadne/bookman/auth/internal/network/rest_api"
	"github.com/mohammadne/bookman/auth/pkg/logger"
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

	// done channel is a trick to pause main groutine
	done := make(chan struct{})

	lg := logger.NewZap(config.Logger)

	cache := cache.NewRedis(config.Cache, lg)

	//
	jwt := jwt.New(config.Jwt, lg)

	userGrpc := grpc_client.NewUser(config.UserGrpc, lg)
	userGrpc.Setup()

	// serving application servers
	servers := []network.Server{
		rest_api.New(config.RestApi, lg, cache, jwt, userGrpc),
		grpc.NewServer(config.AuthGrpc, lg, cache, jwt),
	}

	for _, server := range servers {
		go server.Serve(done)
	}

	// pause main groutine
	<-done
}
