// Modification of the standard log package.
package lg

import (
	"os"

	"github.com/bearatol/lg/color"
)

var (
	infoLogger  *lg = loggerCreateNew(os.Stdout, color.Color.Cyan, "INFO")
	debugLogger *lg = loggerCreateNew(os.Stdout, color.Color.GrayUnderline, "DEBUG")
	traceLogger *lg = loggerCreateNew(os.Stdout, color.Color.Gray, "TRACE")
	warnLogger  *lg = loggerCreateNew(os.Stdout, color.Color.Yellow, "WARN")
	errorLogger *lg = loggerCreateNew(os.Stderr, color.Color.Red, "ERROR")
	fatalLogger *lg = loggerCreateNew(os.Stderr, color.Color.Purple, "FATAL")
	panicLogger *lg = loggerCreateNew(os.Stderr, color.Color.PurpleUnderline, "PANIC")
)

//info functions
func Info(v ...interface{}) {
	infoLogger.print(v...)
}
func Infoln(v ...interface{}) {
	infoLogger.println(v...)
}
func Infof(format string, v ...interface{}) {
	infoLogger.printf(format, v...)
}

//debug functions
func Debug(v ...interface{}) {
	debugLogger.print(v...)
}
func Debugln(v ...interface{}) {
	debugLogger.println(v...)
}
func Debugf(format string, v ...interface{}) {
	debugLogger.printf(format, v...)
}

//trace functions
func Trace(v ...interface{}) {
	traceLogger.print(v...)
}
func Traceln(v ...interface{}) {
	traceLogger.println(v...)
}
func Tracef(format string, v ...interface{}) {
	traceLogger.printf(format, v...)
}

//warning functions
func Warn(v ...interface{}) {
	warnLogger.print(v...)
}
func Warnln(v ...interface{}) {
	warnLogger.println(v...)
}
func Warnf(format string, v ...interface{}) {
	warnLogger.printf(format, v...)
}

//error functions
func Err(v ...interface{}) {
	errorLogger.print(v...)
}
func Errln(v ...interface{}) {
	errorLogger.println(v...)
}
func Errf(format string, v ...interface{}) {
	errorLogger.printf(format, v...)
}

//fatal functions
func Fatal(v ...interface{}) {
	fatalLogger.fatal(v...)
}
func Fatalln(v ...interface{}) {
	fatalLogger.fatalln(v...)
}
func Fatalf(format string, v ...interface{}) {
	fatalLogger.fatalf(format, v...)
}

//panic functions
func Panic(v ...interface{}) {
	panicLogger.panic(v...)
}
func Panicln(v ...interface{}) {
	panicLogger.panicln(v...)
}
func Panicf(format string, v ...interface{}) {
	panicLogger.panicf(format, v...)
}
