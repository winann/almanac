package almanac

import (
	"testing"
)

func TestNewLunar(t *testing.T) {
	var time = Time{-1000, 8, 4, 12, 0, 0}
	var jd = time.getDaysOffJ2000()
	var lunar = NewLunar(jd)
	if lunar.LeapStr != "闰" || lunar.MonthName != "六" || lunar.DayName != "初一" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{-1000, 12, 30, 12, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "十一" || lunar.DayName != "初一" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{-700, 12, 13, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "十二" || lunar.DayName != "三十" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{-500, 12, 22, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "正" || lunar.DayName != "初一" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{-100, 12, 9, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "十" || lunar.DayName != "廿九" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{-1, 12, 25, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "十一" || lunar.DayName != "廿九" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{1582, 10, 22, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "九" || lunar.DayName != "廿六" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{2022, 8, 5, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "七" || lunar.DayName != "初八" {
		t.Error(time, "Lunar error:", lunar)
	}

	time = Time{9999, 12, 31, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.MonthName != "十二" || lunar.DayName != "初二" {
		t.Error(time, "Lunar error:", lunar)
	}

}

// 测试节日
func TestCalcLunarEvents(t *testing.T) {
	var time = Time{-1000, 12, 30, 14, 0, 0}
	var day = NewDay(time)
	if day.Lunar.Events.important[0] != "冬至" || day.Lunar.Events.important[1] != "一九" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}

	time = Time{-700, 12, 13, 14, 0, 0}
	day = NewDay(time)
	if len(day.Lunar.Events.festival) != 1 || day.Lunar.Events.festival[0] != "除夕" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}

	time = Time{-500, 12, 11, 14, 0, 0}
	day = NewDay(time)
	if len(day.Lunar.Events.important) != 1 || day.Lunar.Events.important[0] != "大雪" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}

	time = Time{-100, 12, 9, 14, 0, 0}
	day = NewDay(time)
	if len(day.Lunar.Events.important) != 1 || day.Lunar.Events.important[0] != "大雪" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}

	time = Time{-1, 12, 25, 14, 0, 0}
	day = NewDay(time)
	if len(day.Lunar.Events.important) != 2 || day.Lunar.Events.important[0] != "冬至" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}
	time = Time{1582, 10, 22, 14, 0, 0}
	day = NewDay(time)
	if len(day.Lunar.Events.important) != 1 || day.Lunar.Events.important[0] != "霜降" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}
	time = Time{2022, 7, 26, 14, 0, 0}
	day = NewDay(time)
	if len(day.Lunar.Events.important) != 1 || day.Lunar.Events.important[0] != "中伏" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}
	time = Time{9999, 12, 31, 14, 0, 0}
	day = NewDay(time)
	if len(day.Lunar.Events.important) != 1 || day.Lunar.Events.important[0] != "小寒" {
		t.Error(time, "Events error:", day.Lunar.Events)
	}
}

func TestGetShuoQiDay(t *testing.T) {
	var r = getShuoQiDay(-1095756.53, false)
	var e = -1095756
	if r != e {
		t.Error("getShuoQiDay error:", r, ", But expect:", e)
	}
}

func TestQiHigh(t *testing.T) {
	var r = qiHigh(-18845.105331946175)
	var e = -1.0957534414463292e+06
	if r != e {
		t.Error("qiHigh error:", r, ", But expect:", e)
	}
}

func TestCalcYXJQ(t *testing.T) {
	var l = NewLunar(8259)
	l.calcYXJQ()
	var time = Time{2022, 8, 12, 9, 35, 43}
	if l.phasesOfMoon != "望" && l.phasesOfMoonTime != time {
		t.Error(time, "calcYXJQ error：", l)
	}

	l = NewLunar(2921924)
	l.calcYXJQ()
	time = Time{9999, 12, 16, 8, 26, 11}
	if l.solarTerm != "冬至" && l.solarTermTime != time {
		t.Error(time, "calcYXJQ error：", l)
	}

	l = NewLunar(-10)
	l.calcYXJQ()
	time = Time{1999, 12, 22, 15, 43, 48}
	if l.solarTerm != "冬至" && l.solarTermTime != time {
		t.Error(time, "calcYXJQ error：", l)
	}

	l = NewLunar(-698128)
	l.calcYXJQ()
	time = Time{1999, 12, 22, 16, 45, 56}
	if l.phasesOfMoon != "朔" && l.solarTermTime != time {
		t.Error(time, "calcYXJQ error：", l)
	}

	l = NewLunar(-1095521)
	l.calcYXJQ()
	time = Time{-1000, 8, 4, 12, 6, 18}
	if l.phasesOfMoon != "朔" && l.solarTermTime != time {
		t.Error(time, "calcYXJQ error：", l)
	}
}
