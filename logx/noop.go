package logx

// Noop per
type Noop struct{}

// Trace impletements logx.Logger interface
func (n Noop) Trace(args ...interface{}) {
	panic("noop log")
}

// Tracef impletements logx.Logger interface
func (n Noop) Tracef(format string, args ...interface{}) {
	panic("noop log")
}

// Debug impletements logx.Logger interface
func (n Noop) Debug(args ...interface{}) {
	panic("noop log")
}

// Debugf impletements logx.Logger interface
func (n Noop) Debugf(format string, args ...interface{}) {
	panic("noop log")
}

// Info impletements logx.Logger interface
func (n Noop) Info(args ...interface{}) {
	panic("noop log")
}

// Infof impletements logx.Logger interface
func (n Noop) Infof(format string, args ...interface{}) {
	panic("noop log")
}

// Warn impletements logx.Logger interface
func (n Noop) Warn(args ...interface{}) {
	panic("noop log")
}

// Warnf impletements logx.Logger interface
func (n Noop) Warnf(format string, args ...interface{}) {
	panic("noop log")
}

// Error impletements logx.Logger interface
func (n Noop) Error(args ...interface{}) {
	panic("noop log")
}

// Errorf impletements logx.Logger interface
func (n Noop) Errorf(format string, args ...interface{}) {
	panic("noop log")
}

// Fatal impletements logx.Logger interface
func (n Noop) Fatal(args ...interface{}) {
	panic("noop log")
}

// Fatalf impletements logx.Logger interface
func (n Noop) Fatalf(format string, args ...interface{}) {
	panic("noop log")
}
