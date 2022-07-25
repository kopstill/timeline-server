package logger

import (
	"go.uber.org/zap"
)

func InitLogger() {
	baseLogger, _ := zap.NewDevelopment()
	defer baseLogger.Sync()
	logger := baseLogger.Sugar()
	logger.Debug("Logger initialized.")
}
