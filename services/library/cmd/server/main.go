package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/mohammadne/bookman/library/config"
	"github.com/mohammadne/bookman/library/internal/books"
	books_rest "github.com/mohammadne/bookman/library/internal/books/delivery/rest"
	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/bookman/library/pkg/logger"
	"github.com/mohammadne/bookman/library/pkg/web"
	"github.com/mohammadne/bookman/library/pkg/web/rest"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

type Server struct {
	Config *config.Config
	Logger logger.Logger
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

	db := database.NewMysqlDatabase(server.Config.Database)

	// repositories
	booksRepository := books.NewRepository(db, server.Logger)

	// usecases
	booksUsecase := books.NewUsecase(booksRepository)

	rest := rest.New(server.Config.Rest)
	v1Group := rest.Instance().Group("/api/v1")

	booksHandler := books_rest.NewHandler(booksUsecase)
	booksHandler.Route(v1Group.Group("/books"))

	go func(s web.Server) {
		server.Logger.Info("start serving rest server")
		if err := s.Serve(nil); err != nil {
			server.Logger.Panic("server failed to start", logger.Error(err))
		}
	}(rest)

	<-quit
	server.Logger.Info("Server Exited Properly")
}
