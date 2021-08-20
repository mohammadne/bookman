package main

import (
	"log"

	"github.com/mohammadne/bookman/user/pkg/logger"
)

func main() {
	log.Println("Starting api server")

	appLogger := logger.NewApiLogger(nil)
	appLogger.InitLogger()

}
