package mlogger

// Logger interface to logger
type Logger interface {
	Emergency(args ...interface{})
	Emergencyf(message string, args ...interface{})
	Alert(args ...interface{})
	Alertf(message string, args ...interface{})
	Critical(args ...interface{})
	Criticalf(message string, args ...interface{})
	Error(args ...interface{})
	Errorf(message string, args ...interface{})
	Warning(args ...interface{})
	Warningf(message string, args ...interface{})
	Notice(args ...interface{})
	Noticef(message string, args ...interface{})
	Info(args ...interface{})
	Infof(message string, args ...interface{})
	Debug(args ...interface{})
	Debugf(message string, args ...interface{})
	WithFields(args ...interface{}) Logger
	WithMapFields(map[string]interface{}) Logger
	WithError(err error) Logger
}

// LoggerProvider logger provider
type LoggerProvider interface {
	Logger() Logger
}
