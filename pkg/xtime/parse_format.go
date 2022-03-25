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
		intNs = int64(duration)
	)
	if intNs >= 0 && intNs < int64(time.Second) {
		t = util.Int642String(intNs/int64(time.Millisecond)) + "ms"
	}

	// 大于等于 1秒，小于 1分钟
	if intNs >= int64(time.Second) && intNs < int64(time.Minute) {
		s := intNs / int64(time.Second)
		ms := (intNs - s*int64(time.Second)) / int64(time.Millisecond)
		t = util.Int642String(s) + "s"
		if ms > 0 {
			t += util.Int642String(ms) + "ms"
		}
	}
	// 大于等于 1分钟，小于 1小时
	if intNs >= int64(time.Minute) && intNs < int64(time.Hour) {
		m := intNs / int64(time.Minute)
		s := (intNs - m*int64(time.Minute)) / int64(time.Second)
		t = util.Int642String(m) + "m"
		if s > 0 {
			t += util.Int642String(s) + "s"
		}
	}
	// 大于等于 1小时，小于 1天
	if intNs >= int64(time.Hour) && intNs < 24*int64(time.Hour) {
		h := intNs / int64(time.Hour)
		m := (intNs - h*int64(time.Hour)) / int64(time.Minute)
		s := (intNs - h*int64(time.Hour) - m*int64(time.Minute)) / int64(time.Second)
		t = util.Int642String(h) + "h"
		if m > 0 {
			t += util.Int642String(m) + "m"
		}
		if s > 0 {
			t += util.Int642String(s) + "s"
		}
	}
	// 大于等于 1天
	if intNs >= 24*int64(time.Hour) {
		d := intNs / (24 * int64(time.Hour))
		h := (intNs - d*24*int64(time.Hour)) / int64(time.Hour)
		m := (intNs - d*24*int64(time.Hour) - h*int64(time.Hour)) / int64(time.Minute)
		s := ((intNs - m*int64(time.Minute)) % int64(time.Minute)) / int64(time.Second)

		t = util.Int642String(d) + "d"
		if h > 0 {
			t += util.Int642String(h) + "h"
		}
		if m > 0 {
			t += util.Int642String(m) + "m"
		}
		if s > 0 {
			t += util.Int642String(s) + "s"
		}
	}
	return t
}
