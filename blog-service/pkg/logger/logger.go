package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"runtime"
)

type Level int8

type Fields map[string]interface{}

const (
	LevelDebug Level = iota
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
	LevelPanic
)

func (level Level) String() string {
	switch level {
		case LevelDebug:
			return "debug"
		case LevelInfo:
			return "info"
		case LevelWarn:
			return "warn"
		case LevelError:
			return "error"
		case LevelFatal:
			return "fatal"
		case LevelPanic:
			return "panic"
	}
	return ""
}

type Logger struct {
	newLogger *log.Logger
	ctx context.Context
	level Level
	fields Fields
	callers []string
}

func NewLogger(writer io.Writer, prefix string, flag int) *Logger {
	logger := log.New(writer, prefix, flag)
	return &Logger{newLogger: logger}
}

func (logger *Logger) clone() *Logger {
	newLogger := *logger
	return &newLogger
}

func (logger *Logger) WithLevel(level Level) *Logger {
	newLogger := logger.clone()
	newLogger.level = level
	return newLogger
}

func (logger *Logger) WithFields(fields Fields) *Logger {
	newLogger := logger.clone()
	if newLogger.fields == nil {
		newLogger.fields = make(Fields)
	}
	for key, value := range fields {
		newLogger.fields[key] = value
	}
	return newLogger
}

func (logger *Logger) WithContext(ctx context.Context) *Logger {
	newLogger := logger.clone()
	newLogger.ctx = ctx
	return newLogger
}

func (logger *Logger) WithCaller(skip int) *Logger {
	newLogger := logger.clone()
	pc, file, line, status := runtime.Caller(skip)
	if status {
		f := runtime.FuncForPC(pc)
		newLogger.callers = []string{fmt.Sprintf("%s: %d %s", file, line, f.Name())}
	}
	return newLogger
}

func (logger *Logger) withCallersFrames() *Logger {
	maxCallerDepth := 25
	minCallerDepth := 1
	var callers []string
	pcs := make([]uintptr, maxCallerDepth)
	depth := runtime.Callers(minCallerDepth, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, more := frames.Next(); more; frame, more = frames.Next() {
		callers = append(callers, fmt.Sprintf("%s: %d %s", frame.File, frame.Line, frame.Function))
		if !more {
			break
		}
	}
	newLogger := logger.clone()
	newLogger.callers = callers
	return newLogger
}

