package server

import (
	"fmt"

	"github.com/mohammadne/bookman/user/config"
	"github.com/mohammadne/bookman/user/internal/database"
	"github.com/mohammadne/bookman/user/internal/web/rest"
	"github.com/mohammadne/bookman/user/pkg/logger"
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

	//
	log := logger.NewZap(cfg.Logger)

	db := database.NewMysqlDatabase(cfg.Database, log)

	// start to Handle http endpoints
	web := rest.NewEcho(cfg.Rest, log, db)
	web.StartG()

	// pause main groutine
	<-done
}
