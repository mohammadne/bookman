module github.com/mohammadne/bookman/user

go 1.16

require (
	github.com/go-sql-driver/mysql v1.6.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/prometheus/client_golang v1.11.0
	github.com/spf13/cobra v1.2.1
	go.opentelemetry.io/otel v1.0.0-RC3
	go.opentelemetry.io/otel/exporters/jaeger v1.0.0-RC3
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.0.0-RC3
	go.opentelemetry.io/otel/sdk v1.0.0-RC3
	go.opentelemetry.io/otel/trace v1.0.0-RC3
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	google.golang.org/genproto v0.0.0-20210909211513-a8c4777a87af // indirect
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
)
