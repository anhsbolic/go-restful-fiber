package helper

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger() *logrus.Logger {
	env := "development"
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logLevel := logrus.TraceLevel
	logOutput := os.Stdout
	if env == "production" {
		logLevel = logrus.InfoLevel
		logOutput = os.Stdout
	}
	logger.SetLevel(logLevel)
	logger.SetOutput(logOutput)

	return logger
}
