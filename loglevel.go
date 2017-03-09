package golog

// Level represents a logging level
type Level interface {
	Name() string
	Level() int
	Color() int8
}

// level represents a built-in logging level
type level struct {
	name  string
	level int
	color int8
}

// Built in logging levels
var (
	LevelError Level = &level{name: "ERROR", level: 1000, color: 31}
	LevelWarn  Level = &level{name: "WARN", level: 2000, color: 33}
	LevelInfo  Level = &level{name: "INFO", level: 3000, color: 36}
	LevelDebug Level = &level{name: "DEBUG", level: 5000, color: 35}
)

func (l *level) Name() string {
	return l.name
}

func (l *level) Level() int {
	return l.level
}

func (l *level) Color() int8 {
	return l.color
}
