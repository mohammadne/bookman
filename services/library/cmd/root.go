package main

import (
	"github.com/mohammadne/bookman/library/cmd/server"
	"github.com/mohammadne/bookman/library/config"
	"github.com/mohammadne/bookman/library/pkg/logger"
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
	lg := logger.NewZap(cfg.Logger)

	// root subcommands
	serverCmd := server.Server{
		Config: cfg,
		Logger: lg,
	}.Command()

	// create root command and add sub-commands to it
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(serverCmd)

	// run cobra root cmd
	if err := cmd.Execute(); err != nil {
		lg.Panic(errExecuteCMD, logger.Error(err))
	}
}
