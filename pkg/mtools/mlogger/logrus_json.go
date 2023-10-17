package mlogger

import "github.com/sirupsen/logrus"

type LogrusJSON struct {
	logger *logrus.Entry
}

func (l LogrusJSON) WithMapFields(m map[string]interface{}) Logger {
	l.logger = l.logger.WithFields(m)
	return l
}

func (l LogrusJSON) Emergency(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l LogrusJSON) Emergencyf(message string, args ...interface{}) {
	l.logger.Panicf(message, args...)
}

func (l LogrusJSON) Alert(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l LogrusJSON) Alertf(message string, args ...interface{}) {
	l.logger.Warningf(message, args...)
}

func (l LogrusJSON) Critical(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l LogrusJSON) Criticalf(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
}

func (l LogrusJSON) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l LogrusJSON) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l LogrusJSON) Warning(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l LogrusJSON) Warningf(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l LogrusJSON) Notice(args ...interface{}) {
	l.logger.Print(args...)
}

func (l LogrusJSON) Noticef(message string, args ...interface{}) {
	l.logger.Printf(message, args...)
}

func (l LogrusJSON) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l LogrusJSON) Infof(message string, args ...interface{}) {
	l.logger.Info(args...)
}

func (l LogrusJSON) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l LogrusJSON) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l LogrusJSON) WithFields(args ...interface{}) Logger {
	f := logrus.Fields{}
	for i := 0; i < len(args); i = i + 2 {
		f[args[i].(string)] = args[i+1]
	}

	l.logger = l.logger.WithFields(f)

	return l
}

func (l LogrusJSON) WithError(err error) Logger {
	return &LogrusJSON{
		logger: l.logger.WithError(err),
	}
}

func NewLogrusJSON(logger *logrus.Entry) *LogrusJSON {
	return &LogrusJSON{
		logger: logger,
	}
}

//-----------------------------------------------

type LogrusJSONProvider struct {
	logger *logrus.Logger
}

func (l *LogrusJSONProvider) Logger() Logger {
	return NewLogrusJSON(logrus.NewEntry(l.logger))
}

func NewLogrusJsonProvider() LoggerProvider {
	logger := logrus.New()

	logger.Formatter = &logrus.JSONFormatter{}

	return &LogrusJSONProvider{
		logger: logger,
	}
}
