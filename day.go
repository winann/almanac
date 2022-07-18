package almanac

// Day 为某一天的对象
type Day struct {

	// 2000.0 起算的儒略日，北京时间 12:00
	jd int

	// 公历月内序数
	indexInMonth int

	// 时间
	time Time

	// 所在月首是星期几（周日：0...）
	monthFirstWeekday int

	// 当前是星期几（周日：0...）
	weekday int

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
}

// NewDay 通过日期获取一天的数据
func NewDay(time Time) (d Day) {
	d = *new(Day)

	// 公历日名称
	realTime := timeFromJD(time.toJulianDay())
	// 处理 1582 年 10 月等特殊情况
	time.day = realTime.day

	d.time = time

	time.hour = 12

	// firstDayOff 首月的偏移，numberDaysInMonth 所在月总天数，firstWeekday 所在月第一天的星期
	var firstDayOff, numberDaysInMonth, firstWeekday = time.getMonthFirstDayInfo()
	d.jd = time.getDaysOffJ2000()

	// 本日所在月的序号
	d.indexInMonth = d.jd - firstDayOff
	// 本日所在月的总天数
	d.monthDaysCount = numberDaysInMonth
	// 本日所在月第一天的星期
	d.monthFirstWeekday = firstWeekday
	// 本日的星期
	d.weekday = (firstWeekday + d.indexInMonth) % 7
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
	return
}

// newDayWithMonth 通过传入的月信息和当前第几天获取日对象
func newDayWithMonth(m Month, i int) (d Day) {
	d = *new(Day)
	d.jd = m.firstDayJD + i

	// 真实日期
	d.time = timeFromJD(JulianDay(d.jd + J2000))

	// 本日所在月的序号
	d.indexInMonth = i

	// 本日所在月的总天数
	d.monthDaysCount = m.daysCount
	// 本日所在月第一天的星期
	d.monthFirstWeekday = m.firstWeekday
	// 本日的星期
	d.weekday = (m.firstWeekday + d.indexInMonth) % 7
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
	return
}

// GetGanName 获得天干的名称
func (d Day) GetGanName() (ganName string) {
	ganName = gan[d.chineseEraGan]
	return
}

// GetZhiName 获得地支的名称
func (d Day) GetZhiName() (zhiName string) {
	zhiName = zhi[d.chineseEraZhi]
	return
}

// GetChineseZodiacName 获取生肖名称
func (d Day) GetChineseZodiacName() (zhiName string) {
	zhiName = chineseZodiac[d.chineseEraZhi]
	return
}
