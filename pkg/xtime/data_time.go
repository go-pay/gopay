package xtime

import (
	"time"
)

var daysBefore = [...]int32{
	0,
	31,
	31 + 28,
	31 + 28 + 31,
	31 + 28 + 31 + 30,
	31 + 28 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30,
	31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 + 30 + 31 + 30 + 31,
}

func MonthDays(m time.Month, year int) int {
	if m == time.February && isLeap(year) {
		return 29
	}
	return int(daysBefore[m] - daysBefore[m-1])
}

func isLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// 获取最近 7天 日期
func GetRecentSevenDay() (sevenDays []string) {
	now := time.Now()
	nowDay := time.Date(now.Year(), now.Month(), now.Day()-7, 0, 0, 0, 0, time.Local)
	for i := 0; i < 7; i++ {
		date := nowDay.AddDate(0, 0, i)
		sevenDays = append(sevenDays, date.Format(DateLayout))
	}
	return sevenDays
}

// 获取最近 30天 日期
func GetRecentThirtyDay() (thirtyDays []string) {
	now := time.Now()
	nowDay := time.Date(now.Year(), now.Month(), now.Day()-30, 0, 0, 0, 0, time.Local)
	for i := 0; i < 30; i++ {
		date := nowDay.AddDate(0, 0, i)
		thirtyDays = append(thirtyDays, date.Format(DateLayout))
	}
	return thirtyDays
}

// 获取本周 7天 日期
func GetCurWeekDays() (curWeekDays []string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	if offset > 0 {
		offset = -6
	}
	weekStartDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	for i := 0; i < 7; i++ {
		date := weekStartDate.AddDate(0, 0, i)
		curWeekDays = append(curWeekDays, date.Format(DateLayout))
	}
	return
}

// 获取本月 日期
func GetCurMonthDays() (curMonthDays []string) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	days := MonthDays(month, year)

	monthFirstDay := time.Date(year, month, 01, 0, 0, 0, 0, time.Local)
	for i := 0; i < days; i++ {
		date := monthFirstDay.AddDate(0, 0, i)
		curMonthDays = append(curMonthDays, date.Format(DateLayout))
	}
	return curMonthDays
}

// 获取上一个月 日期
func GetLastMonthDays() (monthDays []string) {
	now := time.Now()
	year := now.Year()
	month := now.Month()
	days := MonthDays(month-1, year)

	monthFirstDay := time.Date(year, month-1, 01, 0, 0, 0, 0, time.Local)
	for i := 0; i < days; i++ {
		date := monthFirstDay.AddDate(0, 0, i)
		monthDays = append(monthDays, date.Format(DateLayout))
	}
	return monthDays
}
