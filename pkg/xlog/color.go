package xlog

type ColorType string

var (
	Reset = ColorType([]byte{27, 91, 48, 109})
	// 标准
	White   = ColorType([]byte{27, 91, 51, 48, 109}) // 白色
	Red     = ColorType([]byte{27, 91, 51, 49, 109}) // 红色
	Green   = ColorType([]byte{27, 91, 51, 50, 109}) // 绿色
	Yellow  = ColorType([]byte{27, 91, 51, 51, 109}) // 黄色
	Blue    = ColorType([]byte{27, 91, 51, 52, 109}) // 蓝色
	Magenta = ColorType([]byte{27, 91, 51, 53, 109}) // 紫色
	Cyan    = ColorType([]byte{27, 91, 51, 54, 109}) // 青色
	// 高亮
	WhiteBright   = ColorType([]byte{27, 91, 49, 59, 51, 48, 109})
	RedBright     = ColorType([]byte{27, 91, 49, 59, 51, 49, 109})
	GreenBright   = ColorType([]byte{27, 91, 49, 59, 51, 50, 109})
	YellowBright  = ColorType([]byte{27, 91, 49, 59, 51, 51, 109})
	BlueBright    = ColorType([]byte{27, 91, 49, 59, 51, 52, 109})
	MagentaBright = ColorType([]byte{27, 91, 49, 59, 51, 53, 109})
	CyanBright    = ColorType([]byte{27, 91, 49, 59, 51, 54, 109})
	// 斜体
	WhiteBevel   = ColorType([]byte{27, 91, 51, 59, 51, 48, 109})
	RedBevel     = ColorType([]byte{27, 91, 51, 59, 51, 49, 109})
	GreenBevel   = ColorType([]byte{27, 91, 51, 59, 51, 50, 109})
	YellowBevel  = ColorType([]byte{27, 91, 51, 59, 51, 51, 109})
	BlueBevel    = ColorType([]byte{27, 91, 51, 59, 51, 52, 109})
	MagentaBevel = ColorType([]byte{27, 91, 51, 59, 51, 53, 109})
	CyanBevel    = ColorType([]byte{27, 91, 51, 59, 51, 54, 109})
	// 下划线
	WhiteUnderLine   = ColorType([]byte{27, 91, 52, 59, 51, 48, 109})
	RedUnderLine     = ColorType([]byte{27, 91, 52, 59, 51, 49, 109})
	GreenUnderLine   = ColorType([]byte{27, 91, 52, 59, 51, 50, 109})
	YellowUnderLine  = ColorType([]byte{27, 91, 52, 59, 51, 51, 109})
	BlueUnderLine    = ColorType([]byte{27, 91, 52, 59, 51, 52, 109})
	MagentaUnderLine = ColorType([]byte{27, 91, 52, 59, 51, 53, 109})
	CyanUnderLine    = ColorType([]byte{27, 91, 52, 59, 51, 54, 109})
	// 背景色
	WhiteBg   = ColorType([]byte{27, 91, 55, 59, 51, 48, 109})
	RedBg     = ColorType([]byte{27, 91, 55, 59, 51, 49, 109})
	GreenBg   = ColorType([]byte{27, 91, 55, 59, 51, 50, 109})
	YellowBg  = ColorType([]byte{27, 91, 55, 59, 51, 51, 109})
	BlueBg    = ColorType([]byte{27, 91, 55, 59, 51, 52, 109})
	MagentaBg = ColorType([]byte{27, 91, 55, 59, 51, 53, 109})
	CyanBg    = ColorType([]byte{27, 91, 55, 59, 51, 54, 109})
	// 删除线
	WhiteDelLine   = ColorType([]byte{27, 91, 57, 59, 51, 48, 109})
	RedDelLine     = ColorType([]byte{27, 91, 57, 59, 51, 49, 109})
	GreenDelLine   = ColorType([]byte{27, 91, 57, 59, 51, 50, 109})
	YellowDelLine  = ColorType([]byte{27, 91, 57, 59, 51, 51, 109})
	BlueDelLine    = ColorType([]byte{27, 91, 57, 59, 51, 52, 109})
	MagentaDelLine = ColorType([]byte{27, 91, 57, 59, 51, 53, 109})
	CyanDelLine    = ColorType([]byte{27, 91, 57, 59, 51, 54, 109})
)

var cl *ColorLogger

type ColorLogger struct {
	Color ColorType
	i     *InfoLogger
	d     *DebugLogger
	w     *WarnLogger
	e     *ErrorLogger
}

func Color(color ColorType) *ColorLogger {
	if cl == nil {
		cl = &ColorLogger{
			Color: color,
			i:     &InfoLogger{},
			d:     &DebugLogger{},
			w:     &WarnLogger{},
			e:     &ErrorLogger{},
		}
		return cl
	}
	cl.Color = color
	return cl
}

func (l *ColorLogger) Info(args ...interface{}) {
	l.i.logOut(&l.Color, nil, args...)
}

func (l *ColorLogger) Infof(format string, args ...interface{}) {
	l.i.logOut(&l.Color, &format, args...)
}

func (l *ColorLogger) Debug(args ...interface{}) {
	l.d.logOut(&l.Color, nil, args...)
}

func (l *ColorLogger) Debugf(format string, args ...interface{}) {
	l.d.logOut(&l.Color, &format, args...)
}

func (l *ColorLogger) Warn(args ...interface{}) {
	l.w.logOut(&l.Color, nil, args...)
}

func (l *ColorLogger) Warnf(format string, args ...interface{}) {
	l.w.logOut(&l.Color, &format, args...)
}

func (l *ColorLogger) Error(args ...interface{}) {
	l.e.logOut(&l.Color, nil, args...)
}

func (l *ColorLogger) Errorf(format string, args ...interface{}) {
	l.e.logOut(&l.Color, &format, args...)
}
