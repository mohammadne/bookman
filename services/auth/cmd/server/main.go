package server

import (
	"fmt"

	"github.com/mohammadne/bookman/auth/config"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/network"
	grpc_server "github.com/mohammadne/bookman/auth/internal/network/grpc/server"
	"github.com/mohammadne/bookman/auth/internal/network/rest"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"

	// server-cmd flags usage
	envUsage = "setting environment, default is dev"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			env, err := cmd.Flags().GetString("env")
			if err != nil {
				fmt.Println(err)
			}

			main(config.EnvFromFlag(env))
		},
	}

	// set server-cmd flags
	cmd.Flags().StringP("env", "e", "", envUsage)

	return cmd
}

func main(environment config.Environment) {
	// done channel is a trick to pause main groutine
	done := make(chan struct{})

	//
	cfg := config.Load(environment)
	log := logger.NewZap(cfg.Logger)
	cache := cache.NewRedis(cfg.Cache, log)

	//
	jwt := jwt.New(cfg.Jwt, log)

	// start serving application servers
	for _, server := range []network.Server{
		rest.New(cfg.Rest, log, cache, jwt),
		grpc_server.New(cfg.Grpc, log, cache, jwt),
	} {
		go server.Serve(done)
	}

	// pause main groutine
	<-done
}
