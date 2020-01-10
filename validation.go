package logger

import (
	"strings"

	"github.com/pkg/errors"
	"go.uber.org/zap/zapcore"
)

var stringToZapLevel = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

func ValidateLevel(level string) (zapcore.Level, error) {
	level = strings.ToLower(level)
	if l, ok := stringToZapLevel[level]; ok {
		return l, nil
	}
	return zapcore.Level(0), errors.Errorf("error undefined level %s", level)
}
