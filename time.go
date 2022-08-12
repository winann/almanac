package almanac

import "time"

//// Weekday 表明当前是星期几（日：0...）
//type Weekday int
//
//const (
//	Sunday Weekday = iota
//	Monday
//	Tuesday
//	Wednesday
//	Thursday
//	Friday
//	Saturday
//)

// Time 时间类
type Time struct {
	year, month, day, hour, minute, second int
}

// NewTime 创建时间
// 为 iOS 提供的接口
func NewTime(year, month, day, hour, minute, second int) (t *Time) {
	t = &Time{year, month, day, hour, minute, second}
	return
}

// TimeNow 获取当前时间
func TimeNow() (t Time) {
	var sysTime = time.Now()
	t = Time{sysTime.Year(), int(sysTime.Month()), sysTime.Day(), sysTime.Hour(), sysTime.Minute(), sysTime.Second()}
	return
}

// toJulianDay 时间实例转儒略日
// 返回儒略日
func (time Time) toJulianDay() (jd JulianDay) {
	jd = julianDay(time.year, time.month, float64(time.day)+((float64(time.second)/60.0+float64(time.minute))/60.0+float64(time.hour))/24)
	return
}

// getDaysOffJ2000 获取当前时间相对于 2000-1-1 的偏移
func (time Time) getDaysOffJ2000() (daysOff int) {
	daysOff = int(time.toJulianDay()) - j2000
	return
}

// 通过 getMonthFirstDay 获取当前日期所在月的月首时间
func (time Time) getYearFirstDayTime() (firstDayTime Time) {
	firstDayTime = time
	firstDayTime.month = 1
	firstDayTime.day = 1
	firstDayTime.hour = 12
	firstDayTime.minute = 0
	firstDayTime.second = 0
	return
}

// 通过 getMonthFirstDay 获取当前日期所在月的月首时间
func (time Time) getMonthFirstDayTime() (firstDayTime Time) {
	firstDayTime = time
	firstDayTime.day = 1
	firstDayTime.hour = 12
	firstDayTime.minute = 0
	firstDayTime.second = 0
	return
}

// getMonthFirstDaysOffJ2000 获取当前所在月的第一天相对于 2000-1-1 的偏移
func (time Time) getMonthFirstDaysOffJ2000() (daysOff int) {
	var firstDay = time.getMonthFirstDayTime()
	daysOff = firstDay.getDaysOffJ2000()
	return
}

// getMonthDaysNum 获取当前日期所在月的天数
func (time Time) getMonthDaysNum() (numberDaysOfMonth int) {
	numberDaysOfMonth = getMonthDaysNum(time, time.getMonthFirstDaysOffJ2000())
	return
}

/// getMonthFirstDay 获取当前日期所在月的第一天的 信息
func (time Time) getMonthFirstDayInfo() (firstDayOff int, numberDaysInMonth int, firstWeekday int) {
	var firstDayTime = time.getMonthFirstDayTime()
	firstDayOff = firstDayTime.getMonthFirstDaysOffJ2000()
	numberDaysInMonth = firstDayTime.getMonthDaysNum()
	firstWeekday = (firstDayOff + j2000 + 1 + 7000000) % 7
	return
}

// getMonthDaysNum 通过 Time 获取到某个月的公历天数
func getMonthDaysNum(firstDayTime Time, firstDayOff int) (numberDaysOfMonth int) {

	// 计算本月的天数 numberDaysOfMonth
	firstDayTime.month++
	if firstDayTime.month > 12 {
		firstDayTime.year++
		firstDayTime.month = 1
	}
	numberDaysOfMonth = firstDayTime.getMonthFirstDaysOffJ2000() - firstDayOff
	return
}
