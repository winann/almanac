package almanac

import (
	"testing"
)

func TestNewDay(t *testing.T) {
	var time = Time{-1000, 12, 30, 14, 0, 0}
	var day = NewDay(time)
	if day.monthFirstWeek != 6 ||
		day.monthDaysCount != 31 ||
		day.GetZodiacName() != "摩羯座" ||
		day.GetChineseZodiacName() != "龙" ||
		day.GetGanName() != "庚" ||
		day.GetZhiName() != "辰" ||
		day.mark != "康王" {
		t.Error(day)
	}

	time = Time{1582, 10, 4, 14, 0, 0}
	day = NewDay(time)
	if day.monthFirstWeek != 1 ||
		day.monthDaysCount != 21 ||
		day.GetZodiacName() != "天秤座" ||
		day.GetChineseZodiacName() != "马" ||
		day.GetGanName() != "壬" ||
		day.GetZhiName() != "午" ||
		day.mark != "神宗" {
		t.Error(day)
	}

	time = Time{1582, 10, 22, 14, 0, 0}
	day = NewDay(time)
	if day.monthFirstWeek != 1 ||
		day.monthDaysCount != 21 ||
		day.GetZodiacName() != "天蝎座" ||
		day.GetChineseZodiacName() != "马" ||
		day.GetGanName() != "壬" ||
		day.GetZhiName() != "午" ||
		day.mark != "神宗" {
		t.Error(day)
	}

	time = Time{2022, 8, 5, 14, 0, 0}
	day = NewDay(time)
	if day.monthFirstWeek != 1 ||
		day.monthDaysCount != 31 ||
		day.GetZodiacName() != "狮子座" ||
		day.GetChineseZodiacName() != "虎" ||
		day.GetGanName() != "壬" ||
		day.GetZhiName() != "寅" ||
		day.mark != "中国" {
		t.Error(day)
	}

	time = Time{9999, 12, 31, 14, 0, 0}
	day = NewDay(time)
	if day.monthFirstWeek != 3 ||
		day.monthDaysCount != 31 ||
		day.GetZodiacName() != "摩羯座" ||
		day.GetChineseZodiacName() != "猪" ||
		day.GetGanName() != "己" ||
		day.GetZhiName() != "亥" ||
		day.mark != "中国" {
		t.Error(day)
	}

}

// 测试回历
func TestCalcHijri(t *testing.T) {
	var time = Time{-1000, 12, 30, 14, 0, 0}
	var day = NewDay(time)
	var e = Hijri{-1671, 8, 27}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = Time{1582, 10, 4, 14, 0, 0}
	day = NewDay(time)
	e = Hijri{990, 9, 16}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = Time{1582, 10, 22, 14, 0, 0}
	day = NewDay(time)
	e = Hijri{990, 9, 24}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = Time{2022, 7, 29, 14, 0, 0}
	day = NewDay(time)
	e = Hijri{1443, 12, 29}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

	time = Time{9999, 12, 31, 14, 0, 0}
	day = NewDay(time)
	e = Hijri{9666, 4, 2}
	if day.Hijri != e {
		t.Error(time, "calcHijri error:", day.Hijri, ", But expect:", e)
	}

}
