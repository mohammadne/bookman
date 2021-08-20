package logger

import (
	"github.com/mohammadne/bookman/user/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// convertFields converts Field To ZapField
func convertFields(fields ...logger.Field) []zapcore.Field {
	zapFileds := make([]zapcore.Field, len(fields), 0)

	for index := 0; index < len(fields); index++ {
		zapField := convertField(fields[index])
		zapFileds = append(zapFileds, zapField)
	}

	return zapFileds
}

func convertField(field logger.Field) zapcore.Field {
	switch field.Type {
	case logger.UnknownType:
		return zap.Any(field.Key, field.Value)
	case logger.BoolType:
		return zap.Bool(field.Key, field.Value.(bool))
	case logger.IntType:
		return zap.Int(field.Key, field.Value.(int))
	case logger.Float64Type:
		return zap.Float64(field.Key, field.Value.(float64))
	case logger.StringType:
		return zap.String(field.Key, field.Value.(string))
	case logger.ErrorType:
		return zap.Error(field.Value.(error))
	}

	return zapcore.Field{}
}

// Unknown constructs a field with the given key and value.
func Unknown(key string, val interface{}) logger.Field {
	return logger.Field{Key: key, Value: val, Type: logger.UnknownType}
}

// Int constructs a field with the given key and value.
func Int(key string, val int) logger.Field {
	return logger.Field{Key: key, Value: val, Type: logger.IntType}
}

// Float constructs a field with the given key and value.
func Float64(key string, val float64) logger.Field {
	return logger.Field{Key: key, Value: val, Type: logger.Float64Type}
}

// String constructs a field with the given key and value.
func String(key string, val string) logger.Field {
	return logger.Field{Key: key, Value: val, Type: logger.StringType}
}

// Error constructs a field with the given key and value.
func Error(val error) logger.Field {
	return logger.Field{Key: "error", Value: val, Type: logger.ErrorType}
}
