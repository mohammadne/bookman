module github.com/mohammadne/bookman/user

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/mohammadne/bookman/core v0.0.0
	github.com/spf13/cobra v1.2.1
	go.uber.org/zap v1.19.0
)

replace github.com/mohammadne/bookman/core => ../core/golang
