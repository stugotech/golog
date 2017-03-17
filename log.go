package golog

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
	"sync"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	// remember if Stderr is a console, to enable coloured output (-1 = no, 0 = not checked, 1 = yes)
	stdErrIsConsole int
	consoleMutex    = &sync.Mutex{}
)

var (
	// CurrentConfig is the current log context options
	CurrentConfig Config
)

func init() {
	CurrentConfig = Config{
		IncludeCallerInfo: true,
		Level:             LevelDebug,
		Writer:            ConsoleWriter,
	}
}

// WriterFunc represents a function which can write a log
type WriterFunc func(tag string, level Level, msg string, fields []Field, frameDepth int)

// Config determines what is logged
type Config struct {
	IncludeCallerInfo bool
	Level             Level
	Writer            WriterFunc
}

// Logger has methods for standard logging levels
type Logger interface {
	Error(format string, fields ...Field) error
	Warn(format string, fields ...Field)
	Info(format string, fields ...Field)
	Debug(format string, fields ...Field)

	Errore(err error) error
	Errorex(desc string, err error, fields ...Field) error
}

type stdLogger struct {
	tag string
}

// NewLogger creates a new standard logger
func NewLogger(tag string) Logger {
	return &stdLogger{
		tag: tag,
	}
}

// NewPackageLogger creates a new logger for the caller package
func NewPackageLogger() Logger {
	frame := GetCallerFrame(1)
	function := frame.Function[:strings.LastIndex(frame.Function, ".")]
	return &stdLogger{
		tag: function,
	}
}

func (l *stdLogger) Error(msg string, fields ...Field) error {
	err := NewError(l.tag, msg)
	defaultWriter(l.tag, LevelError, msg, fields, 1)
	return err
}

func (l *stdLogger) Warn(msg string, fields ...Field) {
	defaultWriter(l.tag, LevelWarn, msg, fields, 1)
}

func (l *stdLogger) Info(msg string, fields ...Field) {
	defaultWriter(l.tag, LevelInfo, msg, fields, 1)
}

func (l *stdLogger) Debug(msg string, fields ...Field) {
	defaultWriter(l.tag, LevelDebug, msg, fields, 1)
}

func (l *stdLogger) Errore(err error) error {
	logged, _ := WrapError(l.tag, err)

	if !logged.Logged() {
		defaultWriter(l.tag, LevelError, err.Error(), nil, 1)
	}

	return logged
}

func (l *stdLogger) Errorex(desc string, err error, fields ...Field) error {
	err = fmt.Errorf("%s: %v", desc, err)
	logged, _ := WrapError(l.tag, err)
	defaultWriter(l.tag, LevelError, err.Error(), nil, 1)
	return logged
}

// ConsoleWriter writes logs to the console
func ConsoleWriter(tag string, level Level, msg string, fields []Field, frameDepth int) {
	if level.Level() <= CurrentConfig.Level.Level() {
		var callerinfo string

		if CurrentConfig.IncludeCallerInfo {
			_, file, line, ok := runtime.Caller(1 + frameDepth)
			if ok {
				callerinfo = fmt.Sprintf("/%s:%d", path.Base(file), line)
			}
		}

		msg = fmt.Sprintf("%s%s%s [%s] %s\n",
			tag,
			callerinfo,
			writeColor(level.Color()),
			level.Name(),
			msg,
		)

		for _, f := range fields {
			msg = fmt.Sprintf("%s\t%s = %s\n", msg, f.Name(), f.String())
		}

		msg += clearColor()
		consoleMutex.Lock()
		fmt.Fprint(os.Stderr, msg)
		consoleMutex.Unlock()
	}
}

// defaultWriter logs to CurrentConfig.Writer if set
func defaultWriter(tag string, level Level, msg string, fields []Field, frameDepth int) {
	if CurrentConfig.Writer != nil {
		frameDepth++
		CurrentConfig.Writer(tag, level, msg, fields, frameDepth)
	}
}

// GetCallerFrame gets the frame of the caller, if this info has been enabled via EnableCallerContext
func GetCallerFrame(skip int) *runtime.Frame {
	pc := make([]uintptr, 1)
	if n := runtime.Callers(2+skip, pc); n < 1 {
		return nil
	}
	frame, _ := runtime.CallersFrames(pc).Next()
	return &frame
}

func writeColor(color int8) string {
	return writeColorCode(fmt.Sprintf("\033[0;%dm", color))
}

func clearColor() string {
	return writeColorCode("\033[0m")
}

func writeColorCode(code string) string {
	if stdErrIsConsole == 0 {
		con := terminal.IsTerminal(int(os.Stderr.Fd()))
		if con {
			stdErrIsConsole = 1
		} else {
			stdErrIsConsole = -1
		}
	}
	if stdErrIsConsole > 0 {
		return code
	}
	return ""
}
