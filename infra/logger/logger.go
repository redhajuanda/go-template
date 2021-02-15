package logger

import (
	"io"

	"github.com/sirupsen/logrus"
)

// Params type, used to pass to `WithParams`.
type Params map[string]interface{}

// LoggerWrapper represent common interface for logging function
type LoggerWrapper interface {
	WithParams(fields Params) LoggerWrapper
	WithParam(key string, value interface{}) LoggerWrapper
	SetFormat(format logrus.Formatter)
	SetOutputs(outputs ...io.Writer)
	Errorf(format string, args ...interface{})
	Error(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatal(args ...interface{})
	Infof(format string, args ...interface{})
	Info(args ...interface{})
	Warnf(format string, args ...interface{})
	Warn(args ...interface{})
	Debugf(format string, args ...interface{})
	Debug(args ...interface{})
}

type loggerWrapper struct {
	*logrus.Entry
}

var logStore *loggerWrapper

// New returns a new wrapper log
func New(serviceName string, environment string) LoggerWrapper {
	logStore = &loggerWrapper{logrus.New().WithField("service", serviceName).WithField("environment", environment)}
	if environment == "production" {
		logStore.SetFormat(&logrus.JSONFormatter{})
	}

	// fmt.Println("Adding hook")
	// hook := logrusly.NewLogglyHook("71000042-f956-4c7e-987d-8694a20695a8", "https://logs-01.loggly.com/bulk/", logrus.InfoLevel, serviceName)
	// logStore.Logger.Hooks.Add(hook)
	return logStore
}

func log() *loggerWrapper {
	if logStore == nil {
		return &loggerWrapper{logrus.NewEntry(logrus.New())}
	}
	return logStore
}

// WithParam create a log entry with param
func WithParam(key string, value interface{}) LoggerWrapper {
	return &loggerWrapper{log().WithFields(logrus.Fields{key: value})}
}

// WithParams create a log entry with params
func WithParams(fields Params) LoggerWrapper {
	return &loggerWrapper{log().WithFields(logrus.Fields(fields))}
}

func (logger *loggerWrapper) WithParams(fields Params) LoggerWrapper {
	return &loggerWrapper{logger.WithFields(logrus.Fields(fields))}
}

func (logger *loggerWrapper) WithParam(key string, value interface{}) LoggerWrapper {
	return &loggerWrapper{logger.WithFields(logrus.Fields{key: value})}
}

func (logger *loggerWrapper) SetFormat(format logrus.Formatter) {
	logger.Logger.SetFormatter(format)
}

func (logger *loggerWrapper) SetOutputs(outputs ...io.Writer) {
	mw := io.MultiWriter(outputs...)
	logger.Logger.SetOutput(mw)
}

func Errorf(format string, args ...interface{}) {
	log().Errorf(format, args...)
}
func Error(args ...interface{}) {
	log().Error(args...)
}
func Fatalf(format string, args ...interface{}) {
	log().Fatalf(format, args...)
}
func Fatal(args ...interface{}) {
	log().Fatal(args...)
}
func Infof(format string, args ...interface{}) {
	log().Infof(format, args...)
}
func Info(args ...interface{}) {
	log().Info(args...)
}
func Warnf(format string, args ...interface{}) {
	log().Warnf(format, args...)
}
func Warn(args ...interface{}) {
	log().Warn(args...)
}
func Debugf(format string, args ...interface{}) {
	log().Debugf(format, args...)
}
func Debug(args ...interface{}) {
	log().Debug(args...)
}
