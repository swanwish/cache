package logs

import "log"

const (
	EMPTY_KEY = ""
)

type LogWriter interface {
	Log(level, key string, v ...interface{})
	Logf(level, key, format string, v ...interface{})
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

func (w DefaultLogWriter) Log(level, key string, v ...interface{}) {
	Log(level, key, v...)
}

func (w DefaultLogWriter) Logf(level, key, format string, v ...interface{}) {
	Logf(level, key, format, v...)
}

func (w DefaultLogWriter) Debug(v ...interface{}) {
	Log(LOG_DEBUG, EMPTY_KEY, v...)
}

func (w DefaultLogWriter) Info(v ...interface{}) {
	Log(LOG_INFO, EMPTY_KEY, v...)
}

func (w DefaultLogWriter) Error(v ...interface{}) {
	Log(LOG_ERROR, EMPTY_KEY, v...)
}

func (w DefaultLogWriter) Fatal(v ...interface{}) {
	Log(LOG_FATAL, EMPTY_KEY, v...)
}

func (w DefaultLogWriter) Debugf(formatString string, v ...interface{}) {
	Logf(LOG_DEBUG, EMPTY_KEY, formatString, v...)
}

func (w DefaultLogWriter) Infof(formatString string, v ...interface{}) {
	Logf(LOG_INFO, EMPTY_KEY, formatString, v...)
}

func (w DefaultLogWriter) Errorf(formatString string, v ...interface{}) {
	Logf(LOG_ERROR, EMPTY_KEY, formatString, v...)
}

func (w DefaultLogWriter) Fatalf(formatString string, v ...interface{}) {
	Logf(LOG_FATAL, EMPTY_KEY, formatString, v...)
}

func Log(logLevel, key string, v ...interface{}) {
	logMessage := LogMessage(v...)
	if key == EMPTY_KEY {
		log.Println(logLevel, "-", logMessage)
	} else {
		log.Println(logLevel, "-", key, "-", logMessage)
	}
}

func Logf(logLevel, key, formatString string, v ...interface{}) {
	logMessage := LogMessagef(formatString, v...)
	if key == EMPTY_KEY {
		log.Println(logLevel, "-", logMessage)
	} else {
		log.Println(logLevel, " - ", key, " - ", logMessage)
	}
}
