package logs

import (
	"fmt"
	"log"
)

type LogWriter interface {
	Debug(v ...interface{})
	Info(v ...interface{})
	Error(v ...interface{})
	Fatal(v ...interface{})
	Debugf(formatString string, v ...interface{})
	Infof(formatString string, v ...interface{})
	Errorf(formatString string, v ...interface{})
	Fatalf(formatString string, v ...interface{})
}

type DefaultLogWriter struct {
}

func (w DefaultLogWriter) Debug(v ...interface{}) {
	Log(LOG_DEBUG, v...)
}

func (w DefaultLogWriter) Info(v ...interface{}) {
	Log(LOG_INFO, v...)
}

func (w DefaultLogWriter) Error(v ...interface{}) {
	Log(LOG_ERROR, v...)
}

func (w DefaultLogWriter) Fatal(v ...interface{}) {
	Log(LOG_FATAL, v...)
}

func (w DefaultLogWriter) Debugf(formatString string, v ...interface{}) {
	Logf(LOG_DEBUG, formatString, v...)
}

func (w DefaultLogWriter) Infof(formatString string, v ...interface{}) {
	Logf(LOG_INFO, formatString, v...)
}

func (w DefaultLogWriter) Errorf(formatString string, v ...interface{}) {
	Logf(LOG_ERROR, formatString, v...)
}

func (w DefaultLogWriter) Fatalf(formatString string, v ...interface{}) {
	Logf(LOG_FATAL, formatString, v...)
}

func Log(logLevel string, v ...interface{}) {
	logMessage := fmt.Sprint(v...)
	log.Println(logLevel, " - ", logMessage)
}

func Logf(logLevel, formatString string, v ...interface{}) {
	logMessage := fmt.Sprintf(formatString, v...)
	log.Println(logLevel, "-", logMessage)
}
