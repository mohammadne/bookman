package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mohammadne/bookman/library/internal/books"
	books_rest "github.com/mohammadne/bookman/library/internal/books/delivery/rest"
	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/rest"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

type Server struct {
	Logger   logger.Logger
	Database database.Database
	Rest     rest.Server
}

func (server Server) Command() *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			server.main()
		},
	}

	return cmd
}

func (server *Server) main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// repositories
	booksRepository := books.NewRepository(server.Database, server.Logger)

	// usecases
	booksUsecase := books.NewUsecase(booksRepository)

	go func() {
		server.Rest.Serve(nil)
		server.Logger.Info("start serving rest server")
	}()

	v1Group := server.Rest.Instance().Group("/api/v1")

	booksHandler := books_rest.NewHandler(booksUsecase)
	booksHandler.Route(v1Group.Group("/books"))

	<-quit
	server.Logger.Info("Server Exited Properly")
}
