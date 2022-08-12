package almanac

import (
	"testing"
)

func TestNewMonth(t *testing.T) {
	var month = NewMonth(-1000, 2)
	var e = `
        -1000年   2月
Sun Mon Tue Wed Thu Fri Sat
              1   2   3   4
  5   6   7   8   9  10  11
 12  13  14  15  16  17  18
 19  20  21  22  23  24  25
 26  27  28  29`

	var r = month.FormatCal()
	if e != r {
		t.Error(month, "\nmonth.FormatCal error:", r, "\nBut Expect:", e)
	}

	month = NewMonth(-1000, 8)
	r = month.FormatCal()
	e = `
        -1000年   8月
Sun Mon Tue Wed Thu Fri Sat
              1   2   3   4
  5   6   7   8   9  10  11
 12  13  14  15  16  17  18
 19  20  21  22  23  24  25
 26  27  28  29  30  31`

	if e != r {
		t.Error(month, "\nmonth.FormatCal error:", r, "\nBut Expect:", e)
	}
	//
	//for _, day := range Month.Days {
	//	t.Log(day.Lunar.LeapStr, day.Lunar.MonthName, day.Lunar.DayName)
	//}

	month = NewMonth(-1000, 9)
	r = month.FormatCal()
	e = `
        -1000年   9月
Sun Mon Tue Wed Thu Fri Sat
                          1
  2   3   4   5   6   7   8
  9  10  11  12  13  14  15
 16  17  18  19  20  21  22
 23  24  25  26  27  28  29
 30`

	if e != r {
		t.Error(month, "\nmonth.FormatCal error:", r, "\nBut Expect:", e)
	}

	//for _, day := range Month.Days {
	//	t.Log(day.Lunar.MonthName, day.Lunar.DayName)
	//}
}

func TestYXJQ(t *testing.T) {
	var m = NewMonth(2022, 8)
	if m.Days[4].PhasesOfMoon != "上弦" {
		t.Error("calYXJQ error:", m.Days[5])
	}
}
