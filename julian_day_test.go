package almanac

import (
	"testing"
)

func TestJulianDay(t *testing.T) {
	var jd = julianDay(2000, 1, 1.5)
	if !isEqual(jd, 2451545.0) {
		t.Error("julianDay(2000, 1, 1.5), 2451545.0)", "Error:", jd)
	}
	jd = julianDay(1987, 1, 27.0)
	if !isEqual(jd, 2446822.5) {
		t.Error("julianDay(1987, 1, 27.0), 2446822.5)", "Error:", jd)
	}
	jd = julianDay(1987, 6, 19.5)
	if !isEqual(jd, 2446966.0) {
		t.Error("julianDay(1987, 6, 19.5), 2446966.0)", "Error:", jd)
	}
	jd = julianDay(1988, 1, 27.0)
	if !isEqual(jd, 2447187.5) {
		t.Error("julianDay(1988, 1, 27.0), 2447187.5)", "Error:", jd)
	}
	jd = julianDay(1988, 6, 19.5)
	if !isEqual(jd, 2447332.0) {
		t.Error("julianDay(1988, 6, 19.5), 2447332.0)", "Error:", jd)
	}
	jd = julianDay(1900, 1, 1.0)
	if !isEqual(jd, 2415020.5) {
		t.Error("julianDay(1900, 1, 1.0), 2415020.5)", "Error:", jd)
	}
	jd = julianDay(1600, 1, 1.0)
	if !isEqual(jd, 2305447.5) {
		t.Error("julianDay(1600, 1, 1.0), 2305447.5)", "Error:", jd)
	}
	jd = julianDay(1600, 12, 31.0)
	if !isEqual(jd, 2305812.5) {
		t.Error("julianDay(1600, 12, 31.0), 2305812.5)", "Error:", jd)
	}
	jd = julianDay(837, 4, 10.3)
	if !isEqual(jd, 2026871.8) {
		t.Error("julianDay(837, 4, 10.3), 2026871.8)", "Error:", jd)
	}
	jd = julianDay(-1000, 7, 12.5)
	if !isEqual(jd, 1356001.0) {
		t.Error("julianDay(-1000, 7, 12.5), 1356001.0)", "Error:", jd)
	}
	jd = julianDay(-1000, 2, 29.0)
	if !isEqual(jd, 1355866.5) {
		t.Error("julianDay(-1000, 2, 29.0), 1355866.5)", "Error:", jd)
	}
	jd = julianDay(-1001, 8, 17.9)
	if !isEqual(jd, 1355671.4) {
		t.Error("julianDay(-1001, 8, 17.9), 1355671.4)", "Error:", jd)
	}
	jd = julianDay(-4712, 1, 1.5)
	if !isEqual(jd, 0.0) {
		t.Error("julianDay(-4712, 1, 1.5), 0.0)", "Error:", jd)
	}
}
func TestTimeFromJD(t *testing.T) {
	var time = Time{2000, 1, 1, 12, 0, 0}
	var r = timeFromJD(2451545.0)
	if r != time {
		t.Error("timeFromJD(2451545.0)", "Error:", r, ",But expect:", time)
	}
	time = Time{1987, 1, 27, 0, 0, 0}
	r = timeFromJD(2446822.5)
	if r != time {
		t.Error("timeFromJD(2446822.5)", "Error:", r, ",But expect:", time)
	}
	time = Time{1987, 6, 19, 12, 0, 0}
	r = timeFromJD(2446966.0)
	if r != time {
		t.Error("timeFromJD(2446966.0)", "Error:", r, ",But expect:", time)
	}
	time = Time{1988, 1, 27, 0, 0, 0}
	r = timeFromJD(2447187.5)
	if r != time {
		t.Error("timeFromJD(2447187.5)", "Error:", r, ",But expect:", time)
	}
	time = Time{1988, 6, 19, 12, 0, 0}
	r = timeFromJD(2447332.0)
	if r != time {
		t.Error("timeFromJD(2447332.0)", "Error:", r, ",But expect:", time)
	}
	time = Time{1900, 1, 1, 0, 0, 0}
	r = timeFromJD(2415020.5)
	if r != time {
		t.Error("timeFromJD(2415020.5)", "Error:", r, ",But expect:", time)
	}
	time = Time{1600, 1, 1, 0, 0, 0}
	r = timeFromJD(2305447.5)
	if r != time {
		t.Error("timeFromJD(2305447.5)", "Error:", r, ",But expect:", time)
	}
	time = Time{1600, 12, 31, 0, 0, 0}
	r = timeFromJD(2305812.5)
	if r != time {
		t.Error("timeFromJD(2305812.5)", "Error:", r, ",But expect:", time)
	}
	time = Time{837, 4, 10, 7, 12, 0}
	r = timeFromJD(2026871.8)
	if r != time {
		t.Error("timeFromJD(2026871.8)", "Error:", r, ",But expect:", time)
	}
	time = Time{-1000, 7, 12, 12, 0, 0}
	r = timeFromJD(1356001.0)
	if r != time {
		t.Error("timeFromJD(1356001.0)", "Error:", r, ",But expect:", time)
	}
	time = Time{-1000, 2, 29, 0, 0, 0}
	r = timeFromJD(1355866.5)
	if r != time {
		t.Error("timeFromJD(1355866.5)", "Error:", r, ",But expect:", time)
	}
	time = Time{-1001, 8, 17, 21, 36, 0}
	r = timeFromJD(1355671.4)
	if r != time {
		t.Error("timeFromJD(1355671.4)", "Error:", r, ",But expect:", time)
	}
	time = Time{-4712, 1, 1, 12, 0, 0}
	r = timeFromJD(0.0)
	if r != time {
		t.Error("timeFromJD(0.0)", "Error:", r, ",But expect:", time)
	}
}
