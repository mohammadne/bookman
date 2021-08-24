package server

import (
	"fmt"

	"github.com/mohammadne/bookman/auth/config"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/web/rest"
	"github.com/mohammadne/bookman/auth/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
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

	cmd.Flags().StringP("env", "e", "", "setting environment, default is dev")

	return cmd
}

func main(environment config.Environment) {
	cfg := config.Load(environment)
	log := logger.NewZap(cfg.Logger)

	// done channel is a trick to pause main groutine
	done := make(chan struct{})

	cache := cache.NewRedis(cfg.Cache, log)
	cache.Initialize()

	jwt := jwt.New(cfg.Jwt, log)

	// start to Handle http endpoints
	web := rest.NewEcho(cfg.Rest, log, cache, jwt)
	web.SetupRoutes()
	web.Start()

	// pause main groutine
	<-done
}
