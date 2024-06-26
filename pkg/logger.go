package pkg

import (
	"github.com/sirupsen/logrus"
	"go-restful-fiber/config"
	"os"
)

func NewLogger() *logrus.Logger {
	// Get Config
	env := config.GetEnvConfig()
	appEnv := env.Get("APP_ENV")

	// Set Logger
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logLevel := logrus.TraceLevel
	logOutput := os.Stdout
	if appEnv == "production" {
		logLevel = logrus.InfoLevel
		logOutput = os.Stdout
	}
	logger.SetLevel(logLevel)
	logger.SetOutput(logOutput)

	return logger
}
