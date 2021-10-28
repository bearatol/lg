// Modification of the standard Lg package.
package lg

import (
	"os"
	"runtime"
)

var c = newColor()
var p = newPrefix()
var Lg = newLog()

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

type Log struct {
	infoLogger  *lg
	debugLogger *lg
	traceLogger *lg
	warnLogger  *lg
	errorLogger *lg
	fatalLogger *lg
	panicLogger *lg
}

func newLog() *Log {
	return &Log{
		infoLogger:  loggerCreateNew(os.Stdout, c.Cyan, p.Info),
		debugLogger: loggerCreateNew(os.Stdout, c.GrayUnderline, p.Debug),
		traceLogger: loggerCreateNew(os.Stdout, c.Gray, p.Trace),
		warnLogger:  loggerCreateNew(os.Stdout, c.Yellow, p.Warn),
		errorLogger: loggerCreateNew(os.Stderr, c.Red, p.Error),
		fatalLogger: loggerCreateNew(os.Stderr, c.Purple, p.Fatal),
		panicLogger: loggerCreateNew(os.Stderr, c.PurpleUnderline, p.Panic),
	}
}

func (l *Log) OffColor() {
	l.infoLogger = loggerCreateNew(os.Stdout, "", p.Info)
	l.debugLogger = loggerCreateNew(os.Stdout, "", p.Debug)
	l.traceLogger = loggerCreateNew(os.Stdout, "", p.Trace)
	l.warnLogger = loggerCreateNew(os.Stdout, "", p.Warn)
	l.errorLogger = loggerCreateNew(os.Stderr, "", p.Error)
	l.fatalLogger = loggerCreateNew(os.Stderr, "", p.Fatal)
	l.panicLogger = loggerCreateNew(os.Stderr, "", p.Panic)
}

//info functions
func Info(v ...interface{}) {
	Lg.infoLogger.print(v...)
}
func Infoln(v ...interface{}) {
	Lg.infoLogger.println(v...)
}
func Infof(format string, v ...interface{}) {
	Lg.infoLogger.printf(format, v...)
}

//debug functions
func Debug(v ...interface{}) {
	Lg.debugLogger.print(v...)
}
func Debugln(v ...interface{}) {
	Lg.debugLogger.println(v...)
}
func Debugf(format string, v ...interface{}) {
	Lg.debugLogger.printf(format, v...)
}

//trace functions
func Trace(v ...interface{}) {
	Lg.traceLogger.print(v...)
}
func Traceln(v ...interface{}) {
	Lg.traceLogger.println(v...)
}
func Tracef(format string, v ...interface{}) {
	Lg.traceLogger.printf(format, v...)
}

//warning functions
func Warn(v ...interface{}) {
	Lg.warnLogger.print(v...)
}
func Warnln(v ...interface{}) {
	Lg.warnLogger.println(v...)
}
func Warnf(format string, v ...interface{}) {
	Lg.warnLogger.printf(format, v...)
}

//error functions
func Error(v ...interface{}) {
	Lg.errorLogger.print(v...)
}
func Errorln(v ...interface{}) {
	Lg.errorLogger.println(v...)
}
func Errorf(format string, v ...interface{}) {
	Lg.errorLogger.printf(format, v...)
}

//fatal functions
func Fatal(v ...interface{}) {
	Lg.fatalLogger.fatal(v...)
}
func Fatalln(v ...interface{}) {
	Lg.fatalLogger.fatalln(v...)
}
func Fatalf(format string, v ...interface{}) {
	Lg.fatalLogger.fatalf(format, v...)
}

//panic functions
func Panic(v ...interface{}) {
	Lg.panicLogger.panic(v...)
}
func Panicln(v ...interface{}) {
	Lg.panicLogger.panicln(v...)
}
func Panicf(format string, v ...interface{}) {
	Lg.panicLogger.panicf(format, v...)
}
