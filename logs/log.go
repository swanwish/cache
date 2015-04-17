package logs

import (
	"fmt"
	"log"
)

const (
	LOG_DEBUG = "debug"
	LOG_INFO  = "info"
	LOG_ERROR = "error"
	LOG_FATAL = "fatal"
)

func Debug(v ...interface{}) {
	Log(LOG_DEBUG, v...)
}

func Info(v ...interface{}) {
	Log(LOG_INFO, v...)
}

func Error(v ...interface{}) {
	Log(LOG_ERROR, v...)
}

func Fatal(v ...interface{}) {
	Log(LOG_FATAL, v...)
}

func Debugf(formatString string, v ...interface{}) {
	Logf(LOG_DEBUG, formatString, v...)
}

func Infof(formatString string, v ...interface{}) {
	Logf(LOG_INFO, formatString, v...)
}

func Errorf(formatString string, v ...interface{}) {
	Logf(LOG_ERROR, formatString, v...)
}

func Log(logLevel string, v ...interface{}) {
	logMessage := fmt.Sprint(v...)
	log.Println(logLevel, " - ", logMessage)
}

func Logf(logLevel, formatString string, v ...interface{}) {
	logMessage := fmt.Sprintf(formatString, v...)
	log.Println(logLevel, "-", logMessage)
}
