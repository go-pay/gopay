package xlog

var (
	debugLog xLogger = &DebugLogger{}
	infoLog  xLogger = &InfoLogger{}
	warnLog  xLogger = &WarnLogger{}
	errLog   xLogger = &ErrorLogger{}
)

type xLogger interface {
	logOut(col *ColorType, format *string, args ...interface{})
}

func Info(args ...interface{}) {
	infoLog.logOut(nil, nil, args...)
}

func Infof(format string, args ...interface{}) {
	infoLog.logOut(nil, &format, args...)
}

func Debug(args ...interface{}) {
	debugLog.logOut(nil, nil, args...)
}

func Debugf(format string, args ...interface{}) {
	debugLog.logOut(nil, &format, args...)
}

func Warn(args ...interface{}) {
	warnLog.logOut(nil, nil, args...)
}

func Warnf(format string, args ...interface{}) {
	warnLog.logOut(nil, &format, args...)
}

func Error(args ...interface{}) {
	errLog.logOut(nil, nil, args...)
}

func Errorf(format string, args ...interface{}) {
	errLog.logOut(nil, &format, args...)
}
