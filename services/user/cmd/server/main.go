package server

import (
	"github.com/mohammadne/bookman/user/internal/configs"
	"github.com/mohammadne/bookman/user/internal/network"
	"github.com/mohammadne/bookman/user/internal/network/grpc"
	"github.com/mohammadne/bookman/user/internal/network/rest_api"

	"github.com/mohammadne/bookman/user/pkg/logger"
	"github.com/mohammadne/bookman/user/pkg/tracer"
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

	authGrpc, err := grpc.NewAuthClient(config.AuthGrpc, lg, tracer)
	if err != nil {
		lg.Panic("error getting auth grpc connection", logger.Error(err))
	}

	// db := database.NewMysqlDatabase(cfg.Database, lg)
	// db := interface{}

	// serving application servers
	servers := []network.Server{
		rest_api.New(config.RestApi, lg, nil, authGrpc),
		grpc.NewServer(config.UserGrpc, lg, nil),
	}

	for _, server := range servers {
		go server.Serve(nil)
	}

}
