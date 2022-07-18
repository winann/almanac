package almanac

import (
	"testing"
)

func TestNewDay(t *testing.T) {
	//var time = Time{1582, 10, 5, 16, 0, 0}
	var time = Time{2022, 7, 18, 16, 0, 0}
	var day = NewDay(time)
	t.Log(day)
	t.Log(day.GetGanName())
	t.Log(day.GetZhiName())
	t.Log(day.GetChineseZodiacName())
}
