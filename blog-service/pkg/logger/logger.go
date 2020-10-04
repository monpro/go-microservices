package logger

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

