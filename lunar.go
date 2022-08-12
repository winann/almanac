package almanac

import (
	"math"
	"strconv"
)

type Lunar struct {
	// 农历纪年(10进制,1984年起算)
	Year int

	// 干支纪年（立春）
	Year2 string

	// 干支纪年（春节/正月）
	Year3 string

	// 黄帝纪年
	Year4 int

	// 月名称
	MonthName string

	// 月大小
	MonthDayCount int

	// 闰月情况（"闰"/""）
	LeapStr string

	// 下个月名称，判断除夕时需要用到
	NextMonthName string

	// 干支纪月
	MonthGanZhi int

	// 干支纪月名称
	MonthGanZhiName string

	// 干支纪日
	DayGanZhiName string

	// 距农历月首的编移量,0对应初一
	IndexInLunarMonth int

	// 农历日名称
	DayName string

	// 节气名称
	SolarTermStr string

	// 月相名称
	PhasesOfMoon string

	// 月相时刻(儒略日)
	phasesOfMoonJD JulianDay

	// 月相时间
	PhasesOfMoonTime Time

	// 定气名称
	SolarTerm string

	// 节气时刻(儒略日)
	solarTermJD JulianDay

	// 节气时间
	SolarTermTime Time

	// 距冬至的天数
	CurDZ int

	// 距夏至的天数
	CurXZ int

	// 距立秋的天数
	CurLQ int

	// 距芒种的天数
	CurMZ int

	// 距小暑的天数
	CurXS int

	// 节日、假期等事件
	Events Event

	jd int // 儒略日
	// 补算二气,确保一年中所有月份的“气”全部被计算在内
	pe1, pe2 int
	// 闰月位置
	leap int
	// 中气表,其中.liqiu 是节气立秋的儒略日,计算三伏时用到
	zq [25]int
	// 合朔表
	hs [15]int
	// 各月大小
	dx [14]int
	// 各月名称
	ymc [14]string
}

// NewLunar 生成农历排序,  jd 为儒略日相对于 j2000 的偏移
// 时间系统全部使用北京时，即使是天象时刻的输出，也是使用北京时
// 如果天象的输出不使用北京时，会造成显示混乱，更严重的是无法与古历比对
// 注意：有 Lunar 对象之后，建议使用 updateLunar 来更新农历，减少计算次数
func NewLunar(jd int) (l *Lunar) {
	l = new(Lunar)
	l.jd = jd
	l.cal()
	l.updateLunar(l.jd)
	return
}

// GetPhasesOfMoonTime 获取月相时间，自己查看 PhasesOfMoon 月相是否存在
// 为 iOS 提供的接口
func (l *Lunar) GetPhasesOfMoonTime() *Time {
	return &l.PhasesOfMoonTime
}

// GetSolarTermTime 获取节气时间，自己查看 SolarTerm 节气是否存在
// 为 iOS 提供的接口
func (l *Lunar) GetSolarTermTime() *Time {
	return &l.SolarTermTime
}

// GetEvents 获取农历节日
// 为 iOS 提供的接口
func (l *Lunar) GetEvents() *Event {
	return &l.Events
}

// 排月序(生成实际年历)
func (l *Lunar) cal() {
	var A, B = &l.zq, &l.hs //中气表,日月合朔表(整日)
	// 该年的气
	var W = math.Floor(float64(l.jd-355+183)/365.2422)*365.2422 + 355 // 355是2000.12冬至,得到较靠近jd的冬至估计值
	if getShuoQiDay(W, true) > l.jd {
		W -= 365.2422
	}
	for i := 0; i < 25; i++ {
		//25个节气时刻(北京时间),从冬至开始到下一个冬至以后
		A[i] = getShuoQiDay(W+15.2184*float64(i), true)
	}
	l.pe1 = getShuoQiDay(W-15.2, true)
	l.pe2 = getShuoQiDay(W-30.4, true) //补算二气,确保一年中所有月份的“气”全部被计算在内

	//今年"首朔"的日月黄经差w
	var w = float64(getShuoQiDay(float64(A[0]), false)) //求较靠近冬至的朔日
	if w > float64(A[0]) {
		w -= 29.53
	}

	//该年所有朔,包含14个月的始末
	for i := 0; i < 15; i++ {
		B[i] = getShuoQiDay(w+29.5306*float64(i), false)
	}

	//月大小
	l.leap = 0
	var ym [14]int
	for i := 0; i < 14; i++ {
		l.dx[i] = l.hs[i+1] - l.hs[i] //月大小
		ym[i] = i                     //月序初始化
	}

	//-721年至-104年的后九月及月建问题,与朔有关，与气无关
	var YY = floorInt(float64(l.zq[0]+10+180)/365.2422) + 2000 //确定年份
	if YY >= -721 && YY <= -104 {
		var ns [9]any
		var yy int
		for i := 0; i < 3; i++ {
			yy = YY + i - 1
			//颁行历年首, 闰月名称, 月建
			if yy >= -721 { //春秋历,ly为-722.12.17
				ns[i] = getShuoQiDay(float64(1457698-j2000)+math.Floor(0.342+float64(yy+721)*12.368422)*29.5306, false)
				ns[i+3] = "十三"
				ns[i+6] = 2
			}
			if yy >= -479 { //战国历,ly为-480.12.11
				ns[i] = getShuoQiDay(float64(1546083-j2000)+math.Floor(0.5+float64(yy+479)*12.368422)*29.5306, false)
				ns[i+3] = "十三"
				ns[i+6] = 2
			}
			if yy >= -220 { //秦汉历,ly为-221.10.31
				ns[i] = getShuoQiDay(float64(1640641-j2000)+math.Floor(0.866+float64(yy+220)*12.369000)*29.5306, false)
				ns[i+3] = "后九"
				ns[i+6] = 11
			}
		}
		var nn, a int
		for i := 0; i < 14; i++ {
			for nn = 2; nn >= 0; nn-- {
				a = ns[nn].(int)
				if l.hs[i] >= a {
					break
				}
			}
			f1 := floorInt(float64(l.hs[i]-a+15) / 29.5306) //该月积数
			if f1 < 12 {
				n6 := ns[nn+6].(int)
				l.ymc[i] = ymc[(f1+n6)%12]
			} else {
				n3 := ns[nn+6].(int)
				ym[i] = n3
			}
		}
		return
	}

	//无中气置闰法确定闰月,(气朔结合法,数据源需有冬至开始的的气和朔)
	var i = 0
	if B[13] <= A[24] { //第13月的月末没有超过冬至(不含冬至),说明今年含有13个月
		for i = 1; B[i+1] > A[2*i] && i < 13; i++ {
			//在13个月中找第1个没有中气的月份
		}
		l.leap = i
		for ; i < 14; i++ {
			ym[i]--
		}
	}

	//名称转换(月建别名)
	for i = 0; i < 14; i++ {
		var Dm = l.hs[i] + j2000
		var v2 = ym[i]      //Dm初一的儒略日,v2为月建序号
		var mc = ymc[v2%12] //月建对应的默认月名称：建子十一,建丑十二,建寅为正……
		if Dm >= 1724360 && Dm <= 1729794 {
			mc = ymc[(v2+1)%12] //  8.01.15至 23.12.02 建子为十二,其它顺推
		} else if Dm >= 1807724 && Dm <= 1808699 {
			mc = ymc[(v2+1)%12] //237.04.12至239.12.13 建子为十二,其它顺推
		} else if Dm >= 1999349 && Dm <= 1999467 {
			mc = ymc[(v2+2)%12] //761.12.02至762.03.30 建子为正月,其它顺推
		} else if Dm >= 1973067 && Dm <= 1977052 {
			//689.12.18至700.11.15 建子为正月,建寅为一月,其它不变
			if v2%12 == 0 {
				mc = "正"
			}
			if v2 == 2 {
				mc = "一"
			}
		}

		if Dm == 1729794 || Dm == 1808699 {
			//239.12.13及23.12.02均为十二月,为避免两个连续十二月，此处改名
			mc = "拾贰"
		}

		l.ymc[i] = mc
	}
}

// updateLunar 相近月的信息使用此方法更新，无需重新定朔气
// 如果不在以计算好的范围，则会重新调用 NewLunar 创建 Lunar
func (l *Lunar) updateLunar(jd int) {
	l.jd = jd
	if jd < l.zq[0] || jd >= l.zq[24] {
		l.cal()
	}

	// 获取人看的信息
	l.calcLunarInfo()

	//// 月相与节气的处理
	//l.calcYXJQ()

	// 获取农历节日
	l.calcLunarEvents()
}

// 计算出给人看的信息
// 不能直接调用，外部调用使用 updateLunar
func (l *Lunar) calcLunarInfo() {
	// 农历所在月的序数
	var mk = floorInt(float64(l.jd-l.hs[0]) / 30)
	if mk < 13 && l.hs[mk+1] <= l.jd {
		mk++
	}
	l.IndexInLunarMonth = l.jd - l.hs[mk] //距农历月首的编移量,0对应初一
	l.DayName = rmc[l.IndexInLunarMonth]  //农历日名称
	l.CurDZ = l.jd - l.zq[0]              //距冬至的天数
	l.CurXZ = l.jd - l.zq[12]             //距夏至的天数
	l.CurLQ = l.jd - l.zq[15]             //距立秋的天数
	l.CurMZ = l.jd - l.zq[11]             //距芒种的天数
	l.CurXS = l.jd - l.zq[13]             //距小暑的天数

	l.MonthName = l.ymc[mk]    //月名称
	l.MonthDayCount = l.dx[mk] //月大小
	if l.leap == mk {
		l.LeapStr = "闰"
	} else {
		l.LeapStr = ""
	}
	if mk < 13 {
		l.NextMonthName = l.ymc[mk+1]
	} else {
		l.NextMonthName = "未知"
	}

	// 节气的取值范围是0-23
	var qk = floorInt(float64(l.jd-l.zq[0]-7) / 15.2184)
	if qk < 23 && l.jd >= l.zq[qk+1] {
		qk++
	}
	if l.jd == l.zq[qk] {
		l.SolarTermStr = jqmc[qk]
	} else {
		l.SolarTermStr = ""
	}

	//干支纪年处理
	//以立春为界定年首
	var D float64
	if l.jd < l.zq[3] {
		D = float64(l.zq[3]) - 365 + 365.25*16 - 35
	} else {
		D = float64(l.zq[3]) + 365.25*16 - 35
	}

	l.Year = floorInt(D/365.2422 + 0.5) //农历纪年(10进制,1984年起算)
	//以下几行以正月初一定年首
	D = float64(l.hs[2])      //一般第3个月为春节
	for j := 0; j < 14; j++ { //找春节
		if l.ymc[j] != "正" || l.leap == j && j > 0 {
			continue
		}
		D = float64(l.hs[j])
		if float64(l.jd) < D {
			D -= 365
			break
		} //无需再找下一个正月
	}
	D = D + 5810                           //计算该年春节与1984年平均春节(立春附近)相差天数估计
	var year0 = floorInt(D/365.2422 + 0.5) //农历纪年(10进制,1984年起算)

	D = float64(l.Year + 12000)
	l.Year2 = gan[int(D)%10] + zhi[int(D)%12] //干支纪年(立春)
	D = float64(year0 + 12000)
	l.Year3 = gan[int(D)%10] + zhi[int(D)%12] //干支纪年(正月)
	l.Year4 = year0 + 1984 + 2698             //黄帝纪年

	//纪月处理,1998年12月7(大雪)开始连续进行节气计数,0为甲子
	mk = floorInt(float64(l.jd-l.zq[0]) / 30.43685)
	//相对大雪的月数计算,mk的取值范围0-12
	if mk < 12 && l.jd >= l.zq[2*mk+1] {
		mk++
	}

	D = float64(mk + floorInt(float64(l.zq[12]+390)/365.2422)*12 + 900000) //相对于1998年12月7(大雪)的月数,900000为正数基数
	l.MonthGanZhi = int(D) % 12
	l.MonthGanZhiName = gan[int(D)%10] + zhi[int(D)%12]

	//纪日,2000年1月7日起算
	var dInt = l.jd - 6 + 9000000
	l.DayGanZhiName = gan[dInt%10] + zhi[dInt%12]
}

// 月相和节气的处理，需要自己调用
func (l *Lunar) calcYXJQ() {
	var jd = float64(l.jd) + dtT(float64(l.jd)) - 8.0/24
	////月相查找
	var w = msALon(jd/36525, 10, 3)
	w = math.Floor((w-0.78)/math.Pi*2) * math.Pi / 2
	var d float64
	var D, xn int
	l.PhasesOfMoon = ""
	l.phasesOfMoonJD = 0
	l.PhasesOfMoonTime = Time{}
	for {
		d = soAccurate(w)
		D = floorInt(d + 0.5)
		xn = floorInt(w/pi2*4+4000000.01) % 4
		w += pi2 / 4

		if D > l.jd {
			break
		}
		if D < l.jd {
			continue
		}
		l.PhasesOfMoon = yxmc[xn] //取得月相名称
		l.phasesOfMoonJD = d
		l.PhasesOfMoonTime = timeFromJD(d + float64(j2000))
		if D+5 >= l.jd {
			break
		}
	}

	//节气查找
	l.SolarTerm = ""
	l.solarTermJD = 0
	l.SolarTermTime = Time{}
	w = sALon(jd/36525, 3)
	w = math.Floor((w-0.13)/pi2*24) * pi2 / 24
	for {
		d = qiAccurate(w)
		D = floorInt(d + 0.5)
		xn = floorInt(w/pi2*24+24000006.01) % 24
		w += pi2 / 24
		if D > l.jd {
			break
		}
		if D < l.jd {
			continue
		}
		l.SolarTerm = jqmc[xn] //取得节气名称
		l.solarTermJD = d
		l.SolarTermTime = timeFromJD(d + float64(j2000))

		if D+12 >= l.jd {
			break
		}
	}

	//var Bd0, Bdn, D, xn int
	//var d, jd2 float64
	//if day == nil {
	//	day = NewDay(timeFromJD(float64(l.jd + j2000)))
	//}
	//Bd0 = day.getMonthFirstDaysOffJ2000()
	//Bdn = day.MonthDaysCount
	//jd2 = float64(Bd0) + dtT(float64(Bd0)) - 8.0/24
	////月相查找
	//var w = msALon(jd2/36525, 10, 3)
	//w = math.Floor((w-0.78)/math.Pi*2) * math.Pi / 2
	//
	//for {
	//	d = so_accurate(w)
	//	D = floorInt(d + 0.5)
	//	xn = floorInt(w/pi2*4+4000000.01) % 4
	//	w += pi2 / 4
	//	if D >= Bd0+Bdn {
	//		break
	//	}
	//	if D < Bd0 {
	//		continue
	//	}
	//	if D-Bd0 == day.IndexInMonth {
	//		l.PhasesOfMoon = PhasesOfMoon[xn] //取得月相名称
	//		l.phasesOfMoonJD = d
	//		l.PhasesOfMoonTime = timeFromJD(d + float64(j2000))
	//	}
	//
	//	if D+5 >= Bd0+Bdn {
	//		break
	//	}
	//}
	//
	////节气查找
	//w = sALon(jd2/36525, 3)
	//w = math.Floor((w-0.13)/pi2*24) * pi2 / 24
	//for {
	//	d = qi_accurate(w)
	//	D = floorInt(d + 0.5)
	//	xn = floorInt(w/pi2*24+24000006.01) % 24
	//	w += pi2 / 24
	//	if D >= Bd0+Bdn {
	//		break
	//	}
	//	if D < Bd0 {
	//		continue
	//	}
	//	if D-Bd0 == day.IndexInMonth {
	//		l.SolarTerm = SolarTerm[xn] //取得节气名称
	//		l.solarTermJD = d
	//		l.SolarTermTime = timeFromJD(d + float64(j2000))
	//	}
	//
	//	if D+12 >= Bd0+Bdn {
	//		break
	//	}
	//}
}

//// 月相和节气的处理
//func (l *Lunar) calcYXJQ(day *Day) {
//	var Bd0, Bdn, D, xn int
//	var d, jd2 float64
//	if day == nil {
//		day = NewDay(timeFromJD(float64(l.jd + j2000)))
//	}
//	Bd0 = day.getMonthFirstDaysOffJ2000()
//	Bdn = day.MonthDaysCount
//	jd2 = float64(Bd0) + dtT(float64(Bd0)) - 8.0/24
//	//月相查找
//	var w = msALon(jd2/36525, 10, 3)
//	w = math.Floor((w-0.78)/math.Pi*2) * math.Pi / 2
//
//	for {
//		d = so_accurate(w)
//		D = floorInt(d + 0.5)
//		xn = floorInt(w/pi2*4+4000000.01) % 4
//		w += pi2 / 4
//		if D >= Bd0+Bdn {
//			break
//		}
//		if D < Bd0 {
//			continue
//		}
//		if D-Bd0 == day.IndexInMonth {
//			l.PhasesOfMoon = PhasesOfMoon[xn] //取得月相名称
//			l.phasesOfMoonJD = d
//			l.PhasesOfMoonTime = timeFromJD(d + float64(j2000))
//		}
//
//		if D+5 >= Bd0+Bdn {
//			break
//		}
//	}
//
//	//节气查找
//	w = sALon(jd2/36525, 3)
//	w = math.Floor((w-0.13)/pi2*24) * pi2 / 24
//	for {
//		d = qi_accurate(w)
//		D = floorInt(d + 0.5)
//		xn = floorInt(w/pi2*24+24000006.01) % 24
//		w += pi2 / 24
//		if D >= Bd0+Bdn {
//			break
//		}
//		if D < Bd0 {
//			continue
//		}
//		if D-Bd0 == day.IndexInMonth {
//			l.SolarTerm = SolarTerm[xn] //取得节气名称
//			l.solarTermJD = d
//			l.SolarTermTime = timeFromJD(d + float64(j2000))
//		}
//
//		if D+12 >= Bd0+Bdn {
//			break
//		}
//	}
//}

// 获取农历节日事件
// 不能直接调用，外部调用使用 updateLunar
func (l *Lunar) calcLunarEvents() {
	l.Events = *new(Event)
	//按农历日期查找重量点节假日
	var d = l.MonthName
	if len([]rune(d)) < 2 {
		d += "月"
	}
	d += l.DayName
	if l.LeapStr != "闰" {
		if d == "正月初一" {
			l.Events.Festival = append(l.Events.Festival, "春节")
		}
		if d == "正月初二" {
			l.Events.Important = append(l.Events.Festival, "大年初二")
		}
		if d == "正月初六" {
			l.Events.Other = append(l.Events.Other, "送穷日")
		}
		if d == "五月初五" {
			l.Events.Festival = append(l.Events.Festival, "端午节")
		}
		if d == "八月十五" {
			l.Events.Festival = append(l.Events.Festival, "中秋节")
		}
		if d == "正月十五" {
			l.Events.Festival = append(l.Events.Festival, "元宵节")
			l.Events.Important = append(l.Events.Important, "上元节")
			l.Events.Other = append(l.Events.Other, "壮族歌墟节", "苗族踩山节", "达斡尔族卡钦")
		}
		if d == "正月十六" {
			l.Events.Other = append(l.Events.Other, "侗族芦笙节(至正月二十)")
		}
		if d == "正月廿五" {
			l.Events.Other = append(l.Events.Other, "填仓节")
		}

		if d == "二月初一" {
			l.Events.Other = append(l.Events.Other, "瑶族忌鸟节")
		}
		if d == "二月初二" {
			l.Events.Important = append(l.Events.Important, "春龙节(龙抬头)")
			l.Events.Other = append(l.Events.Other, "畲族会亲节")
		}
		if d == "二月初八" {
			l.Events.Other = append(l.Events.Other, "傈傈族刀杆节")
		}
		if d == "三月初三" {
			l.Events.Important = append(l.Events.Important, "北帝诞")
			l.Events.Other = append(l.Events.Other, "苗族黎族歌墟节")
		}
		if d == "三月十五" {
			l.Events.Other = append(l.Events.Other, "白族三月街(至三月二十)")
		}
		if d == "三月廿三" {
			l.Events.Important = append(l.Events.Important, "天后诞", "妈祖诞")
		}
		if d == "四月初八" {
			l.Events.Important = append(l.Events.Important, "牛王诞")
		}
		if d == "四月十八" {
			l.Events.Other = append(l.Events.Other, "锡伯族西迁节")
		}
		if d == "五月十三" {
			l.Events.Important = append(l.Events.Important, "关帝诞")
			l.Events.Other = append(l.Events.Other, "阿昌族泼水节")
		}
		if d == "五月廿二" {
			l.Events.Other = append(l.Events.Other, "鄂温克族米阔鲁节")
		}
		if d == "五月廿九" {
			l.Events.Other = append(l.Events.Other, "瑶族达努节")
		}
		if d == "六月初六" {
			l.Events.Important = append(l.Events.Important, "姑姑节", "天贶节")
			l.Events.Other = append(l.Events.Other, "壮族祭田节", "瑶族尝新节")
		}
		if d == "六月廿四" {
			l.Events.Other = append(l.Events.Other, "火把节、星回节(彝、白、佤、阿昌、纳西、基诺族)")
		}
		if d == "七月初七" {
			l.Events.Important = append(l.Events.Important, "七夕(中国情人节,乞巧节,女儿节)")
		}
		if d == "七月十三" {
			l.Events.Other = append(l.Events.Other, "侗族吃新节")
		}
		if d == "七月十五" {
			l.Events.Important = append(l.Events.Important, "中元节、鬼节")
		}
		if d == "九月初九" {
			l.Events.Important = append(l.Events.Important, "重阳节")
		}
		if d == "十月初一" {
			l.Events.Important = append(l.Events.Important, "祭祖节(十月朝)")
		}
		if d == "十月十五" {
			l.Events.Important = append(l.Events.Important, "下元节")
		}
		if d == "十月十六" {
			l.Events.Other = append(l.Events.Other, "瑶族盘王节")
		}
		if d == "十二初八" {
			l.Events.Important = append(l.Events.Important, "腊八节")
		}
	}
	if l.NextMonthName == "正" { //最后一月
		if d == "十二三十" && l.MonthDayCount == 30 {
			l.Events.Festival = append(l.Events.Festival, "除夕")
		}
		if d == "十二廿九" && l.MonthDayCount == 29 {
			l.Events.Festival = append(l.Events.Festival, "除夕")
		}
		if d == "十二廿三" {
			l.Events.Important = append(l.Events.Important, "北方小年")
		}
		if d == "十二廿四" {
			l.Events.Important = append(l.Events.Important, "南方小年")
		}
	}
	if l.SolarTermStr != "" {
		l.Events.Important = append(l.Events.Important, l.SolarTermStr)
	}

	//农历杂节
	var w, w2 string
	if l.CurDZ >= 0 && l.CurDZ < 81 { //数九
		w = numCn[l.CurDZ/9+1]
		if l.CurDZ%9 == 0 {
			l.Events.Important = append(l.Events.Important, w+"九")
		} else {
			l.Events.Other = append(l.Events.Other, w+"九第"+strconv.Itoa(l.CurDZ%9+1)+"天")
		}
	}

	w = string([]rune(l.DayGanZhiName)[0])
	w2 = string([]rune(l.DayGanZhiName)[1])
	if l.CurXZ >= 20 && l.CurXZ < 30 && w == "庚" {
		l.Events.Important = append(l.Events.Important, "初伏")
	}
	if l.CurXZ >= 30 && l.CurXZ < 40 && w == "庚" {
		l.Events.Important = append(l.Events.Important, "中伏")
	}
	if l.CurLQ >= 0 && l.CurLQ < 10 && w == "庚" {
		l.Events.Important = append(l.Events.Important, "末伏")
	}
	if l.CurMZ >= 0 && l.CurMZ < 10 && w == "丙" {
		l.Events.Important = append(l.Events.Important, "入梅")
	}
	if l.CurXS >= 0 && l.CurXS < 12 && w2 == "未" {
		l.Events.Important = append(l.Events.Important, "出梅")
	}
}

// jd 应靠近所要取得的气朔日,isQi 表示是否为'气'，为'气'时算节气的儒略日
func getShuoQiDay(jd float64, isQi bool) int {
	jd += float64(j2000)
	var kb = shuoKB[0:]
	var pc = 14.0
	if isQi {
		kb = qiKB[0:]
		pc = 7
	}
	var f1 = kb[0] - pc
	var f2 = kb[len(kb)-1] - pc
	var f3 float64 = 2436935
	if jd < f1 || jd >= f3 { //平气朔表中首个之前，使用现代天文算法。1960.1.1以后，使用现代天文算法 (这一部分调用了qi_high和so_high,所以需星历表支持)
		if isQi {
			// 2451259是1999.3.21,太阳视黄经为0,春分.定气计算
			return floorInt(qiHigh(math.Floor((jd+pc-2451259)/365.2422*24)*math.Pi/12) + 0.5)
		} else {
			// 2451551是2000.1.7的那个朔日,黄经差为0.定朔计算
			return floorInt(soHigh(math.Floor((jd+pc-2451551)/29.5306)*math.Pi*2) + 0.5)
		}
	}
	if jd >= f1 && jd < f2 { //平气或平朔
		var i = 0
		for i = 0; i < len(kb); i += 2 {
			if jd+pc < kb[i+2] {
				break
			}
		}
		var d = kb[i] + kb[i+1]*math.Floor((jd+pc-kb[i])/kb[i+1])
		var dInt = floorInt(d + 0.5)
		//如果使用太初历计算-103年1月24日的朔日,结果得到的是23日,这里修正为24日(实历)。修正后仍不影响-103的无中置闰。如果使用秦汉历，得到的是24日，本行D不会被执行。
		if dInt == 1683460 {
			dInt++
		}
		return dInt - j2000
	}
	if jd >= f2 && jd < f3 { //定气或定朔
		var D int
		var n string
		if isQi {
			// 2451259是1999.3.21,太阳视黄经为0,春分.定气计算
			D = floorInt(qiLow(math.Floor((jd+pc-2451259)/365.2422*24)*math.Pi/12) + 0.5)
			index := floorInt((jd - f2) / 365.2422 * 24)
			// 找定气修正值
			n = qiFT[index : index+1]
		} else {
			// 2451551是2000.1.7的那个朔日,黄经差为0.定朔计算
			D = floorInt(soLow(math.Floor((jd+pc-2451551)/29.5306)*math.Pi*2) + 0.5)
			index := floorInt((jd - f2) / 29.5306)
			// 找定朔修正值
			n = shuoFT[index : index+1]
		}
		if n == "1" {
			return D + 1
		}
		if n == "2" {
			return D - 1
		}
		return D
	}
	return 0
}

// 低精度定朔计算,在2000年至600，误差在2小时以内(仍比古代日历精准很多)
func soLow(shuo float64) float64 {
	var v = 7771.37714500204
	var t = (shuo + 1.08472) / v
	var off = (-0.0000331*t*t+0.10976*math.Cos(0.785+8328.6914*t)+0.02224*math.Cos(0.187+7214.0629*t)-0.03342*math.Cos(4.669+628.3076*t))/v + (32*(t+1.8)*(t+1.8)-20)/86400/36525
	t -= off
	return t*36525 + 8.0/24
}

// 最大误差小于30分钟，平均5分
func qiLow(qi float64) float64 {
	var t, L float64
	var v = 628.3319653318
	// 第一次估算,误差2天以内
	t = (qi - 4.895062166) / v
	// 第二次估算,误差2小时以内
	t -= (53*t*t + 334116*math.Cos(4.67+628.307585*t) + 2061*math.Cos(2.678+628.3076*t)*t) / v / 10000000

	L = 48950621.66 + 6283319653.318*t + 53*t*t + //平黄经
		334166*math.Cos(4.669257+628.307585*t) + //地球椭圆轨道级数展开
		3489*math.Cos(4.6261+1256.61517*t) + //地球椭圆轨道级数展开
		2060.6*math.Cos(2.67823+628.307585*t)*t - //一次泊松项
		994 - 834*math.Sin(2.1824-33.75705*t) //光行差与章动修正

	t -= (L/10000000-qi)/628.332 + (32*(t+1.8)*(t+1.8)-20)/86400/36525
	return t*36525 + 8/24
}

// 较高精度气
func qiHigh(qi float64) float64 {
	var t = sALonT2(qi) * 36525
	t = t - dtT(t) + 8.0/24
	var v = math.Mod(t+0.5, 1.0) * 86400
	if v < 1200 || v > 86400-1200 {
		t = sALonT(qi)*36525 - dtT(t) + 8.0/24
	}
	return t
}

// 较高精度朔
func soHigh(shuo float64) float64 {
	var t = msALonT2(shuo) * 36525
	t = t - dtT(t) + 8.0/24
	var v = math.Mod(t+0.5, 1) * 86400
	if v < 1800 || v > 86400-1800 {
		t = msALonT(shuo)*36525 - dtT(t) + 8.0/24
	}
	return t
}

// 精气
func qiAccurate(W float64) float64 {
	var t = sALonT(W) * 36525
	return t - dtT(t) + 8.0/24
}

//精朔
func soAccurate(W float64) float64 {
	var t = msALonT(W) * 36525
	return t - dtT(t) + 8.0/24
}
