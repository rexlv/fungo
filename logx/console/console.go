package console

import (
	"fmt"
	"os"
	"strings"

	"github.com/rexlv/fungo/logx"
)

var console = &Console{
	lv:     logx.TRAC | logx.DEBU | logx.INFO | logx.WARN | logx.ERRO | logx.FATA,
	prefix: "FUN]>",
}

func init() {
	logx.Register("console", console)
}

// Console logger
type Console struct {
	lv     int
	prefix string
}

// Trace impletements logx.Logger interface
func (c Console) Trace(args ...interface{}) {
	c.print(logx.TRAC, args...)
}

// Tracef impletements logx.Logger interface
func (c Console) Tracef(format string, args ...interface{}) {
	c.printf(logx.TRAC, format, args...)
}

// Debug impletements logx.Logger interface
func (c Console) Debug(args ...interface{}) {
	c.print(logx.DEBU, args...)
}

// Debugf impletements logx.Logger interface
func (c Console) Debugf(format string, args ...interface{}) {
	c.printf(logx.DEBU, format, args...)
}

// Info impletements logx.Logger interface
func (c Console) Info(args ...interface{}) {
	c.print(logx.INFO, args...)
}

// Infof impletements logx.Logger interface
func (c Console) Infof(format string, args ...interface{}) {
	c.printf(logx.INFO, format, args...)
}

// Warn impletements logx.Logger interface
func (c Console) Warn(args ...interface{}) {
	c.print(logx.WARN, args...)
}

// Warnf impletements logx.Logger interface
func (c Console) Warnf(format string, args ...interface{}) {
	c.printf(logx.WARN, format, args...)
}

// Error impletements logx.Logger interface
func (c Console) Error(args ...interface{}) {
	c.print(logx.ERRO, args...)
}

// Errorf impletements logx.Logger interface
func (c Console) Errorf(format string, args ...interface{}) {
	c.printf(logx.ERRO, format, args...)
}

// Fatal impletements logx.Logger interface
func (c Console) Fatal(args ...interface{}) {
	c.print(logx.FATA, args...)
}

// Fatalf impletements logx.Logger interface
func (c Console) Fatalf(format string, args ...interface{}) {
	c.printf(logx.FATA, format, args...)
}

func (c Console) print(lv int, args ...interface{}) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf(c.prefix+"(%d)"+color(lv)+"%s", os.Getpid(), label(lv), fmt.Sprintf(strings.Repeat(" %+v", len(args)), args...)))
}

func (c Console) printf(lv int, format string, args ...interface{}) {
	fmt.Fprintln(os.Stdout, fmt.Sprintf(c.prefix+"(%d)"+color(lv)+"%s", os.Getpid(), label(lv), fmt.Sprintf(format, args...)))
}

func color(lv int) string {
	if c, ok := colorMap[lv]; ok {
		return c
	}

	return cyanFormat
}

func label(lv int) string {
	if l, ok := labelMap[lv]; ok {
		return l
	}

	return "Unknown"
}

var colorMap = map[int]string{
	logx.TRAC: magentaFormat,
	logx.DEBU: blueFormat,
	logx.INFO: greenFormat,
	logx.WARN: yellowFormat,
	logx.ERRO: redFormat,
	logx.FATA: redFormat,
}

var labelMap = map[int]string{
	logx.TRAC: "TRAC",
	logx.DEBU: "DEBU",
	logx.INFO: "INFO",
	logx.WARN: "WARN",
	logx.ERRO: "ERRO",
	logx.FATA: "FATA",
}

const (
	blackFormat   = "\x1b[30m%s\x1b[0m"
	redFormat     = "\x1b[31m%s\x1b[0m"
	greenFormat   = "\x1b[32m%s\x1b[0m"
	yellowFormat  = "\x1b[33m%s\x1b[0m"
	blueFormat    = "\x1b[34m%s\x1b[0m"
	magentaFormat = "\x1b[35m%s\x1b[0m"
	cyanFormat    = "\x1b[36m%s\x1b[0m"
	whiteFormat   = "\x1b[37m%s\x1b[0m"
)
