package server

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/labstack/echo/v4"
	"github.com/mohammadne/bookman/library/internal/books"
	"github.com/mohammadne/bookman/library/internal/books/delivery/rest"
	"github.com/mohammadne/bookman/library/pkg/database"
	"github.com/mohammadne/go-pkgs/logger"
	"github.com/spf13/cobra"
)

const (
	use   = "server"
	short = "run server"
)

func Command(logger logger.Logger, db database.Database) *cobra.Command {
	cmd := &cobra.Command{
		Use:   use,
		Short: short,
		Run: func(cmd *cobra.Command, args []string) {
			main(logger, db)
		},
	}

	return cmd
}

func main(logger logger.Logger, db database.Database) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// repositories
	booksRepository := books.NewRepository(db)

	// usecases
	booksUsecase := books.NewUsecase(booksRepository)

	e := echo.New()
	v1Group := e.Group("/api/v1")

	// healthGroup := v1Group.Group("/health")

	booksGroup := v1Group.Group("/books")
	booksHandler := rest.NewHandler(booksUsecase)
	rest.Route(booksGroup, booksHandler)

	// servers := []network.Server{
	// 	rest.New(cfg.Rest, log, db, authGrpc),
	// 	grpc.NewServer(cfg.GrpcServer, log, db),
	// }

	// for _, server := range servers {
	// 	go server.Serve(done)
	// }

	<-quit
	logger.Info("Server Exited Properly")
}
