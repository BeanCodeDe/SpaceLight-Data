package main

import (
	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/api"
	"github.com/BeanCodeDe/SpaceLight-Data/internal/app/data/config"
	"github.com/BeanCodeDe/authi/pkg/middleware"
	"github.com/BeanCodeDe/authi/pkg/parser"
	log "github.com/sirupsen/logrus"
)

//const rootPath = "/spacelight"

func main() {
	setLogLevel(config.LogLevel)
	log.Info("Start Server")
	_, err := api.NewApi()
	if err != nil {
		log.Fatal("Error while starting api: %w", err)
	}
}

func setLogLevel(logLevel string) {
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	default:
		log.SetLevel(log.DebugLevel)
		log.Errorf("Log level %s unknow", logLevel)
	}

}

func initAuthMiddleware() middleware.Middleware {
	tokenParser, err := parser.NewJWTParser()
	if err != nil {
		log.Fatalf("Error while init auth middleware: %v", err)
		return nil
	}
	return middleware.NewEchoMiddleware(tokenParser)
}
