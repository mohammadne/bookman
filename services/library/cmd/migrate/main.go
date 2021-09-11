package migrate

import (
	"context"
	"os"

	"github.com/mohammadne/bookman/library/internal/configs"
	"github.com/mohammadne/bookman/library/internal/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "migrate"
	short = "run migration"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	previewFlag := "if set to true, it will only preview changes and doesn't execute them"
	cmd.Flags().BoolP("preview", "p", false, previewFlag)

	return cmd
}

func main(cmd *cobra.Command, args []string) {
	env := cmd.Flag("env").Value.String()
	config := configs.Migrate(env)

	lg := logger.NewZap(config.Logger)
	db, err := database.NewClient(config.Database)
	if err != nil {
		lg.Panic("", logger.Error(err))
	}

	lg.Info("following changes will be executed to database")
	if err := db.MigratePreview(context.TODO(), os.Stdout); err != nil {
		lg.Panic("error while previewing migration", logger.Error(err))
	}

	preview, err := cmd.Flags().GetBool("preview")
	if err != nil {
		lg.Panic("", logger.Error(err))
	}

	if !preview {
		if err := db.Migrate(context.TODO()); err != nil {
			lg.Panic("error while running migration", logger.Error(err))
		}

		lg.Info("schemas successfully migrated to database")
	}
}
