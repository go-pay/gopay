package xlog

const (
	ErrorLevel LogLevel = iota + 1
	WarnLevel
	InfoLevel
	DebugLevel
)

type LogLevel int

var (
	debugLog XLogger = &DebugLogger{}
	infoLog  XLogger = &InfoLogger{}
	warnLog  XLogger = &WarnLogger{}
	errLog   XLogger = &ErrorLogger{}

	Level LogLevel
)

type XLogger interface {
	LogOut(col *ColorType, format *string, args ...any)
}

func SetDebugLog(logger XLogger) {
	debugLog = logger
}

func SetInfoLog(logger XLogger) {
	infoLog = logger
}

func SetWarnLog(logger XLogger) {
	warnLog = logger
}

func SetErrLog(logger XLogger) {
	errLog = logger
}

func Info(args ...any) {
	infoLog.LogOut(nil, nil, args...)
}

func Infof(format string, args ...any) {
	infoLog.LogOut(nil, &format, args...)
}

func Debug(args ...any) {
	debugLog.LogOut(nil, nil, args...)
}

func Debugf(format string, args ...any) {
	debugLog.LogOut(nil, &format, args...)
}

func Warn(args ...any) {
	warnLog.LogOut(nil, nil, args...)
}

func Warnf(format string, args ...any) {
	warnLog.LogOut(nil, &format, args...)
}

func Error(args ...any) {
	errLog.LogOut(nil, nil, args...)
}

func Errorf(format string, args ...any) {
	errLog.LogOut(nil, &format, args...)
}
