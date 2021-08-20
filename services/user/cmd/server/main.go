package server

import (
	"github.com/mohammadne/bookman/user/config"
	"github.com/mohammadne/bookman/user/pkg/logger"
	"github.com/spf13/cobra"
)

func Command(cfg *config.Config, log *logger.Logger) *cobra.Command {
	return &cobra.Command{
		Use:   "driverServer",
		Short: "Run driver server",
		Run: func(cmd *cobra.Command, args []string) {
			main(cfg, log)
		},
	}
}

func main(cfg *config.Config, log *logger.Logger) {

}
