package logger

import (
	"go.uber.org/zap"
)

var globalLogger = (&Logger{
	config: NewConfig(zap.InfoLevel),
}).applyConfig()

func New() *Logger {
	return globalLogger.Clone()
}

func SetConfig(config *Config) {
	globalLogger.SetConfig(config)
}

func SetFields(fields map[string]string) {
	globalLogger.SetFields(fields)
}

func S() *zap.SugaredLogger {
	return globalLogger.Sugar()
}

func L() *zap.Logger {
	return globalLogger.Logger()
}

func Sync() error {
	return globalLogger.Sync()
}
