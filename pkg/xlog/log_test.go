package xlog

import (
	"testing"
)

func TestLog(t *testing.T) {

	// default log
	Info(White, "白色 info", Reset, WhiteBright, "高亮 info", Reset, "恢复默认颜色", WhiteDelLine, "删除线", Reset, WhiteUnderLine, "下划线", Reset, WhiteBevel, "斜体 info", Reset, WhiteBg, "背景", Reset)
	Debug(Cyan, "青色 debug", Reset, CyanBright, "高亮 debug", Reset, "恢复默认颜色", CyanDelLine, "删除线", Reset, CyanUnderLine, "下划线", Reset, CyanBevel, "斜体 debug", Reset, CyanBg, "背景", Reset)
	Warn(Yellow, "黄色 warning", Reset, YellowBright, "高亮 warning", Reset, "恢复默认颜色", YellowDelLine, "删除线", Reset, YellowUnderLine, "下划线", Reset, YellowBevel, "斜体 warning", Reset, YellowBg, "背景", Reset)
	Error(Red, "红色 error", Reset, RedBright, "高亮 error", Reset, "恢复默认颜色", RedDelLine, "删除线", Reset, RedUnderLine, "下划线", Reset, RedBevel, "斜体 error", Reset, RedBg, "背景", Reset)

	// color log
	Color(White).Info("color log info")
	Color(Cyan).Debug("color log debug")
	Color(Yellow).Warn("color log warn")
	Color(Red).Error("color log error")
}
