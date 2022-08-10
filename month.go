package almanac

import (
	"math"
	"strconv"
	"strings"
)

// Month 月份对象
type Month struct {
	// 公历年
	Year int

	// 公历月
	Month int

	// 本月第一天 2000.0 起算的儒略日，北京时间 12:00
	firstDayJD int

	// 月首日期
	FirstDayTime Time

	// 月首是星期几（周日：0...）
	FirstWeekday int

	// 本月的总周数
	WeekCount int

	// 本日所在月的天数（公历）
	DaysCount int

	// 所属公历年对应的农历干支纪年
	ChineseEraDay int

	// 干支纪年-天干的索引
	ChineseEraGan int

	// 干支纪年-地支的索引
	ChineseEraZhi int

	// 生肖
	ChineseZodiac int

	// 朝代
	Dynasty

	// 本月所有的日期
	Days []Day
}

// NewMonth 获取指定的月份
func NewMonth(year, month int) (m *Month) {
	m = new(Month)
	m.Year = year
	m.Month = month
	m.FirstDayTime = Time{year, month, 1, 12, 0, 0}
	var firstDayOff = m.FirstDayTime.getDaysOffJ2000()

	m.firstDayJD = firstDayOff

	// 本月天数
	m.DaysCount = m.FirstDayTime.getMonthDaysNum()
	// 本月月首星期
	m.FirstWeekday = (firstDayOff + j2000 + 1 + 7000000) % 7
	// 本月周数
	m.WeekCount = ((m.FirstWeekday + m.DaysCount - 1) / 7) + 1
	// 所属公历年对应的农历干支纪年
	m.ChineseEraDay = m.Year - 1984 + 12000
	m.ChineseEraGan = m.ChineseEraDay % 10
	m.ChineseEraZhi = m.ChineseEraDay % 12

	// 公历年对应的生肖
	m.ChineseZodiac = m.ChineseEraZhi

	// 年号
	m.Dynasty = *NewDynastyInfo(m.Year)

	var firstDayLunar = NewLunar(m.firstDayJD)
	for i := 0; i < m.DaysCount; i++ {
		var day = newDayWithMonth(m, i, firstDayLunar)
		m.Days = append(m.Days, *day)
	}

	// 月相节气计算
	m.calcYXJQ()

	return
}

// 月相和节气的处理
func (m *Month) calcYXJQ() {
	var Bd0, Bdn, D, xn int
	var d, jd2 float64

	Bd0 = m.firstDayJD
	Bdn = m.DaysCount
	jd2 = float64(Bd0) + dtT(float64(Bd0)) - 8.0/24
	//月相查找
	var w = MsALon(jd2/36525, 10, 3)
	w = math.Floor((w-0.78)/math.Pi*2) * math.Pi / 2

	for {
		d = soAccurate(w)
		D = floorInt(d + 0.5)
		xn = floorInt(w/pi2*4+4000000.01) % 4
		w += pi2 / 4
		if D >= Bd0+Bdn {
			break
		}
		if D < Bd0 {
			continue
		}
		var l = &(m.Days[D-Bd0].Lunar)
		l.phasesOfMoon = yxmc[xn] //取得月相名称
		l.phasesOfMoonJD = d
		l.phasesOfMoonTime = timeFromJD(d + float64(j2000))
		m.Days[D-Bd0].Lunar = *l
		if D+5 >= Bd0+Bdn {
			break
		}
	}

	//节气查找
	w = sALon(jd2/36525, 3)
	w = math.Floor((w-0.13)/pi2*24) * pi2 / 24
	for {
		d = qiAccurate(w)
		D = floorInt(d + 0.5)
		xn = floorInt(w/pi2*24+24000006.01) % 24
		w += pi2 / 24
		if D >= Bd0+Bdn {
			break
		}
		if D < Bd0 {
			continue
		}
		var l = &(m.Days[D-Bd0].Lunar)
		l.solarTerm = jqmc[xn] //取得节气名称
		l.solarTermJD = d
		l.solarTermTime = timeFromJD(d + float64(j2000))
		m.Days[D-Bd0].Lunar = *l
		if D+12 >= Bd0+Bdn {
			break
		}
	}
}

// FormatCal 格式化输出公历的月信息，可以作为日历使用
func (m Month) FormatCal() string {
	var sb = strings.Builder{}
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat(" ", 8))
	sb.WriteString(strconv.Itoa(m.Year))
	sb.WriteString("年  ")
	if m.Month < 10 {
		sb.WriteString(" ")
	}
	sb.WriteString(strconv.Itoa(m.Month))
	sb.WriteString("月\n")
	sb.WriteString("Sun Mon Tue Wed Thu Fri Sat\n")
	sb.WriteString(strings.Repeat(" ", m.FirstWeekday*4))

	var weekNum = 0
	for i, day := range m.Days {
		if weekNum != day.WeekIndexInMonth {
			sb.WriteString("\n")
			weekNum = day.WeekIndexInMonth
		}
		if day.Time.day < 10 {
			sb.WriteString("  ")
		} else {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.Itoa(day.Time.day))
		if day.Week != 6 && i < len(m.Days)-1 {
			sb.WriteString(" ")
		}
	}
	//sb.WriteString("\n")
	return sb.String()
}
