package main

import (
	"github.com/joho/godotenv"
	"github.com/mohammadne/bookman/auth/cmd/server"
	"github.com/spf13/cobra"
)

const (
	errLoadEnv    = "Error loading .env file"
	errExecuteCMD = "failed to execute root command"

	use   = "bookman_auth"
	short = "short"
	long  = `long`
)

func main() {
	// loads environment variables from .env
	if err := godotenv.Load(); err != nil {
		panic(map[string]interface{}{"err": err, "msg": errLoadEnv})
	}

	// root subcommands
	serverCmd := server.Command()

	// create root command and add sub-commands to it
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(serverCmd)

	// run cobra root cmd
	if err := cmd.Execute(); err != nil {
		panic(map[string]interface{}{"err": err, "msg": errExecuteCMD})
	}
}
