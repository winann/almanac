package almanac

import (
	"math"
	"strconv"
)

// Hijri 为回历
type Hijri struct {
	HYear  int
	HMonth int
	HDay   int
}

// Day 为某一天的对象
type Day struct {

	// 2000.0 起算的儒略日，北京时间 12:00
	jd int

	// 公历月内序数
	IndexInMonth int

	// 时间
	Time

	// 所在月首是星期几（周日：0...）
	MonthFirstWeek int

	// 当前是星期几（周日：0...）
	Week int

	// 当前所在年的周序号(第一周:0...)
	WeekIndexInYear int

	// 本日所在月的周序号(第一周:0...)
	WeekIndexInMonth int

	// 本月的总周数
	WeekCountInMonth int

	// 本日所在月的天数（公历）
	MonthDaysCount int

	// 所属公历年对应的农历干支纪年
	ChineseEraDay int

	// 干支纪年-天干的索引
	ChineseEraGan int

	// 干支纪年-地支的索引
	ChineseEraZhi int

	// 生肖
	ChineseZodiac int

	// 星座
	Zodiac int

	// 朝代
	Dynasty

	// 农历信息
	Lunar

	// 回历信息
	Hijri

	// 节假日信息
	Events Event
}

// NewDay 通过日期获取一天的数据
func NewDay(time Time) (d *Day) {
	d = new(Day)

	// 公历日名称
	realTime := timeFromJD(time.toJulianDay())
	// 处理 1582 年 10 月等特殊情况
	time.day = realTime.day

	d.Time = time

	time.hour = 12

	// firstDayOff 月首日的偏移，numberDaysInMonth 所在月总天数，FirstWeekday 所在月第一天的星期
	var firstDayOff, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	d.jd = time.getDaysOffJ2000()

	// 本日所在月的序号
	d.IndexInMonth = d.jd - firstDayOff
	// 本日所在月的总天数
	d.MonthDaysCount = numberDaysInMonth
	// 本日所在月第一天的星期
	d.MonthFirstWeek = firstWeekday
	// 本日的星期
	d.Week = (firstWeekday + d.IndexInMonth) % 7
	// 本日所在月的周序号
	d.WeekIndexInMonth = (firstWeekday + d.IndexInMonth) / 7
	// 所在月总周数
	d.WeekCountInMonth = ((firstWeekday + numberDaysInMonth - 1) / 7) + 1
	// 获取今天在本年的第几周
	d.getWeekIndexInYear()

	// 所属公历年对应的农历干支纪年
	d.ChineseEraDay = time.year - 1984 + 12000
	d.ChineseEraGan = d.ChineseEraDay % 10
	d.ChineseEraZhi = d.ChineseEraDay % 12

	// 公历年对应的生肖
	d.ChineseZodiac = d.ChineseEraZhi

	// 年号
	d.Dynasty = *NewDynastyInfo(time.year)

	// 农历信息
	var l = &(d.Lunar)
	if firstDayOff < d.zq[0] || firstDayOff >= d.zq[24] {
		l = NewLunar(d.jd)
	}

	// 月相和节气
	l.calcYXJQ()

	d.Lunar = *l

	//星座
	var mk = int(float64(d.jd-d.zq[0]-15) / 30.43685)
	//星座所在月的序数,(如果j=13,ob.d0不会超过第14号中气)
	if mk < 11 && d.jd >= d.zq[2*mk+2] {
		mk++
	}
	d.Zodiac = mk

	// 回历
	d.calcHijri()

	// 公历节日
	d.calcEvents()

	return
}

// newDayWithMonth 通过传入的月信息和当前第几天获取日对象
func newDayWithMonth(m *Month, i int, l *Lunar) (d *Day) {
	d = new(Day)
	d.jd = m.firstDayJD + i
	d.Lunar = *l
	// 真实日期
	d.Time = timeFromJD(JulianDay(d.jd + j2000))

	// 本日所在月的序号
	d.IndexInMonth = i

	// 本日所在月的总天数
	d.MonthDaysCount = m.DaysCount
	// 本日所在月第一天的星期
	d.MonthFirstWeek = m.FirstWeekday
	// 本日的星期
	d.Week = (m.FirstWeekday + d.IndexInMonth) % 7
	// 本日所在月的周序号
	d.WeekIndexInMonth = (m.FirstWeekday + d.IndexInMonth) / 7
	// 所在月总周数
	d.WeekCountInMonth = m.WeekCount

	// 所属公历年对应的农历干支纪年
	d.ChineseEraDay = m.ChineseEraDay
	d.ChineseEraGan = m.ChineseEraGan
	d.ChineseEraZhi = m.ChineseEraZhi

	// 公历年对应的生肖
	d.ChineseZodiac = d.ChineseEraZhi

	// 年号
	d.Dynasty = *NewDynastyInfo(m.Year)

	// 农历信息
	l.updateLunar(d.jd)

	d.Lunar = *l

	//星座
	var mk = int(float64(d.jd-d.zq[0]-15) / 30.43685)
	//星座所在月的序数,(如果j=13,ob.d0不会超过第14号中气)
	if mk < 11 && d.jd >= d.zq[2*mk+2] {
		mk++
	}
	d.Zodiac = mk

	// 回历
	d.calcHijri()

	// 公历节日
	d.calcEvents()

	return
}

// GetGanName 获得天干的名称
func (day Day) GetGanName() (ganName string) {
	ganName = gan[day.ChineseEraGan]
	return
}

// GetZhiName 获得地支的名称
func (day Day) GetZhiName() (zhiName string) {
	zhiName = zhi[day.ChineseEraZhi]
	return
}

// GetChineseZodiacName 获取生肖名称
func (day Day) GetChineseZodiacName() (chineseZodiacStr string) {
	chineseZodiacStr = chineseZodiac[day.ChineseEraZhi]
	return
}

// GetZodiacName 获取星座名称
func (day Day) GetZodiacName() (zodiacStr string) {
	zodiacStr = zodiac[(day.Zodiac+12)%12] + "座"
	return
}

// 获取公历节日
func (day *Day) calcEvents() {
	var m0, d0 string
	if day.month < 10 {
		m0 = "0"
	}
	m0 += strconv.Itoa(day.month)

	if day.day < 10 {
		d0 = "0"
	}
	d0 += strconv.Itoa(day.day)
	day.Events.isWeekend = day.Week == 0 || day.Week == 6 // 是否是星期日或星期六
	var s, t string
	//按公历日期查找
	for i := 0; i < len(sFtv[day.month-1]); i++ { //公历节日或纪念日,遍历本月节日表
		s = sFtv[day.month-1][i]
		if s[0:2] != d0 {
			continue
		}
		s = s[2:]
		t = s[0:1]
		if s[5:6] == "-" { //有年限的
			var start, _ = strconv.Atoi(s[1:5])
			var end, _ = strconv.Atoi(s[6:10])
			if day.Time.year < start || day.Time.year > end {
				continue
			}
			s = s[10:]
		} else {
			if day.Time.year < 1850 {
				continue
			}
			s = s[1:]
		}
		if t == "#" {
			day.Events.festival = append(day.Events.festival, s)
		}
		if t == "I" {
			day.Events.important = append(day.Events.important, s)
		}
		if t == "." {
			day.Events.other = append(day.Events.other, s)
		}
	}

	//按周查找
	var w = day.WeekIndexInMonth
	if day.Week >= day.MonthFirstWeek {
		w += 1
	}
	var w2 = w
	if day.WeekIndexInMonth == day.WeekCountInMonth-1 {
		w2 = 5
	}
	var wStr = m0 + strconv.Itoa(w) + strconv.Itoa(day.Week) //d日在本月的第几个星期某
	var w2Str = m0 + strconv.Itoa(w2) + strconv.Itoa(day.Week)

	for i := 0; i < len(wFtv); i++ {
		s = wFtv[i]
		var s2 = s[0:4]
		if s2 != wStr && s2 != w2Str {
			continue
		}
		t = s[4:5]
		s = s[5:]
		if t == "#" {
			day.Events.festival = append(day.Events.festival, s)
		}
		if t == "I" {
			day.Events.important = append(day.Events.important, s)
		}
		if t == "." {
			day.Events.other = append(day.Events.other, s)
		}
	}
}

// 获取回历
func (day *Day) calcHijri() {
	//以下算法使用Excel测试得到,测试时主要关心年临界与月临界
	var z, y, m, d float64
	d = float64(day.jd + 503105)
	z = math.Floor(d / 10631) //10631为一周期(30年)
	d -= z * 10631
	y = math.Floor((float64(d) + 0.50000001) / 354.366) //加0.5的作用是保证闰年正确(一周中的闰年是第2,5,7,10,13,16,18,21,24,26,29年)
	d -= math.Floor(y*354.366 + 0.5000001)
	m = math.Floor((d + 0.11) / 29.51) //分子加0.11,分母加0.01的作用是第354或355天的的月分保持为12月(m=11)
	d -= math.Floor(m*29.5 + 0.5000001)
	day.HYear = int(z*30 + y + 1)
	day.HMonth = int(m + 1)
	day.HDay = int(d + 1)
}

// getYearFirstDayWeek 获取某一年的第一天是周几
func (day *Day) getWeekIndexInYear() {
	var firstDayTime = Time{day.year, 1, 1, 12, 0, 0}
	var firstDayOff = firstDayTime.getDaysOffJ2000()
	var yearFirstDayWeek = (firstDayOff + j2000 + 1 + 7000000) % 7
	day.WeekIndexInYear = (day.jd - firstDayOff + yearFirstDayWeek) / 7
}
