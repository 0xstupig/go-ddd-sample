package log

import "io"

type Logger interface {
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})

	Errorw(msg string, args ...interface{})
	Warnw(msg string, args ...interface{})
	Infow(msg string, args ...interface{})
	Debugw(msg string, args ...interface{})

	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(format string, args ...interface{})
}

var writer *io.Writer

func Writer() *io.Writer {
	return writer
}