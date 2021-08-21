package server

import (
	"github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/config"
	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/web"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command(cfg *config.Config, log *logger.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg, log)
		},
	}
}

func main(cfg *config.Config, log *logger.Logger) {
	// done channel is a trick to pause main groutine
	done := make(chan struct{})

	db := database.NewMysqlDatabase(cfg.Database, log)

	// start to Handle http endpoints
	web := web.NewEcho(cfg.Web, log, &db)
	web.StartG()

	// pause main groutine
	<-done
}
