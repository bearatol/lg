// Сreated a new structure based on the logger structure of the standard log package.
package lg

import (
	"fmt"
	"log"
	"os"
)

type lg struct {
	*log.Logger
}

func loggerCreateNew(std *os.File, color string, prefix string) *lg {
	return &lg{
		log.New(std, color+prefix+"\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *lg) print(v ...interface{})   { l.Output(3, fmt.Sprint(v...)) }
func (l *lg) println(v ...interface{}) { l.Output(3, fmt.Sprintln(v...)) }
func (l *lg) printf(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
}

func (l *lg) fatal(v ...interface{}) {
	l.Output(3, fmt.Sprint(v...))
	os.Exit(1)
}
func (l *lg) fatalln(v ...interface{}) {
	l.Output(3, fmt.Sprintln(v...))
	os.Exit(1)
}
func (l *lg) fatalf(format string, v ...interface{}) {
	l.Output(3, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *lg) panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.Output(3, s)
	panic(s)
}
func (l *lg) panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	l.Output(3, s)
	panic(s)
}
func (l *lg) panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.Output(3, s)
	panic(s)
}
