package logger

import (
	"go.uber.org/zap"
)

func (logger *Logger) With(args ...interface{}) *zap.SugaredLogger {
	return logger.Sugar().With(args...)
}

func (logger *Logger) Debug(args ...interface{}) {
	logger.Sugar().Debug(args...)
}

func (logger *Logger) Debugf(template string, args ...interface{}) {
	logger.Sugar().Debugf(template, args...)
}

func (logger *Logger) Error(args ...interface{}) {
	logger.Sugar().Error(args...)
}

func (logger *Logger) Errorf(template string, args ...interface{}) {
	logger.Sugar().Errorf(template, args...)
}

func (logger *Logger) Info(args ...interface{}) {
	logger.Sugar().Info(args...)
}

func (logger *Logger) Infof(template string, args ...interface{}) {
	logger.Sugar().Infof(template, args...)
}

func (logger *Logger) Warn(args ...interface{}) {
	logger.Sugar().Warn(args...)
}

func (logger *Logger) Warnf(template string, args ...interface{}) {
	logger.Sugar().Warnf(template, args...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	logger.Sugar().Fatal(args...)
}

func (logger *Logger) Fatalf(template string, args ...interface{}) {
	logger.Sugar().Fatalf(template, args...)
}

func (logger *Logger) Panic(args ...interface{}) {
	logger.Sugar().Panic(args...)
}

func (logger *Logger) Panicf(template string, args ...interface{}) {
	logger.Sugar().Panicf(template, args...)
}

// Global logger shortcuts
func With(args ...interface{}) *zap.SugaredLogger {
	return globalLogger.Sugar().With(args...)
}

func Debug(args ...interface{}) {
	globalLogger.Sugar().Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	globalLogger.Sugar().Debugf(template, args...)
}

func Error(args ...interface{}) {
	globalLogger.Sugar().Error(args...)
}

func Errorf(template string, args ...interface{}) {
	globalLogger.Sugar().Errorf(template, args...)
}

func Info(args ...interface{}) {
	globalLogger.Sugar().Info(args...)
}

func Infof(template string, args ...interface{}) {
	globalLogger.Sugar().Infof(template, args...)
}

func Warn(args ...interface{}) {
	globalLogger.Sugar().Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	globalLogger.Sugar().Warnf(template, args...)
}

func Fatal(args ...interface{}) {
	globalLogger.Sugar().Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	globalLogger.Sugar().Fatalf(template, args...)
}

func Panic(args ...interface{}) {
	globalLogger.Sugar().Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	globalLogger.Sugar().Panicf(template, args...)
}
