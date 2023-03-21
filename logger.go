// Modification of the standard Lg package.
package lg

import (
	"log"
	"os"
	"runtime"
)

var logger *logEnd

func init() {
	logger = newLogEnd()
}

func GetLogger() *logEnd {
	return logger
}

type prefix struct {
	Info  string
	Debug string
	Trace string
	Warn  string
	Error string
	Fatal string
	Panic string
}

func newPrefix() *prefix {
	return &prefix{
		Info:  "INFO",
		Debug: "DEBUG",
		Trace: "TRACE",
		Warn:  "WARN",
		Error: "ERROR",
		Fatal: "FATAL",
		Panic: "PANIC",
	}
}

type color struct {
	Reset,
	Red,
	Yellow,
	Purple,
	PurpleUnderline,
	Cyan,
	Gray,
	GrayUnderline string
}

func newColor() *color {
	if runtime.GOOS != "windows" {
		return &color{
			Reset:           "\033[0;0;0m",
			Red:             "\033[0;0;31m",
			Yellow:          "\033[0;0;33m",
			Purple:          "\033[0;0;35m",
			PurpleUnderline: "\033[0;4;35m",
			Cyan:            "\033[0;0;36m",
			Gray:            "\033[0;0;37m",
			GrayUnderline:   "\033[0;4;37m",
		}
	}
	return &color{}
}

type logEnd struct {
	infoLogger  *lg
	debugLogger *lg
	traceLogger *lg
	warnLogger  *lg
	errorLogger *lg
	fatalLogger *lg
	panicLogger *lg
}

func newLogEnd() *logEnd {
	p := newPrefix()
	c := newColor()
	return &logEnd{
		infoLogger:  loggerCreateNew(os.Stdout, c.Cyan, p.Info),
		debugLogger: loggerCreateNew(os.Stdout, c.GrayUnderline, p.Debug),
		traceLogger: loggerCreateNew(os.Stdout, c.Gray, p.Trace),
		warnLogger:  loggerCreateNew(os.Stdout, c.Yellow, p.Warn),
		errorLogger: loggerCreateNew(os.Stderr, c.Red, p.Error),
		fatalLogger: loggerCreateNew(os.Stderr, c.Purple, p.Fatal),
		panicLogger: loggerCreateNew(os.Stderr, c.PurpleUnderline, p.Panic),
	}
}

func (l *logEnd) OffColor() {
	offColor = true
	l.rewrite()
}

func (l *logEnd) OffPrefix() {
	offPrefix = true
	l.rewrite()
}

func (l *logEnd) OffDate() {
	logFlags = logFlags &^ log.Ldate
	l.infoLogger.log.SetFlags(logFlags)
	l.debugLogger.log.SetFlags(logFlags)
	l.traceLogger.log.SetFlags(logFlags)
	l.warnLogger.log.SetFlags(logFlags)
	l.errorLogger.log.SetFlags(logFlags)
	l.fatalLogger.log.SetFlags(logFlags)
	l.panicLogger.log.SetFlags(logFlags)
}

func (l *logEnd) OffTime() {
	logFlags = logFlags &^ log.Ltime
	l.infoLogger.log.SetFlags(logFlags)
	l.debugLogger.log.SetFlags(logFlags)
	l.traceLogger.log.SetFlags(logFlags)
	l.warnLogger.log.SetFlags(logFlags)
	l.errorLogger.log.SetFlags(logFlags)
	l.fatalLogger.log.SetFlags(logFlags)
	l.panicLogger.log.SetFlags(logFlags)
}

func (l *logEnd) OffShortFile() {
	logFlags = logFlags &^ log.Lshortfile
	l.infoLogger.log.SetFlags(logFlags)
	l.debugLogger.log.SetFlags(logFlags)
	l.traceLogger.log.SetFlags(logFlags)
	l.warnLogger.log.SetFlags(logFlags)
	l.errorLogger.log.SetFlags(logFlags)
	l.fatalLogger.log.SetFlags(logFlags)
	l.panicLogger.log.SetFlags(logFlags)
}

func (l *logEnd) rewrite() {
	log := newLogEnd()
	*l = *log
}

// info functions
func Info(v ...interface{}) error {
	return logger.infoLogger.print(v...)
}
func Infoln(v ...interface{}) error {
	return logger.infoLogger.println(v...)
}
func Infof(format string, v ...interface{}) error {
	return logger.infoLogger.printf(format, v...)
}

// debug functions
func Debug(v ...interface{}) error {
	return logger.debugLogger.print(v...)
}
func Debugln(v ...interface{}) error {
	return logger.debugLogger.println(v...)
}
func Debugf(format string, v ...interface{}) error {
	return logger.debugLogger.printf(format, v...)
}

// trace functions
func Trace(v ...interface{}) error {
	return logger.traceLogger.print(v...)
}
func Traceln(v ...interface{}) error {
	return logger.traceLogger.println(v...)
}
func Tracef(format string, v ...interface{}) error {
	return logger.traceLogger.printf(format, v...)
}

// warning functions
func Warn(v ...interface{}) error {
	return logger.warnLogger.print(v...)
}
func Warnln(v ...interface{}) error {
	return logger.warnLogger.println(v...)
}
func Warnf(format string, v ...interface{}) error {
	return logger.warnLogger.printf(format, v...)
}

// error functions
func Error(v ...interface{}) error {
	return logger.errorLogger.print(v...)
}
func Errorln(v ...interface{}) error {
	return logger.errorLogger.println(v...)
}
func Errorf(format string, v ...interface{}) error {
	return logger.errorLogger.printf(format, v...)
}

// fatal functions
func Fatal(v ...interface{}) error {
	return logger.fatalLogger.fatal(v...)
}
func Fatalln(v ...interface{}) error {
	return logger.fatalLogger.fatalln(v...)
}
func Fatalf(format string, v ...interface{}) error {
	return logger.fatalLogger.fatalf(format, v...)
}

// panic functions
func Panic(v ...interface{}) error {
	return logger.panicLogger.panic(v...)
}
func Panicln(v ...interface{}) error {
	return logger.panicLogger.panicln(v...)
}
func Panicf(format string, v ...interface{}) error {
	return logger.panicLogger.panicf(format, v...)
}
