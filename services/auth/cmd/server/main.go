package server

import (
	"github.com/mohammadne/bookman/auth/config"
	"github.com/mohammadne/bookman/auth/internal/cache"
	"github.com/mohammadne/bookman/auth/internal/jwt"
	"github.com/mohammadne/bookman/auth/internal/web/rest"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command(cfg *config.Config, log logger.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg, log)
		},
	}
}

func main(cfg *config.Config, log logger.Logger) {
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
