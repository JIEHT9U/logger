package logger

import (
	"context"

	"github.com/opentracing/opentracing-go"
	tag "github.com/opentracing/opentracing-go/ext"
	"go.uber.org/zap"

	"github.com/jieht9u/logger/fieldadapter"
)

type Opentracing struct {
	logger *zap.Logger
	span   opentracing.Span
	fields []zap.Field
}

func For(ctx context.Context) *Opentracing {
	return &Opentracing{span: opentracing.SpanFromContext(ctx), logger: globalLogger.zapLogger}
}

func (o *Opentracing) Info(msg string, args ...zap.Field) {
	o.createDefaultMsg("info", msg, args...).fillSpan().info(msg, args...)
}

func (o *Opentracing) Error(msg string, args ...zap.Field) {
	o.createDefaultMsg("error", msg, args...).fillSpan().setErrorTag().error(msg, args...)
}

func (o *Opentracing) Fatal(msg string, args ...zap.Field) {
	o.createDefaultMsg("fatal", msg, args...).fillSpan().fatal(msg, args...)
}

func (o *Opentracing) With(fields ...zap.Field) *Opentracing {
	return &Opentracing{
		span:   o.span,
		logger: o.logger.With(fields...),
		fields: append(o.fields, fields...),
	}
}

func (o *Opentracing) fillSpan() *Opentracing {
	return o.execIfSpanExist(func() {
		var fa = fieldadapter.New(0, len(o.fields))
		for _, field := range o.fields {
			field.AddTo(&fa)
		}
		o.span.LogFields(fa...)
	})
}

func (o *Opentracing) createDefaultMsg(level, msg string, fields ...zap.Field) *Opentracing {
	return o.addField(zap.String("level", level), zap.String("event", msg)).addField(fields...)
}

func (o *Opentracing) addField(fields ...zap.Field) *Opentracing {
	o.fields = append(o.fields, fields...)
	return o
}

func (o *Opentracing) setErrorTag() *Opentracing {
	tag.Error.Set(o.span, true)
	return o
}

func (o *Opentracing) info(msg string, args ...zap.Field) *Opentracing {
	o.logger.Info(msg, args...)
	return o
}

func (o *Opentracing) error(msg string, args ...zap.Field) *Opentracing {
	o.logger.Error(msg, args...)
	return o
}

func (o *Opentracing) fatal(msg string, args ...zap.Field) *Opentracing {
	o.logger.Fatal(msg, args...)
	return o
}

func (o *Opentracing) execIfSpanExist(f func()) *Opentracing {
	if o.span != nil {
		f()
	}
	return o
}
