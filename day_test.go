package almanac

import (
	"testing"
)

func TestNewDay(t *testing.T) {
	var time = NewTime(-1000, 12, 30, 14, 0, 0)
	var day = NewDay(time)
	if day.MonthFirstWeek != 6 ||
		day.MonthDaysCount != 31 ||
		day.GetZodiacName() != "摩羯座" ||
		day.GetChineseZodiacName() != "龙" ||
		day.GetGanName() != "庚" ||
		day.GetZhiName() != "辰" ||
		day.Mark != "康王" {
		t.Error(day)
	}

	time = NewTime(1582, 10, 4, 14, 0, 0)
	day = NewDay(time)
	if day.MonthFirstWeek != 1 ||
		day.MonthDaysCount != 21 ||
		day.GetZodiacName() != "天秤座" ||
		day.GetChineseZodiacName() != "马" ||
		day.GetGanName() != "壬" ||
		day.GetZhiName() != "午" ||
		day.Mark != "神宗" {
		t.Error(day)
	}

	time = NewTime(1582, 10, 22, 14, 0, 0)
	day = NewDay(time)
	if day.MonthFirstWeek != 1 ||
		day.MonthDaysCount != 21 ||
		day.GetZodiacName() != "天蝎座" ||
		day.GetChineseZodiacName() != "马" ||
		day.GetGanName() != "壬" ||
		day.GetZhiName() != "午" ||
		day.Mark != "神宗" {
		t.Error(day)
	}

	time = NewTime(2022, 8, 5, 14, 0, 0)
	day = NewDay(time)
	if day.MonthFirstWeek != 1 ||
		day.MonthDaysCount != 31 ||
		day.GetZodiacName() != "狮子座" ||
		day.GetChineseZodiacName() != "虎" ||
		day.GetGanName() != "壬" ||
		day.GetZhiName() != "寅" ||
		day.Mark != "中国" {
		t.Error(day)
	}

	time = NewTime(9999, 12, 31, 14, 0, 0)
	day = NewDay(time)
	if day.MonthFirstWeek != 3 ||
		day.MonthDaysCount != 31 ||
		day.GetZodiacName() != "摩羯座" ||
		day.GetChineseZodiacName() != "猪" ||
		day.GetGanName() != "己" ||
		day.GetZhiName() != "亥" ||
		day.Mark != "中国" {
		t.Error(day)
	}

}

// 测试回历
func TestCalcHijri(t *testing.T) {
	var time = NewTime(-1000, 12, 30, 14, 0, 0)
	var day = NewDay(time)
	var e = Hijri{-1671, 8, 27}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = NewTime(1582, 10, 4, 14, 0, 0)
	day = NewDay(time)
	e = Hijri{990, 9, 16}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = NewTime(1582, 10, 22, 14, 0, 0)
	day = NewDay(time)
	e = Hijri{990, 9, 24}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = NewTime(2022, 7, 29, 14, 0, 0)
	day = NewDay(time)
	e = Hijri{1443, 12, 29}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = NewTime(9999, 12, 31, 14, 0, 0)
	day = NewDay(time)
	e = Hijri{9666, 4, 2}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}
}

func TestWeekIndex(t *testing.T) {
	var day = NewDay(NewTime(2022, 1, 1, 0, 0, 0))
	if day.WeekIndexInYear != 0 || day.WeekIndexInMonth != 0 {
		t.Error(day.Time, " Week Index error:")
	}

	day = NewDay(NewTime(2022, 8, 10, 0, 0, 0))
	if day.WeekIndexInYear != 32 || day.WeekIndexInMonth != 1 {
		t.Error(day.Time, " Week Index error:")
	}

	day = NewDay(NewTime(2022, 12, 31, 0, 0, 0))
	if day.WeekIndexInYear != 52 || day.WeekIndexInMonth != 4 {
		t.Error(day.Time, " Week Index error:")
	}
}
