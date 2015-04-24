package logs

import "fmt"

const (
	LOG_DEBUG = "debug"
	LOG_INFO  = "info"
	LOG_ERROR = "error"
	LOG_FATAL = "fatal"
)

var Writer LogWriter = DefaultLogWriter{}

func Debug(v ...interface{}) {
	if Writer != nil {
		Writer.Debug(v...)
	}
}

func Info(v ...interface{}) {
	if Writer != nil {
		Writer.Info(v...)
	}
}

func Error(v ...interface{}) {
	if Writer != nil {
		Writer.Error(v...)
	}
}

func Fatal(v ...interface{}) {
	if Writer != nil {
		Writer.Fatal(v...)
	}
}

func Debugf(formatString string, v ...interface{}) {
	if Writer != nil {
		Writer.Debugf(formatString, v...)
	}
}

func Infof(formatString string, v ...interface{}) {
	if Writer != nil {
		Writer.Infof(formatString, v...)
	}
}

func Errorf(formatString string, v ...interface{}) {
	if Writer != nil {
		Writer.Errorf(formatString, v...)
	}
}

func Fatalf(formatString string, v ...interface{}) {
	if Writer != nil {
		Writer.Fatalf(formatString, v...)
	}
}

func LogMessage(v ...interface{}) string {
	return fmt.Sprint(v...)
}

func LogMessagef(formatString string, v ...interface{}) string {
	return fmt.Sprintf(formatString, v...)
}
