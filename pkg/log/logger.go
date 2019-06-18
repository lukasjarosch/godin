package log

import (
	"fmt"
	"os"
	"strings"

	kitlog "github.com/go-kit/kit/log"
)

type severity int

func (s severity) String() string {
	return logLevelName[s]
}

// GCP log-severity levels, see https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry#LogSeverity
const (
	DEFAULT severity = iota
	DEBUG
	INFO
	NOTICE
	WARNING
	ERROR
	CRITICAL
	ALERT
	EMERGENCY
)

var (
	logLevel severity
)

var logLevelName = [...]string{
	"DEFAULT",
	"DEBUG",
	"INFO",
	"NOTICE",
	"WARNING",
	"ERROR",
	"CRITICAL",
	"ALERT",
	"EMERGENCY",
}

var logLevelValue = map[string]severity{
	"DEFAULT":   DEFAULT,
	"DEBUG":     DEBUG,
	"INFO":      INFO,
	"NOTICE":    NOTICE,
	"WARNING":   WARNING,
	"ERROR":     ERROR,
	"CRITICAL":  CRITICAL,
	"ALERT":     ALERT,
	"EMERGENCY": EMERGENCY,
}

func init() {
	ll, ok := logLevelValue[strings.ToUpper(os.Getenv("LOG_LEVEL"))]
	if !ok {
		fmt.Println("logger WARN: LOG_LEVEL is not set, defaulting to INFO")
		logLevel = logLevelValue[INFO.String()]
	} else {
		logLevel = ll
	}
}

type Logger struct {
	kitlogger kitlog.Logger
}

func New() *Logger {
	logger := kitlog.NewJSONLogger(kitlog.NewSyncWriter(os.Stdout))
	logger = kitlog.With(logger, "timestamp", kitlog.DefaultTimestamp)
	logger = kitlog.With(logger, "caller", kitlog.DefaultCaller)

	return &Logger{
		kitlogger: logger,
	}
}

func (l *Logger) Debug(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(DEBUG, message, keyvals))
}

func (l *Logger) Info(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(INFO, message, keyvals))
}

func (l *Logger) Notice(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(NOTICE, message, keyvals))
}

func (l *Logger) Warning(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(WARNING, message, keyvals))
}

func (l *Logger) Error(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(ERROR, message, keyvals))
}

func (l *Logger) Critical(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(CRITICAL, message, keyvals))
}

func (l *Logger) Alert(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(ALERT, message, keyvals))
}

func (l *Logger) Emergency(message string, keyvals ...interface{}) error {
	return l.kitlogger.Log(l.mergeKeyValues(EMERGENCY, message, keyvals))
}

// mergeKeyValues will append the level and message field to already existing keyvals
func (l *Logger) mergeKeyValues(level severity, message string, keyvals ...interface{}) []interface{} {
	levelData := []interface{}{
		"level",
		strings.ToLower(logLevelName[level]),
		"message",
		message,
	}
	keyvals = append(keyvals, levelData...)

	return keyvals
}
