package logging

import (
	"fmt"
	"log"
)

type StdOutLogger struct {
	AbstractLogger
}

func (l *StdOutLogger) Log(v ...interface{}) {
	log.Print(v...)
}

func (l *StdOutLogger) Logf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *StdOutLogger) Logs(service string, v ...interface{}) {
	log.Print(FormatLogMessage(service, fmt.Sprint(v...)))
}

func (l *StdOutLogger) Logsf(service string, format string, v ...interface{}) {
	l.Logs(service, fmt.Sprintf(format, v...))
}

var StdOutLog *StdOutLogger = &StdOutLogger{}
