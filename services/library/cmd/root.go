package main

import (
	"github.com/mohammadne/bookman/library/cmd/migrate"
	"github.com/mohammadne/bookman/library/cmd/server"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	use   = "library"
	short = "short library"
	long  = "long library"
)

func main() {
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(server.Command(), migrate.Command())

	if err := cmd.Execute(); err != nil {
		panic(map[string]string{"reason": errExecuteCMD, "error": err.Error()})
	}
}

// cfg := config.Load(config.Development)
// lg := logger.NewZap(cfg.Logger)

// // root subcommands
// serverCmd := server.Server{
// 	Config: cfg,
// 	Logger: lg,
// }.Command()

// // create root command and add sub-commands to it
// cmd := &cobra.Command{Use: use, Short: short, Long: long}
// cmd.AddCommand(serverCmd)

// // run cobra root cmd
// if err := cmd.Execute(); err != nil {
// 	lg.Panic(errExecuteCMD, logger.Error(err))
// }
