package main

import (
	"os"

	coreLogger "github.com/mohammadne/bookman/core/logger"
	"github.com/mohammadne/bookman/user/cmd/server"
	"github.com/mohammadne/bookman/user/config"
	"github.com/mohammadne/bookman/user/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"
	exitFailure   = 1

	use   = "bookman_user"
	short = "short"
	long  = "long long"
)

func main() {
	env := config.Development
	cfg := config.Load(env)
	log := logger.NewZap(cfg.Logger)

	root := &cobra.Command{Use: use, Short: short, Long: long}

	// register server sub-command
	serverCMD := server.Command(cfg, log)
	root.AddCommand(serverCMD)

	// run cobra root cmd
	if err := root.Execute(); err != nil {
		log.Error(errExecuteCMD, coreLogger.Error(err))
		os.Exit(exitFailure)
	}
}
