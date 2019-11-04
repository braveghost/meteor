package itime

import (
	"time"
)

const (
	Second = 1
	Minute = 60 * Second
	Hour   = 60 * Minute
	Day    = 24 * Hour
)

func NowFormat(ft ...string) string {
	switch len(ft) {
	case 0:
		return time.Now().Format("2006-01-02 15:04:05")
	case 1:
		return time.Now().Format(ft[0])
	default:
		return "input format param is error"
	}
}
func NowMinuteStart() int64 {
	t := nowMinute()
	return t.Unix()
}
func nowMinute() time.Time {
	timeStr := time.Now().Format("2006-01-02 15:04")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t
}

func LastMinuteStart() int64 {
	return NowMinuteStart() - Minute
}
func LastMinutesStart(n int64) int64 {
	return NowMinuteStart() - n*Minute
}
func NextMinuteStart() int64 {
	return NowMinuteStart() + Minute
}
func NextMinutesStart(n int64) int64 {
	return NowMinuteStart() + n*Minute
}

func NowHourStart() int64 {
	t := nowHour()
	return t.Unix()
}
func nowHour() time.Time {
	timeStr := time.Now().Format("2006-01-02 15")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t
}

func LastHourStart() int64 {
	return NowHourStart() - Hour
}
func LastHoursStart(n int64) int64 {
	return NowHourStart() - n*Hour
}
func NextHourStart() int64 {
	return NowHourStart() + Hour
}
func NextHoursStart(n int64) int64 {
	return NowHourStart() + n*Hour
}

func NowDayStart() int64 {
	t := nowDay()
	return t.Unix()
}
func nowDay() time.Time {
	timeStr := time.Now().Format("2006-01-02")
	t, _ := time.Parse("2006-01-02", timeStr)
	return t
}

func LastDayStart() int64 {
	return NowDayStart() - Day
}

func LastDaysStart(n int64) int64 {
	return NowHourStart() - n*Day
}
func NextDayStart() int64 {
	return NowDayStart() + Day
}

func NextDaysStart(n int64) int64 {
	return NowDayStart() + n*Day
}

func NowMonthStart() int64 {
	t := nowMonth()
	return t.Unix()
}

func nowMonth() time.Time {
	t, _ := time.Parse("2006-01", time.Now().Format("2006-01"))
	return t
}

func LastMonthStart() int64 {
	return LastMonthsStart(1)
}

func LastMonthsStart(n int64) int64 {
	t := nowMonth()
	tt := t.AddDate(0, int(-n), 0)
	return tt.Unix()
}

func NextMonthStart() int64 {
	return NextMonthsStart(1)
}

func NextMonthsStart(n int64) int64 {
	t := nowMonth()
	tt := t.AddDate(0, int(n), 0)
	return tt.Unix()
}
