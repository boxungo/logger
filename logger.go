package logger

import (
	"fmt"
	"io"
	"sync"
)

// DebugLevel 调试等级
// InfoLevel 信息记录
// WarnLevel 警告
// ErrorLevel 错误
// FatalLevel 严重错误
const (
	DebugLevel = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	FatalLevel
)

var (
	levelName = map[int]string{
		DebugLevel: "DEBUG",
		InfoLevel:  "INFO",
		WarnLevel:  "WARN",
		ErrorLevel: "ERROR",
		FatalLevel: "FATAL",
	}
)

// Logger ...
type Logger struct {
	sync.RWMutex
	level         int
	out           io.Writer
	formatHandler func(out io.Writer, level int, arg string)
}

// New ...
func New(out io.Writer, level int) *Logger {
	l := &Logger{
		level:         level,
		out:           out,
		formatHandler: defaultFormatHandler,
	}
	return l
}

// SetLevel ...
func (l *Logger) SetLevel(level int) *Logger {
	l.level = level
	return l
}

// SetFormatHandler 设置日志格式化函数
func (l *Logger) SetFormatHandler(f func(out io.Writer, level int, arg string)) {
	l.formatHandler = f
}

// Debug 调试
func (l *Logger) Debug(arg string) {
	l.write(DebugLevel, arg)
}

// Info 信息
func (l *Logger) Info(arg string) {
	l.write(InfoLevel, arg)
}

// Warn 警告
func (l *Logger) Warn(arg string) {
	l.write(WarnLevel, arg)
}

// Error 错误
func (l *Logger) Error(arg string) {
	l.write(ErrorLevel, arg)
}

// Fatal 严重错误
func (l *Logger) Fatal(arg string) {
	l.write(FatalLevel, arg)
}

// Debugf 调试
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.writef(DebugLevel, format, args...)
}

// Infof 信息
func (l *Logger) Infof(format string, args ...interface{}) {
	l.writef(InfoLevel, format, args...)
}

// Warnf 警告
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.writef(WarnLevel, format, args...)
}

// Errorf 错误
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.writef(ErrorLevel, format, args...)
}

// Fatalf 严重错误
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.writef(FatalLevel, format, args...)
}

func (l *Logger) write(level int, arg string) {
	if l.level > level {
		return
	}
	l.Lock()
	defer l.Unlock()

	l.formatHandler(l.out, level, arg)
}

func (l *Logger) writef(level int, format string, args ...interface{}) {
	if l.level > level {
		return
	}

	str := fmt.Sprintf(format, args...)

	l.Lock()
	defer l.Unlock()

	l.formatHandler(l.out, level, str)
}
