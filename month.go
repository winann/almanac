package almanac

import (
	"math"
	"strconv"
	"strings"
)

// Month 月份对象
type Month struct {
	// 公历年
	year int

	// 公历月
	month int

	// 本月第一天 2000.0 起算的儒略日，北京时间 12:00
	firstDayJD int

	// 月首日期
	firstDayTime Time

	// 月首是星期几（周日：0...）
	firstWeekday int

	// 本月的总周数
	weekCount int

	// 本日所在月的天数（公历）
	daysCount int

	// 所属公历年对应的农历干支纪年
	chineseEraDay int

	// 干支纪年-天干的索引
	chineseEraGan int

	// 干支纪年-地支的索引
	chineseEraZhi int

	// 生肖
	chineseZodiac int

	// 朝代
	dynasty

	// 本月所有的日期
	days []Day
}

// NewMonth 获取指定的月份
func NewMonth(year, month int) (m *Month) {
	m = new(Month)
	m.year = year
	m.month = month
	m.firstDayTime = Time{year, month, 1, 12, 0, 0}
	var firstDayOff = m.firstDayTime.getDaysOffJ2000()

	m.firstDayJD = firstDayOff

	// 本月天数
	m.daysCount = m.firstDayTime.getMonthDaysNum()
	// 本月月首星期
	m.firstWeekday = (firstDayOff + J2000 + 1 + 7000000) % 7
	// 本月周数
	m.weekCount = ((m.firstWeekday + m.daysCount - 1) / 7) + 1
	// 所属公历年对应的农历干支纪年
	m.chineseEraDay = m.year - 1984 + 12000
	m.chineseEraGan = m.chineseEraDay % 10
	m.chineseEraZhi = m.chineseEraDay % 12

	// 公历年对应的生肖
	m.chineseZodiac = m.chineseEraZhi

	// 年号
	m.dynasty = *NewDynastyInfo(m.year)

	var firstDayLunar = NewLunar(m.firstDayJD)
	for i := 0; i < m.daysCount; i++ {
		var day = newDayWithMonth(m, i, firstDayLunar)
		m.days = append(m.days, *day)
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
	Bdn = m.daysCount
	jd2 = float64(Bd0) + dt_T(float64(Bd0)) - 8.0/24
	//月相查找
	var w = MS_aLon(jd2/36525, 10, 3)
	w = math.Floor((w-0.78)/math.Pi*2) * math.Pi / 2

	for {
		d = so_accurate(w)
		D = floorInt(d + 0.5)
		xn = floorInt(w/pi2*4+4000000.01) % 4
		w += pi2 / 4
		if D >= Bd0+Bdn {
			break
		}
		if D < Bd0 {
			continue
		}
		var l = &m.days[D-Bd0]
		l.yxmc = yxmc[xn] //取得月相名称
		l.yxjd = d
		l.yxsj = timeFromJD(d + float64(J2000))

		if D+5 >= Bd0+Bdn {
			break
		}
	}

	//节气查找
	w = S_aLon(jd2/36525, 3)
	w = math.Floor((w-0.13)/pi2*24) * pi2 / 24
	for {
		d = qi_accurate(w)
		D = floorInt(d + 0.5)
		xn = floorInt(w/pi2*24+24000006.01) % 24
		w += pi2 / 24
		if D >= Bd0+Bdn {
			break
		}
		if D < Bd0 {
			continue
		}
		var l = &m.days[D-Bd0]
		l.jqmc = jqmc[xn] //取得节气名称
		l.jqjd = d
		l.jqsj = timeFromJD(d + float64(J2000))

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
	sb.WriteString(strconv.Itoa(m.year))
	sb.WriteString("年  ")
	if m.month < 10 {
		sb.WriteString(" ")
	}
	sb.WriteString(strconv.Itoa(m.month))
	sb.WriteString("月\n")
	sb.WriteString("Sun Mon Tue Wed Thu Fri Sat\n")
	sb.WriteString(strings.Repeat(" ", m.firstWeekday*4))

	var weekNum = 0
	for i, day := range m.days {
		if weekNum != day.weekNumInMonth {
			sb.WriteString("\n")
			weekNum = day.weekNumInMonth
		}
		if day.Time.day < 10 {
			sb.WriteString("  ")
		} else {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.Itoa(day.Time.day))
		if day.week != 6 && i < len(m.days)-1 {
			sb.WriteString(" ")
		}
	}
	//sb.WriteString("\n")
	return sb.String()
}
