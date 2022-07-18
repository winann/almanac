package almanac

import (
	"testing"
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
	var time = Time{2000, 1, 1, 12, 0, 0}
	var off = time.getMonthFirstDaysOffJ2000()
	if off != 0 {
		t.Error("J2000 days off error:", off, "But Expect:", 0)
	}

	time = Time{2022, 7, 14, 12, 0, 0}
	off = time.getMonthFirstDaysOffJ2000()
	if off != 8230 {
		t.Error("2022-7-14 12:00:00 days off error:", off, "But Expect:", 8230)
	}

	time = Time{1582, 10, 1, 12, 0, 0}
	off = time.getMonthFirstDaysOffJ2000()
	if off != -152388 {
		t.Error("1582-10-1 12:00:00 days off error:", off, "But Expect:", -152388)
	}

	time = Time{1582, 10, 4, 12, 0, 0}
	off = time.getMonthFirstDaysOffJ2000()
	if off != -152385 {
		t.Error("1582-10-4 12:00:00 days off error:", off, "But Expect:", -152385)
	}

	time = Time{1582, 10, 15, 12, 0, 0}
	off = time.getMonthFirstDaysOffJ2000()
	if off != -152384 {
		t.Error("1582-10-15 12:00:00 days off error:", off, "But Expect:", -152384)
	}

	time = Time{1582, 10, 18, 12, 0, 0}
	off = time.getMonthFirstDaysOffJ2000()
	if off != -152381 {
		t.Error("1582-10-18 12:00:00 days off error:", off, "But Expect:", -152381)
	}

	time = Time{1582, 10, 31, 12, 0, 0}
	off = time.getMonthFirstDaysOffJ2000()
	if off != -152368 {
		t.Error("1582-10-31 12:00:00 days off error:", off, "But Expect:", -152368)
	}

	time = Time{9999, 12, 31, 12, 0, 0}
	off = time.getMonthFirstDaysOffJ2000()
	if off != 2921939 {
		t.Error("9999-12-31 12:00:00 days off error:", off, "But Expect:", 2921939)
	}
}

func TestGetMonthFirstDayTime(t *testing.T) {
	var time = Time{1582, 7, 13, 11, 30, 0}
	var result = time.getMonthFirstDayTime()
	if result != (Time{1582, 7, 1, 12, 0, 0}) {
		t.Error("Time{1582, 7, 13, 11, 30, 0} getMonthFirstDayTime error:", result, "But Expect:", Time{1582, 7, 1, 12, 0, 0})
	}

	time = Time{-999, 7, 13, 11, 30, 0}
	result = time.getMonthFirstDayTime()
	if result != (Time{-999, 7, 1, 12, 0, 0}) {
		t.Error("Time{-999, 7, 13, 11, 30, 0} getMonthFirstDayTime error:", result, "But Expect:", Time{-999, 7, 1, 12, 0, 0})
	}
}

func TestGetMonthDaysNum(t *testing.T) {
	var time = Time{2022, 7, 14, 0, 0, 0}
	var sum = time.getMonthDaysNum()

	if sum != 31 || getMonthDaysNum(time, time.getMonthFirstDaysOffJ2000()) != 31 {
		t.Error(time, "getMonthDaysNum error:", sum, "But Expect:", 31)
	}

	time = Time{2022, 2, 14, 0, 0, 0}
	sum = time.getMonthDaysNum()

	if sum != 28 || getMonthDaysNum(time, time.getMonthFirstDaysOffJ2000()) != 28 {
		t.Error(time, "getMonthDaysNum error:", sum, "But Expect:", 28)
	}

	time = Time{1582, 10, 6, 0, 0, 0}
	sum = time.getMonthDaysNum()

	if sum != 21 || getMonthDaysNum(time, time.getMonthFirstDaysOffJ2000()) != 21 {
		t.Error(time, "getMonthDaysNum error:", sum, "But Expect:", 21)
	}

	time = Time{1582, 10, 1, 0, 0, 0}
	sum = time.getMonthDaysNum()

	if sum != 21 || getMonthDaysNum(time, time.getMonthFirstDaysOffJ2000()) != 21 {
		t.Error(time, "getMonthDaysNum error:", sum, "But Expect:", 21)
	}

	time = Time{1582, 10, 15, 0, 0, 0}
	sum = time.getMonthDaysNum()

	if sum != 21 || getMonthDaysNum(time, time.getMonthFirstDaysOffJ2000()) != 21 {
		t.Error(time, "getMonthDaysNum error:", sum, "But Expect:", 21)
	}

	time = Time{-1000, 2, 1, 0, 0, 0}
	sum = time.getMonthDaysNum()

	if sum != 29 || getMonthDaysNum(time, time.getMonthFirstDaysOffJ2000()) != 29 {
		t.Error(time, "getMonthDaysNum error:", sum, "But Expect:", 29)
	}
}

func TestGetMonthFirstDayInfo(t *testing.T) {
	var time = Time{2022, 7, 14, 0, 0, 0}
	var _, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	var rNumberDaysInMonth, rFirstWeekday = 31, 5
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(time, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	time = Time{2022, 2, 14, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 28, 2
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(time, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	time = Time{1582, 10, 6, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 21, 1
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(time, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	time = Time{1582, 9, 30, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 30, 6
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(time, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	time = Time{1582, 1, 1, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 31, 1
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(time, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	time = Time{-1000, 2, 1, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 29, 3
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(time, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

	time = Time{9999, 12, 31, 0, 0, 0}
	_, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	rNumberDaysInMonth, rFirstWeekday = 31, 3
	if numberDaysInMonth != rNumberDaysInMonth || firstWeekday != rFirstWeekday {
		t.Error(time, "error: (", numberDaysInMonth, firstWeekday, "), But Expect:(", rNumberDaysInMonth, ",", rFirstWeekday, ")")
	}

}

// 获取当前时间
func TestTimeNow(t *testing.T) {
	var time = TimeNow()
	t.Log(time)
}
