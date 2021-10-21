package xtime

import (
	"strings"
	"time"

	"github.com/go-pay/gopay/pkg/util"
)

//解析时间
//    时间字符串格式：2006-01-02 15:04:05
func ParseDateTime(timeStr string) (datetime time.Time) {
	datetime, _ = time.ParseInLocation(TimeLayout, timeStr, time.Local)
	return
}

//解析日期
//    日期字符串格式：2006-01-02
func ParseDate(timeStr string) (date time.Time) {
	date, _ = time.ParseInLocation(DateLayout, timeStr, time.Local)
	return
}

//格式化Datetime字符串
//    格式化前输入样式：2019-01-04T15:40:00Z 或 2019-01-04T15:40:00+08:00
//    格式化后返回样式：2019-01-04 15:40:00
func FormatDateTime(timeStr string) (formatTime string) {
	if timeStr == "" {
		return ""
	}
	replace := strings.Replace(timeStr, "T", " ", 1)
	formatTime = replace[:19]
	return
}

//格式化Date成字符串
//    格式化前输入样式：2019-01-04T15:40:00Z 或 2019-01-04T15:40:00+08:00
//    格式化后返回样式：2019-01-04
func FormatDate(dateStr string) (formatDate string) {
	if dateStr == "" {
		return ""
	}
	split := strings.Split(dateStr, "T")
	formatDate = split[0]
	return
}

func DurationToUnit(duration time.Duration) string {
	var (
		t     string
		intNs = int(duration)
	)
	if intNs >= 0 && intNs < int(time.Second) {
		t = util.Int2String(intNs/int(time.Millisecond)) + "ms"
	}

	// 大于等于 1秒，小于 1分钟
	if intNs >= int(time.Second) && intNs < int(time.Minute) {
		s := intNs / int(time.Second)
		ms := (intNs - s*int(time.Second)) / int(time.Millisecond)
		t = util.Int2String(s) + "s"
		if ms > 0 {
			t += util.Int2String(ms) + "ms"
		}
	}
	// 大于等于 1分钟，小于 1小时
	if intNs >= int(time.Minute) && intNs < int(time.Hour) {
		m := intNs / int(time.Minute)
		s := (intNs - m*int(time.Minute)) / int(time.Second)
		t = util.Int2String(m) + "m"
		if s > 0 {
			t += util.Int2String(s) + "s"
		}
	}
	// 大于等于 1小时，小于 1天
	if intNs >= int(time.Hour) && intNs < 24*int(time.Hour) {
		h := intNs / int(time.Hour)
		m := (intNs - h*int(time.Hour)) / int(time.Minute)
		s := (intNs - h*int(time.Hour) - m*int(time.Minute)) / int(time.Second)
		t = util.Int2String(h) + "h"
		if m > 0 {
			t += util.Int2String(m) + "m"
		}
		if s > 0 {
			t += util.Int2String(s) + "s"
		}
	}
	// 大于等于 1天
	if intNs >= 24*int(time.Hour) {
		d := intNs / (24 * int(time.Hour))
		h := (intNs - d*24*int(time.Hour)) / int(time.Hour)
		m := (intNs - d*24*int(time.Hour) - h*int(time.Hour)) / int(time.Minute)
		s := ((intNs - m*int(time.Minute)) % int(time.Minute)) / int(time.Second)

		t = util.Int2String(d) + "d"
		if h > 0 {
			t += util.Int2String(h) + "h"
		}
		if m > 0 {
			t += util.Int2String(m) + "m"
		}
		if s > 0 {
			t += util.Int2String(s) + "s"
		}
	}
	return t
}
