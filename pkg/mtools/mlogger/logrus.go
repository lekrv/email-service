package mlogger

import "github.com/sirupsen/logrus"

type Logrus struct {
	logger *logrus.Entry
}

func (l *Logrus) WithMapFields(m map[string]interface{}) Logger {
	//TODO implement me
	panic("implement me")
}

func (l *Logrus) Emergency(args ...interface{}) {
	l.logger.Panic(args...)
}

func (l *Logrus) Emergencyf(message string, args ...interface{}) {
	l.logger.Panicf(message, args...)
}

func (l *Logrus) Alert(args ...interface{}) {
	l.logger.Warning(args...)
}

func (l *Logrus) Alertf(message string, args ...interface{}) {
	l.logger.Warningf(message, args...)
}

func (l *Logrus) Critical(args ...interface{}) {
	l.logger.Fatal(args...)
}

func (l *Logrus) Criticalf(message string, args ...interface{}) {
	l.logger.Fatalf(message, args...)
}

func (l *Logrus) Error(args ...interface{}) {
	l.logger.Error(args...)
}

func (l *Logrus) Errorf(message string, args ...interface{}) {
	l.logger.Errorf(message, args...)
}

func (l *Logrus) Warning(args ...interface{}) {
	l.logger.Warn(args...)
}

func (l *Logrus) Warningf(message string, args ...interface{}) {
	l.logger.Warnf(message, args...)
}

func (l *Logrus) Notice(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logrus) Noticef(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *Logrus) Info(args ...interface{}) {
	l.logger.Info(args...)
}

func (l *Logrus) Infof(message string, args ...interface{}) {
	l.logger.Infof(message, args...)
}

func (l *Logrus) Debug(args ...interface{}) {
	l.logger.Debug(args...)
}

func (l *Logrus) Debugf(message string, args ...interface{}) {
	l.logger.Debugf(message, args...)
}

func (l *Logrus) WithFields(args ...interface{}) Logger {
	f := logrus.Fields{}
	for i := 0; i < len(args); i = i + 2 {
		f[args[i].(string)] = args[i+1]
	}

	l.logger = l.logger.WithFields(f)

	return l
}

func (l *Logrus) WithError(err error) Logger {
	return &Logrus{
		logger: l.logger.WithError(err),
	}
}

func NewLogrus(logger *logrus.Entry) *Logrus {
	return &Logrus{
		logger: logger,
	}
}

//-----------------------------------------------

type LogrusProvider struct {
	logger *logrus.Logger
}

func (l *LogrusProvider) Logger() Logger {
	return NewLogrus(logrus.NewEntry(l.logger))
}

func NewLogrusProvider() LoggerProvider {
	logger := logrus.New()

	logger.Formatter = &logrus.JSONFormatter{}

	return &LogrusProvider{
		logger: logger,
	}
}
