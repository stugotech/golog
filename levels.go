package golog

// Level is the severity of the log entry
type Level interface {
	Name() string
	Level() int
}

// level represents a built-in log level
type level struct {
	name  string
	level int
}

// The following log levels are built in
var (
	LevelError Level = &level{name: "ERROR", level: 5000}
	LevelWarn  Level = &level{name: "WARN", level: 6000}
	LevelInfo  Level = &level{name: "INFO", level: 7000}
	LevelDebug Level = &level{name: "DEBUG", level: 8000}
)

func (l *level) Name() string {
	return l.name
}

func (l *level) Level() int {
	return l.level
}
