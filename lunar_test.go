package almanac

import (
	"testing"
)

func TestNewLunar(t *testing.T) {
	var time = Time{-1000, 8, 4, 12, 0, 0}
	var jd = time.getDaysOffJ2000()
	var lunar = NewLunar(jd)
	if lunar.LeapStr != "闰" || lunar.Lmc != "六" || lunar.Ldc != "初一" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{-1000, 12, 30, 12, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "十一" || lunar.Ldc != "初一" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{-700, 12, 13, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "十二" || lunar.Ldc != "三十" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{-500, 12, 22, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "正" || lunar.Ldc != "初一" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{-100, 12, 9, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "十" || lunar.Ldc != "廿九" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{-1, 12, 25, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "十一" || lunar.Ldc != "廿九" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{1582, 10, 22, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "九" || lunar.Ldc != "廿六" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{2022, 8, 5, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "七" || lunar.Ldc != "初八" {
		t.Error(time, "lunar error:", lunar)
	}

	time = Time{9999, 12, 31, 14, 0, 0}
	jd = time.getDaysOffJ2000()
	lunar = NewLunar(jd)
	if lunar.Lmc != "十二" || lunar.Ldc != "初二" {
		t.Error(time, "lunar error:", lunar)
	}

}

// 测试节日
func TestCalcLunarEvents(t *testing.T) {
	var time = Time{-1000, 12, 30, 14, 0, 0}
	var day = NewDay(time)
	if day.lunar.events.important[0] != "冬至" || day.lunar.events.important[1] != "一九" {
		t.Error(time, "events error:", day.lunar.events)
	}

	time = Time{-700, 12, 13, 14, 0, 0}
	day = NewDay(time)
	if len(day.lunar.events.festival) != 1 || day.lunar.events.festival[0] != "除夕" {
		t.Error(time, "events error:", day.lunar.events)
	}

	time = Time{-500, 12, 11, 14, 0, 0}
	day = NewDay(time)
	if len(day.lunar.events.important) != 1 || day.lunar.events.important[0] != "大雪" {
		t.Error(time, "events error:", day.lunar.events)
	}

	time = Time{-100, 12, 9, 14, 0, 0}
	day = NewDay(time)
	if len(day.lunar.events.important) != 1 || day.lunar.events.important[0] != "大雪" {
		t.Error(time, "events error:", day.lunar.events)
	}

	time = Time{-1, 12, 25, 14, 0, 0}
	day = NewDay(time)
	if len(day.lunar.events.important) != 2 || day.lunar.events.important[0] != "冬至" {
		t.Error(time, "events error:", day.lunar.events)
	}
	time = Time{1582, 10, 22, 14, 0, 0}
	day = NewDay(time)
	if len(day.lunar.events.important) != 1 || day.lunar.events.important[0] != "霜降" {
		t.Error(time, "events error:", day.lunar.events)
	}
	time = Time{2022, 7, 26, 14, 0, 0}
	day = NewDay(time)
	if len(day.lunar.events.important) != 1 || day.lunar.events.important[0] != "中伏" {
		t.Error(time, "events error:", day.lunar.events)
	}
	time = Time{9999, 12, 31, 14, 0, 0}
	day = NewDay(time)
	if len(day.lunar.events.important) != 1 || day.lunar.events.important[0] != "小寒" {
		t.Error(time, "events error:", day.lunar.events)
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
