package server

import (
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}

	return cmd
}

func main() {

}
