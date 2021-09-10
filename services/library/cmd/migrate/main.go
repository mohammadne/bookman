package migrate

import (
	"github.com/spf13/cobra"
)

const (
	use   = "migrate"
	short = "run migration"
)

func Command() *cobra.Command {
	return &cobra.Command{Use: use, Short: short, Run: main}
}

func main(cmd *cobra.Command, args []string) {}
