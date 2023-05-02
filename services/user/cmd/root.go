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
	cmd := &cobra.Command{Use: use, Short: short, Long: long}
	cmd.AddCommand(server.Command())

	if err := cmd.Execute(); err != nil {
		panic(map[string]interface{}{"err": err, "msg": errExecuteCMD})
	}
}
