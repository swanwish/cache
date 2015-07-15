package logs

import (
	"fmt"
	"net/smtp"

	"github.com/AeroNotix/libsmtp"
)

type EmailConfig struct {
	Username      string
	Password      string
	Host          string
	Port          int64
	From          string
	To            []string
	SubjectPrefix string
}

const (
	DEFAULT_SUBJECT_PREFIX = "EMAIL LOG"
)

var EMailLogger LogWriter = EMailLogWriter{}

type EMailLogWriter struct {
	EmailConfig
}

func (w EMailLogWriter) Log(level, key string, v ...interface{}) {
	message := LogMessage(v...)
	w.sendMail(w.getMailSubject(level, key), message)
}

func (w EMailLogWriter) Logf(level, key, format string, v ...interface{}) {
	message := LogMessagef(format, v...)
	w.sendMail(w.getMailSubject(level, key), message)
}

func (w EMailLogWriter) Debug(v ...interface{}) {
	w.Log(LOG_DEBUG, EMPTY_KEY, v...)
}

func (w EMailLogWriter) Info(v ...interface{}) {
	w.Log(LOG_INFO, EMPTY_KEY, v...)
}

func (w EMailLogWriter) Error(v ...interface{}) {
	w.Log(LOG_ERROR, EMPTY_KEY, v...)
}

func (w EMailLogWriter) Fatal(v ...interface{}) {
	w.Log(LOG_FATAL, EMPTY_KEY, v...)
}

func (w EMailLogWriter) Debugf(formatString string, v ...interface{}) {
	w.Logf(LOG_DEBUG, EMPTY_KEY, formatString, v...)
}

func (w EMailLogWriter) Infof(formatString string, v ...interface{}) {
	w.Logf(LOG_INFO, EMPTY_KEY, formatString, v...)
}

func (w EMailLogWriter) Errorf(formatString string, v ...interface{}) {
	w.Logf(LOG_ERROR, EMPTY_KEY, formatString, v...)
}

func (w EMailLogWriter) Fatalf(formatString string, v ...interface{}) {
	w.Logf(LOG_FATAL, EMPTY_KEY, formatString, v...)
}

func (l EMailLogWriter) sendMail(title, message string) {
	if l.Host == "" {
		Info("The email log is not enabled")
		return
	}
	auth := smtp.PlainAuth("", l.Username, l.Password, l.Host)
	err := libsmtp.SendMailWithAttachments(fmt.Sprintf("%s:%d", l.Host, l.Port), &auth, l.From, title, l.To, []byte(message), nil)
	if err != nil {
		Errorf("Failed to send mail, the error is %v", err)
	}
}

func (w EMailLogWriter) getMailSubject(level, key string) string {
	subject := ""
	if w.SubjectPrefix == "" {
		subject = DEFAULT_SUBJECT_PREFIX
	} else {
		subject = w.SubjectPrefix
	}
	subject = fmt.Sprintf("%s [%s]", subject, level)

	if key != "" {
		subject += " - " + key
	}

	return subject
}
