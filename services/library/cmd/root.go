package main

import (
	"github.com/mohammadne/bookman/library/cmd/server"
	"github.com/mohammadne/bookman/library/config"
	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/rest"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	use   = "bookman_auth"
	short = "short"
	long  = `long`
)

func main() {
	cfg := config.Load(config.Development)
	logger := logger.NewZap(cfg.Logger)
	database := database.NewMysqlDatabase(cfg.Database)
	rest := rest.New(cfg.Rest)

	// root subcommands
	serverCmd := server.Server{
		Logger:   logger,
		Database: database,
		Rest:     rest,
	}.Command()

	// create root command and add sub-commands to it
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(serverCmd)

	// run cobra root cmd
	if err := cmd.Execute(); err != nil {
		panic(map[string]interface{}{"err": err, "msg": errExecuteCMD})
	}
}
