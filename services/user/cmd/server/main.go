package server

import (
	"github.com/mohammadne/bookman/user/config"
	"github.com/mohammadne/bookman/user/pkg/logger"
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

}
