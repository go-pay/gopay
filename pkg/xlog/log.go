package xlog

var (
	debugLog XLogger = &DebugLogger{}
	infoLog  XLogger = &InfoLogger{}
	warnLog  XLogger = &WarnLogger{}
	errLog   XLogger = &ErrorLogger{}
)

type XLogger interface {
	LogOut(col *ColorType, format *string, args ...interface{})
}

func Info(args ...interface{}) {
	infoLog.LogOut(nil, nil, args...)
}

func Infof(format string, args ...interface{}) {
	infoLog.LogOut(nil, &format, args...)
}

func Debug(args ...interface{}) {
	debugLog.LogOut(nil, nil, args...)
}

func Debugf(format string, args ...interface{}) {
	debugLog.LogOut(nil, &format, args...)
}

func Warn(args ...interface{}) {
	warnLog.LogOut(nil, nil, args...)
}

func Warnf(format string, args ...interface{}) {
	warnLog.LogOut(nil, &format, args...)
}

func Error(args ...interface{}) {
	errLog.LogOut(nil, nil, args...)
}

func Errorf(format string, args ...interface{}) {
	errLog.LogOut(nil, &format, args...)
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
