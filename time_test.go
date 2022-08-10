package almanac

import (
	"testing"
	"time"
)

func TestToJulianDay(t *testing.T) {
	var jd = (Time{2000, 1, 1, 12, 0, 0}).toJulianDay()
	if !isEqual(jd, 2451545.0) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2451545.0)
	}
	jd = (Time{1987, 1, 27, 0, 0, 0}).toJulianDay()
	if !isEqual(jd, 2446822.5) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2446822.5)
	}
	jd = (Time{1987, 6, 19, 12, 0, 0}).toJulianDay()
	if !isEqual(jd, 2446966.0) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2446966.0)
	}
	jd = (Time{1988, 1, 27, 0, 0, 0}).toJulianDay()
	if !isEqual(jd, 2447187.5) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2447187.5)
	}
	jd = (Time{1988, 6, 19, 12, 0, 0}).toJulianDay()
	if !isEqual(jd, 2447332.0) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2447332.0)
	}
	jd = (Time{1900, 1, 1, 0, 0, 0}).toJulianDay()
	if !isEqual(jd, 2415020.5) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2415020.5)
	}
	jd = (Time{1600, 1, 1, 0, 0, 0}).toJulianDay()
	if !isEqual(jd, 2305447.5) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2305447.5)
	}
	jd = (Time{1600, 12, 31, 0, 0, 0}).toJulianDay()
	if !isEqual(jd, 2305812.5) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2305812.5)
	}
	jd = (Time{837, 4, 10, 7, 12, 0}).toJulianDay()
	if !isEqual(jd, 2026871.8) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 2026871.8)
	}
	jd = (Time{-1000, 7, 12, 12, 0, 0}).toJulianDay()
	if !isEqual(jd, 1356001.0) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 1356001.0)
	}
	jd = (Time{-1000, 2, 29, 0, 0, 0}).toJulianDay()
	if !isEqual(jd, 1355866.5) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 1355866.5)
	}
	jd = (Time{-1001, 8, 17, 21, 36, 0}).toJulianDay()
	if !isEqual(jd, 1355671.4) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 1355671.4)
	}
	jd = (Time{-4712, 1, 1, 12, 0, 0}).toJulianDay()
	if !isEqual(jd, 0.0) {
		t.Error(jd, "toJulianDay() Error:", jd, ", But expect:", 0.0)
	}
}

func TestGetDaysOffJ2000(t *testing.T) {
	var tt = Time{2000, 1, 1, 12, 0, 0}
	var off = tt.getDaysOffJ2000()
	if off != 0 {
		t.Error("j2000 Days off error:", off, "But Expect:", 0)
	}

	tt = Time{2022, 7, 14, 12, 0, 0}
	off = tt.getDaysOffJ2000()
	if off != 8230 {
		t.Error("2022-7-14 12:00:00 Days off error:", off, "But Expect:", 8230)
	}

	tt = Time{1582, 10, 1, 12, 0, 0}
	off = tt.getDaysOffJ2000()
	if off != -152388 {
		t.Error("1582-10-1 12:00:00 Days off error:", off, "But Expect:", -152388)
	}

	tt = Time{1582, 10, 4, 12, 0, 0}
	off = tt.getDaysOffJ2000()
	if off != -152385 {
		t.Error("1582-10-4 12:00:00 Days off error:", off, "But Expect:", -152385)
	}

	tt = Time{1582, 10, 15, 12, 0, 0}
	off = tt.getDaysOffJ2000()
	if off != -152384 {
		t.Error("1582-10-15 12:00:00 Days off error:", off, "But Expect:", -152384)
	}

	tt = Time{1582, 10, 18, 12, 0, 0}
	off = tt.getDaysOffJ2000()
	if off != -152381 {
		t.Error("1582-10-18 12:00:00 Days off error:", off, "But Expect:", -152381)
	}

	tt = Time{1582, 10, 31, 12, 0, 0}
	off = tt.getDaysOffJ2000()
	if off != -152368 {
		t.Error("1582-10-31 12:00:00 Days off error:", off, "But Expect:", -152368)
	}

	tt = Time{9999, 12, 31, 12, 0, 0}
	off = tt.getDaysOffJ2000()
	if off != 2921939 {
		t.Error("9999-12-31 12:00:00 Days off error:", off, "But Expect:", 2921939)
	}
}

func TestGetMonthFirstDayTime(t *testing.T) {
	var tt = Time{1582, 7, 13, 11, 30, 0}
	var result = tt.getMonthFirstDayTime()
	if result != (Time{1582, 7, 1, 12, 0, 0}) {
		t.Error("Time{1582, 7, 13, 11, 30, 0} getMonthFirstDayTime error:", result, "But Expect:", Time{1582, 7, 1, 12, 0, 0})
	}

	tt = Time{-999, 7, 13, 11, 30, 0}
	result = tt.getMonthFirstDayTime()
	if result != (Time{-999, 7, 1, 12, 0, 0}) {
		t.Error("Time{-999, 7, 13, 11, 30, 0} getMonthFirstDayTime error:", result, "But Expect:", Time{-999, 7, 1, 12, 0, 0})
	}
}

func TestGetMonthDaysNum(t *testing.T) {
	var tt = Time{2022, 7, 14, 0, 0, 0}
	var sum = tt.getMonthDaysNum()

	if sum != 31 || getMonthDaysNum(tt, tt.getMonthFirstDaysOffJ2000()) != 31 {
		t.Error(tt, "getMonthDaysNum error:", sum, "But Expect:", 31)
	}

	tt = Time{2022, 2, 14, 0, 0, 0}
	sum = tt.getMonthDaysNum()

	if sum != 28 || getMonthDaysNum(tt, tt.getMonthFirstDaysOffJ2000()) != 28 {
		t.Error(tt, "getMonthDaysNum error:", sum, "But Expect:", 28)
	}

	tt = Time{1582, 10, 6, 0, 0, 0}
	sum = tt.getMonthDaysNum()

	if sum != 21 || getMonthDaysNum(tt, tt.getMonthFirstDaysOffJ2000()) != 21 {
		t.Error(tt, "getMonthDaysNum error:", sum, "But Expect:", 21)
	}

	tt = Time{1582, 10, 1, 0, 0, 0}
	sum = tt.getMonthDaysNum()

	if sum != 21 || getMonthDaysNum(tt, tt.getMonthFirstDaysOffJ2000()) != 21 {
		t.Error(tt, "getMonthDaysNum error:", sum, "But Expect:", 21)
	}

	tt = Time{1582, 10, 15, 0, 0, 0}
	sum = tt.getMonthDaysNum()

	if sum != 21 || getMonthDaysNum(tt, tt.getMonthFirstDaysOffJ2000()) != 21 {
		t.Error(tt, "getMonthDaysNum error:", sum, "But Expect:", 21)
	}

	tt = Time{-1000, 2, 1, 0, 0, 0}
	sum = tt.getMonthDaysNum()

	if sum != 29 || getMonthDaysNum(tt, tt.getMonthFirstDaysOffJ2000()) != 29 {
		t.Error(tt, "getMonthDaysNum error:", sum, "But Expect:", 29)
	}
}

func TestGetMonthFirstDayInfo(t *testing.T) {
	var tt = Time{2022, 7, 14, 0, 0, 0}
	var _, numberDaysInMonth, firstWeekday = tt.getMonthFirstDayInfo()
	var rNumberDaysInMonth, rFirstWeekday = 31, 5
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(tt, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	tt = Time{2022, 2, 14, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = tt.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 28, 2
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(tt, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	tt = Time{1582, 10, 6, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = tt.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 21, 1
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(tt, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	tt = Time{1582, 9, 30, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = tt.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 30, 6
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(tt, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	tt = Time{1582, 1, 1, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = tt.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 31, 1
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(tt, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	tt = Time{-1000, 2, 1, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = tt.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 29, 3
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(tt, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	tt = Time{9999, 12, 31, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = tt.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 31, 3
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(tt, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

}

// 获取当前时间
func TestTimeNow(t *testing.T) {
	var t1 = TimeNow()
	var now = time.Now()
	var e = Time{year: now.Year(), month: int(now.Month()), day: now.Day(), hour: now.Hour(), minute: now.Minute(), second: now.Second()}
	if t1 != e {
		t.Error("TimeNow may")
	}
}
