package logx

var (
	logger  = Noop{} // default logger
	loggers = make(map[string]Logger)
)

// Logger interface{}
type Logger interface {
	Trace(...interface{})
	Tracef(string, ...interface{})
	Debug(...interface{})
	Debugf(string, ...interface{})
	Info(...interface{})
	Infof(string, ...interface{})
	Warn(...interface{})
	Warnf(string, ...interface{})
	Error(...interface{})
	Errorf(string, ...interface{})
	Fatal(...interface{})
	Fatalf(string, ...interface{})
}

const (
	TRAC = 1 << iota
	DEBU
	INFO
	WARN
	ERRO
	FATA
)

// Trace impletements logx.Logger interface
func Trace(args ...interface{}) {
	logger.Trace(args...)
}

// Tracef impletements logx.Logger interface
func Tracef(format string, args ...interface{}) {
	logger.Tracef(format, args...)
}

// Debug impletements logx.Logger interface
func Debug(args ...interface{}) {
	logger.Debug(args...)
}

// Debugf impletements logx.Logger interface
func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

// Info impletements logx.Logger interface
func Info(args ...interface{}) {
	logger.Info(args...)
}

// Infof impletements logx.Logger interface
func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

// Warn impletements logx.Logger interface
func Warn(args ...interface{}) {
	logger.Warn(args...)
}

// Warnf impletements logx.Logger interface
func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

// Error impletements logx.Logger interface
func Error(args ...interface{}) {
	logger.Error(args...)
}

// Errorf impletements logx.Logger interface
func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

// Fatal impletements logx.Logger interface
func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

// Fatalf impletements logx.Logger interface
func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Register register
func Register(name string, logger Logger) {
	if logger == nil {
		panic("log: Register logger is nil")
	}

	if _, dup := loggers[name]; dup {
		panic("log: Register failed")
	}

	loggers[name] = logger
}

// Log returns named Logger
func Log(name string) Logger {
	if logger, exist := loggers[name]; !exist {
		panic("log: undefined logger " + name)
	} else {
		return logger
	}
}
