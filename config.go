package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Config logger config
type Config struct {
	fields    map[string]string
	zapFields []zap.Field

	zapConfig zap.Config
}

func NewConfig(level zapcore.Level) *Config {
	return &Config{
		zapConfig: NewDevelopmentConfig(level),
	}
}

func (c *Config) SetZapConfig(cfg zap.Config) *Config {
	c.zapConfig = cfg
	return c
}

func (c *Config) AddZapFields(fields ...zap.Field) *Config {
	c.zapFields = append(c.zapFields, fields...)
	return c
}

func (c *Config) SetFields(fields map[string]string) *Config {
	c.fields = fields
	return c
}

func (c *Config) Clone() *Config {
	var newConfig = &Config{
		zapConfig: c.zapConfig,
		fields:    make(map[string]string, len(c.fields)),
		zapFields: make([]zap.Field, len(c.zapFields)),
	}

	for n, v := range c.fields {
		newConfig.fields[n] = v
	}
	copy(newConfig.zapFields, c.zapFields)
	return newConfig
}

func (c *Config) createZapFields() []zap.Field {
	for key, value := range c.fields {
		c.zapFields = append(c.zapFields, zap.String(key, value))
	}
	return c.zapFields
}

func NewDevelopmentConfig(level zapcore.Level) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: true,
		Encoding:    "console",
		EncoderConfig: zapcore.EncoderConfig{
			// Keys can be anything except the empty string.
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder, // Enable color output
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.StringDurationEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func NewProductionConfig(level zapcore.Level) zap.Config {
	return zap.Config{
		Level:       zap.NewAtomicLevelAt(level),
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "ts",
			LevelKey:       "level",
			NameKey:        "logger",
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.EpochTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
		},
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}
