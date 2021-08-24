package logger

import (
	"github.com/mohammadne/go-pkgs/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// convertFields converts Field To ZapField
func convertFields(fields ...logger.Field) []zapcore.Field {
	zapFileds := make([]zapcore.Field, 0, len(fields))

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
