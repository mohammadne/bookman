package logger

// Logger methods interface
type Logger interface {
	Debug(string, ...Field)
	Info(string, ...Field)
	Warn(string, ...Field)
	Error(string, ...error)
	Panic(string, ...Field)
	Fatal(string, ...Field)
}
