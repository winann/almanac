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
	indexInMonth int

	// 时间
	Time

	// 所在月首是星期几（周日：0...）
	monthFirstWeek int

	// 当前是星期几（周日：0...）
	week int

	// 本日所在月的周序号
	weekNumInMonth int

	// 本月的总周数
	weekCountInMonth int

	// 本日所在月的天数（公历）
	monthDaysCount int

	// 所属公历年对应的农历干支纪年
	chineseEraDay int

	// 干支纪年-天干的索引
	chineseEraGan int

	// 干支纪年-地支的索引
	chineseEraZhi int

	// 生肖
	chineseZodiac int

	// 星座
	zodiac int

	// 朝代
	dynasty

	// 农历信息
	lunar

	// 回历信息
	Hijri

	// 节假日信息
	events event
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

	// firstDayOff 月首日的偏移，numberDaysInMonth 所在月总天数，firstWeekday 所在月第一天的星期
	var firstDayOff, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	d.jd = time.getDaysOffJ2000()

	// 本日所在月的序号
	d.indexInMonth = d.jd - firstDayOff
	// 本日所在月的总天数
	d.monthDaysCount = numberDaysInMonth
	// 本日所在月第一天的星期
	d.monthFirstWeek = firstWeekday
	// 本日的星期
	d.week = (firstWeekday + d.indexInMonth) % 7
	// 本日所在月的周序号
	d.weekNumInMonth = (firstWeekday + d.indexInMonth) / 7
	// 所在月总周数
	d.weekCountInMonth = ((firstWeekday + numberDaysInMonth - 1) / 7) + 1

	// 所属公历年对应的农历干支纪年
	d.chineseEraDay = time.year - 1984 + 12000
	d.chineseEraGan = d.chineseEraDay % 10
	d.chineseEraZhi = d.chineseEraDay % 12

	// 公历年对应的生肖
	d.chineseZodiac = d.chineseEraZhi

	// 年号
	d.dynasty = *NewDynastyInfo(time.year)

	// 农历信息
	if firstDayOff < d.ZQ[0] || firstDayOff >= d.ZQ[24] {
		d.lunar = *(NewLunar(d.jd))
	}

	//星座
	var mk = int(float64(d.jd-d.ZQ[0]-15) / 30.43685)
	//星座所在月的序数,(如果j=13,ob.d0不会超过第14号中气)
	if mk < 11 && d.jd >= d.ZQ[2*mk+2] {
		mk++
	}
	d.zodiac = mk

	// 回历
	d.calcHijri()

	// 公历节日
	d.calcEvents()
	return
}

// newDayWithMonth 通过传入的月信息和当前第几天获取日对象
func newDayWithMonth(m Month, i int, l *lunar) (d *Day) {
	d = new(Day)
	d.jd = m.firstDayJD + i
	d.lunar = *l
	// 真实日期
	d.Time = timeFromJD(JulianDay(d.jd + J2000))

	// 本日所在月的序号
	d.indexInMonth = i

	// 本日所在月的总天数
	d.monthDaysCount = m.daysCount
	// 本日所在月第一天的星期
	d.monthFirstWeek = m.firstWeekday
	// 本日的星期
	d.week = (m.firstWeekday + d.indexInMonth) % 7
	// 本日所在月的周序号
	d.weekNumInMonth = (m.firstWeekday + d.indexInMonth) / 7
	// 所在月总周数
	d.weekCountInMonth = m.weekCount

	// 所属公历年对应的农历干支纪年
	d.chineseEraDay = m.chineseEraDay
	d.chineseEraGan = m.chineseEraGan
	d.chineseEraZhi = m.chineseEraZhi

	// 公历年对应的生肖
	d.chineseZodiac = d.chineseEraZhi

	// 年号
	d.dynasty = *NewDynastyInfo(m.year)

	// 农历信息
	l.updateLunar(d.jd)
	d.lunar = *l

	//星座
	var mk = int(float64(d.jd-d.ZQ[0]-15) / 30.43685)
	//星座所在月的序数,(如果j=13,ob.d0不会超过第14号中气)
	if mk < 11 && d.jd >= d.ZQ[2*mk+2] {
		mk++
	}
	d.zodiac = mk

	// 回历
	d.calcHijri()

	// 公历节日
	d.calcEvents()

	return
}

// GetGanName 获得天干的名称
func (day Day) GetGanName() (ganName string) {
	ganName = gan[day.chineseEraGan]
	return
}

// GetZhiName 获得地支的名称
func (day Day) GetZhiName() (zhiName string) {
	zhiName = zhi[day.chineseEraZhi]
	return
}

// GetChineseZodiacName 获取生肖名称
func (day Day) GetChineseZodiacName() (chineseZodiacStr string) {
	chineseZodiacStr = chineseZodiac[day.chineseEraZhi]
	return
}

// GetZodiacName 获取星座名称
func (day Day) GetZodiacName() (zodiacStr string) {
	zodiacStr = zodiac[(day.zodiac+12)%12] + "座"
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
	day.events.isWeekend = day.week == 0 || day.week == 6 // 是否是星期日或星期六
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
			if day.year < start || day.year > end {
				continue
			}
			s = s[10:]
		} else {
			if day.year < 1850 {
				continue
			}
			s = s[1:]
		}
		if t == "#" {
			day.events.festival = append(day.events.festival, s)
		}
		if t == "I" {
			day.events.important = append(day.events.important, s)
		}
		if t == "." {
			day.events.other = append(day.events.other, s)
		}
	}

	//按周查找
	var w = day.weekNumInMonth
	if day.week >= day.monthFirstWeek {
		w += 1
	}
	var w2 = w
	if day.weekNumInMonth == day.weekCountInMonth-1 {
		w2 = 5
	}
	var wStr = m0 + strconv.Itoa(w) + strconv.Itoa(day.week) //d日在本月的第几个星期某
	var w2Str = m0 + strconv.Itoa(w2) + strconv.Itoa(day.week)

	for i := 0; i < len(wFtv); i++ {
		s = wFtv[i]
		var s2 = s[0:4]
		if s2 != wStr && s2 != w2Str {
			continue
		}
		t = s[4:5]
		s = s[5:]
		if t == "#" {
			day.events.festival = append(day.events.festival, s)
		}
		if t == "I" {
			day.events.important = append(day.events.important, s)
		}
		if t == "." {
			day.events.other = append(day.events.other, s)
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
