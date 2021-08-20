package logger

import (
	"os"

	"github.com/mohammadne/bookman/user/pkg/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// zapLogger
type zapLogger struct {
	// passed dependencies
	config *Config

	// internal instance
	instance *zap.Logger
}

func NewZap(cfg *Config) logger.Logger {
	zLog := &zapLogger{config: cfg}

	zLog.instance = zap.New(
		zapcore.NewCore(
			zLog.getEncoder(),
			zLog.getWriteSyncer(),
			zLog.getLoggerLevel(),
		),
		zLog.getOptions()...,
	)

	return zLog
}

// For mapping config logger to app logger levels
var loggerLevelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func (l *zapLogger) getEncoder() zapcore.Encoder {
	var encoderConfig zapcore.EncoderConfig
	if l.config.Development {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	} else {
		encoderConfig = zap.NewProductionEncoderConfig()
	}

	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if l.config.Encoding == "console" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	return encoder
}

func (l *zapLogger) getWriteSyncer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

func (l *zapLogger) getLoggerLevel() zap.AtomicLevel {
	level, exist := loggerLevelMap[l.config.Level]
	if !exist {
		return zap.NewAtomicLevelAt(zapcore.DebugLevel)
	}

	return zap.NewAtomicLevelAt(level)
}

func (l *zapLogger) getOptions() []zap.Option {
	options := make([]zap.Option, 0, 2)

	if l.config.EnableCaller {
		options = append(options, zap.AddCaller())
	}

	if l.config.EnableStacktrace {
		options = append(options, zap.AddStacktrace(zap.ErrorLevel))
	}

	return options
}

func (l *zapLogger) Debug(msg string, fields ...logger.Field) {
	l.instance.Debug(msg, convertField(fields...)...)
}

func (l *zapLogger) Info(msg string, fields ...logger.Field) {
	l.instance.Info(msg, convertField(fields...)...)
}

func (l *zapLogger) Warn(msg string, fields ...logger.Field) {
	l.instance.Warn(msg, convertField(fields...)...)
}

func (l *zapLogger) Error(msg string, errors ...error) {
	l.instance.Panic(msg, convertError(errors...)...)
}

func (l *zapLogger) Panic(msg string, fields ...logger.Field) {
	l.instance.Panic(msg, convertField(fields...)...)
}

func (l *zapLogger) Fatal(msg string, fields ...logger.Field) {
	l.instance.Fatal(msg, convertField(fields...)...)
}

// convertField converts Field To ZapField
func convertField(fields ...logger.Field) []zapcore.Field {
	zapFileds := make([]zapcore.Field, len(fields), 0)

	for index := 0; index < len(fields); index++ {
		zapField := zap.Any(fields[index].Key, fields[index].Value)
		zapFileds = append(zapFileds, zapField)
	}

	return zapFileds
}

// convertError converts Error To ZapField
func convertError(errors ...error) []zapcore.Field {
	zapErrors := make([]zapcore.Field, len(errors), 0)

	for index := 0; index < len(errors); index++ {
		zapErrors = append(zapErrors, zap.Error(errors[index]))
	}

	return zapErrors
}
