package server

import (
	"github.com/mohammadne/bookman/library/internal/configs"
	"github.com/mohammadne/bookman/library/internal/database"
	"github.com/mohammadne/bookman/library/internal/network/grpc"
	"github.com/mohammadne/bookman/library/internal/network/rest_api"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/tracer"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command() *cobra.Command {
	cmd := &cobra.Command{Use: use, Short: short, Run: main}

	envFlag := "set config environment, default is dev"
	cmd.Flags().StringP("env", "e", "", envFlag)

	return cmd
}

func main(cmd *cobra.Command, args []string) {
	env := cmd.Flag("env").Value.String()
	config := configs.Server(env)

	lg := logger.NewZap(config.Logger)
	tracer, err := tracer.New(config.Tracer)
	if err != nil {
		lg.Panic("error getting tracer object", logger.Error(err))
	}

	db, err := database.NewClient(config.Database, lg, tracer)
	if err != nil {
		lg.Panic("error getting database client", logger.Error(err))
	}

	authGrpc, err := grpc.NewAuthClient(config.AuthGrpc, lg, tracer)
	if err != nil {
		lg.Panic("error getting auth grpc connection", logger.Error(err))
	}

	rest := rest_api.NewServer(config.RestApi, lg, tracer, db, authGrpc)
	rest.Serve(nil)
}
