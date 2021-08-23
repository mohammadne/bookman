module github.com/mohammadne/bookman/auth

go 1.16

require (
	github.com/go-redis/redis/v8 v8.11.3
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/joho/godotenv v1.3.0
	github.com/labstack/echo/v4 v4.5.0
	github.com/mohammadne/bookman/core/failures v0.0.0
	github.com/mohammadne/bookman/core/logger v0.0.0
	github.com/satori/go.uuid v1.2.0 // indirect
	go.uber.org/zap v1.19.0
)

replace (
	github.com/mohammadne/bookman/core/failures => ../core/golang/failures
	github.com/mohammadne/bookman/core/logger => ../core/golang/logger
)
