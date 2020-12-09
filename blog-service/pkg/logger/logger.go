package logger

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"runtime"
	"time"
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
	ctx       context.Context
	level     Level
	fields    Fields
	callers   []string
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

func (logger *Logger) WithCallersFrames() *Logger {
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

func (logger *Logger) WithTrace() *Logger {
	context, status := logger.ctx.(*gin.Context)
	if status {
		return logger.WithFields(Fields{
			"trace_id": context.MustGet("X-Trace-ID"),
			"span_id":  context.MustGet("X-Span-ID"),
		})
	}
	return logger
}

func (logger *Logger) JSONFormat(msg string) map[string]interface{} {
	data := make(Fields, len(logger.fields)+4)
	data["level"] = logger.level.String()
	data["time"] = time.Now().Local().UnixNano()
	data["message"] = msg
	data["callers"] = logger.callers
	if len(logger.fields) > 0 {
	}
	for key, value := range logger.fields {
		if _, status := data[key]; !status {
			data[key] = value
		}
	}
	return data
}

func (logger *Logger) Output(level Level, msg string) {
	body, _ := json.Marshal(logger.JSONFormat(msg))
	content := string(body)
	switch logger.level {
	case LevelDebug:
		logger.newLogger.Print(content)
	case LevelInfo:
		logger.newLogger.Printf(content)
	case LevelWarn:
		logger.newLogger.Printf(content)
	case LevelError:
		logger.newLogger.Print(content)
	case LevelFatal:
		logger.newLogger.Fatalf(content)
	case LevelPanic:
		logger.newLogger.Panic(content)
	}
}

func (logger *Logger) Debug(ctx context.Context, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprint(v...))
}

func (logger *Logger) DebugFormat(ctx context.Context, format string, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelDebug, fmt.Sprintf(format, v...))
}

func (logger *Logger) Info(ctx context.Context, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

func (logger *Logger) InfoFormat(ctx context.Context, format string, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprintf(format, v...))
}

func (logger *Logger) Warn(ctx context.Context, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

func (logger *Logger) WarnFormat(ctx context.Context, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelInfo, fmt.Sprint(v...))
}

func (logger *Logger) Error(ctx context.Context, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprint(v...))
}

func (logger *Logger) ErrorFormat(ctx context.Context, format string, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelError, fmt.Sprintf(format, v...))
}

func (logger *Logger) Fatal(ctx context.Context, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprint(v...))
}

func (logger *Logger) FatalFormat(ctx context.Context, format string, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelFatal, fmt.Sprintf(format, v...))
}

func (logger *Logger) Panic(ctx context.Context, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprint(v...))
}

func (logger *Logger) PanicFormat(ctx context.Context, format string, v ...interface{}) {
	logger.WithContext(ctx).WithTrace().Output(LevelPanic, fmt.Sprintf(format, v...))
}
