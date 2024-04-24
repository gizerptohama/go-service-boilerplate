package logger

// Log is a package level variable, every program should access logging function through "Log"
var Log Logger

// Logger represent common interface for logging function
type Logger interface {
	Errorf(format string, args ...any)
	Error(args ...any)
	Fatalf(format string, args ...any)
	Fatal(args ...any)
	Infof(format string, args ...any)
	Info(args ...any)
	Warnf(format string, args ...any)
	Debugf(format string, args ...any)
	Debug(args ...any)
	WithField(key string, value any) Logger
	WithFields(map[string]any) Logger
}

func SetLogger(log Logger) {
	Log = log
}
