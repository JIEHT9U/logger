package fieldadapter

import (
	"time"

	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap/zapcore"
)

type FieldAdapter []log.Field

func New(len, cup int) FieldAdapter {
	return make(FieldAdapter, len, cup)
}

func (fa *FieldAdapter) AddBool(key string, value bool) {
	*fa = append(*fa, log.Bool(key, value))
}

func (fa *FieldAdapter) AddFloat64(key string, value float64) {
	*fa = append(*fa, log.Float64(key, value))
}

func (fa *FieldAdapter) AddFloat32(key string, value float32) {
	*fa = append(*fa, log.Float64(key, float64(value)))
}

func (fa *FieldAdapter) AddInt(key string, value int) {
	*fa = append(*fa, log.Int(key, value))
}

func (fa *FieldAdapter) AddInt64(key string, value int64) {
	*fa = append(*fa, log.Int64(key, value))
}

func (fa *FieldAdapter) AddInt32(key string, value int32) {
	*fa = append(*fa, log.Int64(key, int64(value)))
}

func (fa *FieldAdapter) AddInt16(key string, value int16) {
	*fa = append(*fa, log.Int64(key, int64(value)))
}

func (fa *FieldAdapter) AddInt8(key string, value int8) {
	*fa = append(*fa, log.Int64(key, int64(value)))
}

func (fa *FieldAdapter) AddUint(key string, value uint) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

func (fa *FieldAdapter) AddUint64(key string, value uint64) {
	*fa = append(*fa, log.Uint64(key, value))
}

func (fa *FieldAdapter) AddUint32(key string, value uint32) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

func (fa *FieldAdapter) AddUint16(key string, value uint16) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

func (fa *FieldAdapter) AddUint8(key string, value uint8) {
	*fa = append(*fa, log.Uint64(key, uint64(value)))
}

func (fa *FieldAdapter) AddUintptr(key string, value uintptr)                        {}
func (fa *FieldAdapter) AddArray(key string, marshaler zapcore.ArrayMarshaler) error { return nil }
func (fa *FieldAdapter) AddComplex128(key string, value complex128)                  {}
func (fa *FieldAdapter) AddComplex64(key string, value complex64)                    {}
func (fa *FieldAdapter) AddObject(key string, value zapcore.ObjectMarshaler) error   { return nil }
func (fa *FieldAdapter) AddReflected(key string, value interface{}) error            { return nil }
func (fa *FieldAdapter) OpenNamespace(key string)                                    {}

func (fa *FieldAdapter) AddDuration(key string, value time.Duration) {
	// TODO inefficient
	*fa = append(*fa, log.String(key, value.String()))
}

func (fa *FieldAdapter) AddTime(key string, value time.Time) {
	// TODO inefficient
	*fa = append(*fa, log.String(key, value.String()))
}

func (fa *FieldAdapter) AddBinary(key string, value []byte) {
	*fa = append(*fa, log.Object(key, value))
}

func (fa *FieldAdapter) AddByteString(key string, value []byte) {
	*fa = append(*fa, log.Object(key, value))
}

func (fa *FieldAdapter) AddString(key, value string) {
	if key != "" && value != "" {
		*fa = append(*fa, log.String(key, value))
	}
}
