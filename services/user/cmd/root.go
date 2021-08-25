package main

import (
	"github.com/mohammadne/bookman/user/cmd/server"
	"github.com/spf13/cobra"
)

const (
	errExecuteCMD = "failed to execute root command"

	use   = "bookman_user"
	short = "short"
	long  = "long long"
)

func main() {
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
