package mlogger

import (
	formatters "github.com/fabienm/go-logrus-formatters"
	"github.com/sirupsen/logrus"
	"os"
)

type LogrusGelf struct {
	logger *logrus.Entry
}

func (l LogrusGelf) WithMapFields(m map[string]interface{}) Logger {
	//TODO implement me
	panic("implement me")
}

func (l LogrusGelf) Emergency(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l LogrusGelf) Emergencyf(message string, args ...interface{}) {
	l.logger.Panicf(message, args...)
}

func (l LogrusGelf) Alert(args ...interface{}) {
	panic("implement me")
}

func (l LogrusGelf) Alertf(message string, args ...interface{}) {
	panic("implement me")
}

func (l LogrusGelf) Critical(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l LogrusGelf) Criticalf(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
}

func (l LogrusGelf) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l LogrusGelf) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l LogrusGelf) Warning(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l LogrusGelf) Warningf(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l LogrusGelf) Notice(args ...interface{}) {
	panic("implement me")
}

func (l LogrusGelf) Noticef(message string, args ...interface{}) {
	panic("implement me")
}

func (l LogrusGelf) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l LogrusGelf) Infof(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l LogrusGelf) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l LogrusGelf) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l LogrusGelf) WithFields(args ...interface{}) Logger {
	f := logrus.Fields{}
	for i := 0; i < len(args); i = i + 2 {
		f[args[i].(string)] = args[i+1]
	}

	l.logger = l.logger.WithFields(f)
	return l
}

func (l LogrusGelf) WithError(err error) Logger {
	return &LogrusGelf{
		logger: l.logger.WithError(err),
	}
}

func NewLogrusGelf(logger *logrus.Entry) *LogrusGelf {
	return &LogrusGelf{
		logger: logger,
	}
}

//-----------------------------------------------

type LogrusGelfProvider struct {
	logger *logrus.Logger
}

func (l *LogrusGelfProvider) Logger() Logger {
	return NewLogrusGelf(logrus.NewEntry(l.logger))
}

func NewLogrusGelfProvider() LoggerProvider {
	logger := logrus.New()

	hostname, err := os.Hostname()

	if err != nil {
		panic("No hostname")
	}

	logger.SetFormatter(formatters.NewGelf(hostname))

	return &LogrusGelfProvider{
		logger: logger,
	}
}
