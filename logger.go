package logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
)


type Logger struct {
	config    *Config
	mux       sync.RWMutex
	zapLogger *zap.Logger
}

func (logger *Logger) Clone() *Logger {
	return (&Logger{
		config: logger.config.Clone(),
	}).applyConfig()
}

func (logger *Logger) SetConfig(config *Config) {
	logger.changeAndApplyConfig(func(l *Logger) { l.config = config })
}

func (logger *Logger) Sugar() *zap.SugaredLogger {
	return logger.read().zapLogger.Sugar()
}

func (logger *Logger) Logger() *zap.Logger {
	return logger.read().zapLogger
}

func (logger *Logger) SetFields(fields map[string]string) {
	logger.changeAndApplyConfig(func(l *Logger) { l.config.fields = fields })
}

func (logger *Logger) Sync() error {
	return logger.read().zapLogger.Sync()
}

func (logger *Logger) change(f func(l *Logger)) *Logger {
	logger.mux.Lock()
	f(logger)
	logger.mux.Unlock()
	return logger
}

func (logger *Logger) changeAndApplyConfig(f func(l *Logger)) {
	logger.change(f).applyConfig()
}

func (logger *Logger) read() (l *Logger) {
	logger.mux.RLock()
	defer logger.mux.RUnlock()
	return logger
}

func (logger *Logger) applyConfig() *Logger {
	logger.mux.Lock()
	defer logger.mux.Unlock()

	var err error
	logger.zapLogger, err = logger.config.zapConfig.Build(
		zap.Fields(logger.config.createZapFields()...),
	)
	if err != nil {
		log.Fatal(err)
	}
	return logger
}
