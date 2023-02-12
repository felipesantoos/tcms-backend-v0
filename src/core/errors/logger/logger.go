package logger

import (
	"github.com/rs/zerolog"
	"os"
	"runtime"
	"tcms/src/core/errors"
	"time"
)

const (
	timeKey = "time"
	fileKey = "file"
	lineKey = "line"
)

type Logger interface {
	Log(err errors.Error)
}

type TCMSLogger struct {
	zeroLogger zerolog.Logger
}

func (instance TCMSLogger) Log(err errors.Error) {
	instance.getEventLog(err.LogLevel()).
		Time(timeKey, time.Now()).
		Str(fileKey, err.File()).
		Int(lineKey, err.Line()).
		Msg(err.Error())
}

func (instance TCMSLogger) LogWithMessage(logLevel int, message string) {
	_, file, line, _ := runtime.Caller(1)
	instance.getEventLog(logLevel).
		Time(timeKey, time.Now()).
		Str(fileKey, file).
		Int(lineKey, line).
		Msg(message)
}

func (instance TCMSLogger) getEventLog(level int) *zerolog.Event {
	switch level {
	case errors.DebugLevel:
		return instance.zeroLogger.Debug()
	case errors.WarnLevel:
		return instance.zeroLogger.Warn()
	case errors.ErrorLevel:
		return instance.zeroLogger.Error()
	case errors.FatalLevel:
		return instance.zeroLogger.Fatal()
	default:
		return instance.zeroLogger.Info()
	}
}

func New() *TCMSLogger {
	logger := zerolog.New(os.Stdout)
	return &TCMSLogger{zeroLogger: logger}
}
