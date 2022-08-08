package almanac

import (
	"math"
	"strconv"
)

type lunar struct {
	Lyear     int    // 农历纪年(10进制,1984年起算)
	Lyear2    string // 干支纪年（立春）
	Lyear3    string // 干支纪年（春节/正月）
	Lyear4    int    // 黄帝纪年
	Lmc       string // 月名称
	Ldn       int    // 月大小
	LeapStr   string // 闰月情况（"闰"/""）
	Lmc2      string // 下个月名称，判断除夕时需要用到
	Lmonth    int    // 干支纪月
	LmonthStr string // 干支纪月名称
	Lday2     string // 干支纪日

	Ldi  int    // 距农历月首的编移量,0对应初一
	Ldc  string // 农历日名称
	Ljq  string // 节气名称
	yxmc string // 月相名称
	yxjd string // 月相时刻(儒略日)
	yxsj string // 月相时间串
	jqmc string // 定气名称
	jqjd int    // 节气时刻(儒略日)
	jqsj string // 节气时间串

	CurDZ int // 距冬至的天数
	CurXZ int // 距夏至的天数
	CurLQ int // 距立秋的天数
	CurMZ int // 距芒种的天数
	CurXS int // 距小暑的天数

	events event // 节日、假期等事件

	jd int // 儒略日
	// 补算二气,确保一年中所有月份的“气”全部被计算在内
	pe1, pe2 int
	leap     int        // 闰月位置
	ZQ       [25]int    // 中气表,其中.liqiu 是节气立秋的儒略日,计算三伏时用到
	HS       [15]int    // 合朔表
	dx       [14]int    // 各月大小
	ymc      [14]string // 各月名称
}

// NewLunar 生成农历排序,  jd 为儒略日相对于 J2000 的偏移
// 时间系统全部使用北京时，即使是天象时刻的输出，也是使用北京时
// 如果天象的输出不使用北京时，会造成显示混乱，更严重的是无法与古历比对
// 注意：有 lunar 对象之后，建议使用 updateLunar 来更新农历，减少计算次数
func NewLunar(jd int) (l *lunar) {
	l = new(lunar)
	l.jd = jd
	l.cal()
	l.updateLunar(l.jd)
	return
}

// 排月序(生成实际年历)
func (l *lunar) cal() {
	var A, B = &l.ZQ, &l.HS //中气表,日月合朔表(整日)
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
		l.dx[i] = l.HS[i+1] - l.HS[i] //月大小
		ym[i] = i                     //月序初始化
	}

	//-721年至-104年的后九月及月建问题,与朔有关，与气无关
	var YY = floorInt(float64(l.ZQ[0]+10+180)/365.2422) + 2000 //确定年份
	if YY >= -721 && YY <= -104 {
		var ns [9]any
		var yy int
		for i := 0; i < 3; i++ {
			yy = YY + i - 1
			//颁行历年首, 闰月名称, 月建
			if yy >= -721 { //春秋历,ly为-722.12.17
				ns[i] = getShuoQiDay(float64(1457698-J2000)+math.Floor(0.342+float64(yy+721)*12.368422)*29.5306, false)
				ns[i+3] = "十三"
				ns[i+6] = 2
			}
			if yy >= -479 { //战国历,ly为-480.12.11
				ns[i] = getShuoQiDay(float64(1546083-J2000)+math.Floor(0.5+float64(yy+479)*12.368422)*29.5306, false)
				ns[i+3] = "十三"
				ns[i+6] = 2
			}
			if yy >= -220 { //秦汉历,ly为-221.10.31
				ns[i] = getShuoQiDay(float64(1640641-J2000)+math.Floor(0.866+float64(yy+220)*12.369000)*29.5306, false)
				ns[i+3] = "后九"
				ns[i+6] = 11
			}
		}
		var nn, a int
		for i := 0; i < 14; i++ {
			for nn = 2; nn >= 0; nn-- {
				a = ns[nn].(int)
				if l.HS[i] >= a {
					break
				}
			}
			f1 := floorInt(float64(l.HS[i]-a+15) / 29.5306) //该月积数
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
		var Dm = l.HS[i] + J2000
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
// 如果不在以计算好的范围，则会重新调用 NewLunar 创建 lunar
func (l *lunar) updateLunar(jd int) {
	l.jd = jd
	if jd < l.ZQ[0] || jd >= l.ZQ[24] {
		l.cal()
	}

	// 获取人看的信息
	l.calcLunarInfo()

	// 获取农历节日
	l.calcLunarEvents()
}

// 计算出给人看的信息
// 不能直接调用，外部调用使用 updateLunar
func (l *lunar) calcLunarInfo() {
	// 农历所在月的序数
	var mk = floorInt(float64(l.jd-l.HS[0]) / 30)
	if mk < 13 && l.HS[mk+1] <= l.jd {
		mk++
	}
	l.Ldi = l.jd - l.HS[mk]   //距农历月首的编移量,0对应初一
	l.Ldc = rmc[l.Ldi]        //农历日名称
	l.CurDZ = l.jd - l.ZQ[0]  //距冬至的天数
	l.CurXZ = l.jd - l.ZQ[12] //距夏至的天数
	l.CurLQ = l.jd - l.ZQ[15] //距立秋的天数
	l.CurMZ = l.jd - l.ZQ[11] //距芒种的天数
	l.CurXS = l.jd - l.ZQ[13] //距小暑的天数

	l.Lmc = l.ymc[mk] //月名称
	l.Ldn = l.dx[mk]  //月大小
	if l.leap == mk {
		l.LeapStr = "闰"
	} else {
		l.LeapStr = ""
	}
	if mk < 13 {
		l.Lmc2 = l.ymc[mk+1]
	} else {
		l.Lmc2 = "未知"
	}

	// 节气的取值范围是0-23
	var qk = floorInt(float64(l.jd-l.ZQ[0]-7) / 15.2184)
	if qk < 23 && l.jd >= l.ZQ[qk+1] {
		qk++
	}
	if l.jd == l.ZQ[qk] {
		l.Ljq = jqmc[qk]
	} else {
		l.Ljq = ""
	}

	//干支纪年处理
	//以立春为界定年首
	var D float64
	if l.jd < l.ZQ[3] {
		D = float64(l.ZQ[3]) - 365 + 365.25*16 - 35
	} else {
		D = float64(l.ZQ[3]) + 365.25*16 - 35
	}

	l.Lyear = floorInt(D/365.2422 + 0.5) //农历纪年(10进制,1984年起算)
	//以下几行以正月初一定年首
	D = float64(l.HS[2])      //一般第3个月为春节
	for j := 0; j < 14; j++ { //找春节
		if l.ymc[j] != "正" || l.leap == j && j > 0 {
			continue
		}
		D = float64(l.HS[j])
		if float64(l.jd) < D {
			D -= 365
			break
		} //无需再找下一个正月
	}
	D = D + 5810                            //计算该年春节与1984年平均春节(立春附近)相差天数估计
	var lyear0 = floorInt(D/365.2422 + 0.5) //农历纪年(10进制,1984年起算)

	D = float64(l.Lyear + 12000)
	l.Lyear2 = gan[int(D)%10] + zhi[int(D)%12] //干支纪年(立春)
	D = float64(lyear0 + 12000)
	l.Lyear3 = gan[int(D)%10] + zhi[int(D)%12] //干支纪年(正月)
	l.Lyear4 = lyear0 + 1984 + 2698            //黄帝纪年

	//纪月处理,1998年12月7(大雪)开始连续进行节气计数,0为甲子
	mk = floorInt(float64(l.jd-l.ZQ[0]) / 30.43685)
	//相对大雪的月数计算,mk的取值范围0-12
	if mk < 12 && l.jd >= l.ZQ[2*mk+1] {
		mk++
	}

	D = float64(mk + floorInt(float64(l.ZQ[12]+390)/365.2422)*12 + 900000) //相对于1998年12月7(大雪)的月数,900000为正数基数
	l.Lmonth = int(D) % 12
	l.LmonthStr = gan[int(D)%10] + zhi[int(D)%12]

	//纪日,2000年1月7日起算
	var dInt = l.jd - 6 + 9000000
	l.Lday2 = gan[dInt%10] + zhi[dInt%12]
}

// 获取农历节日事件
// 不能直接调用，外部调用使用 updateLunar
func (l *lunar) calcLunarEvents() {
	l.events = *new(event)
	//按农历日期查找重量点节假日
	var d = l.Lmc
	if len([]rune(d)) < 2 {
		d += "月"
	}
	d += l.Ldc
	if l.LeapStr != "闰" {
		if d == "正月初一" {
			l.events.festival = append(l.events.festival, "春节")
		}
		if d == "正月初二" {
			l.events.important = append(l.events.festival, "大年初二")
		}
		if d == "正月初六" {
			l.events.other = append(l.events.other, "送穷日")
		}
		if d == "五月初五" {
			l.events.festival = append(l.events.festival, "端午节")
		}
		if d == "八月十五" {
			l.events.festival = append(l.events.festival, "中秋节")
		}
		if d == "正月十五" {
			l.events.festival = append(l.events.festival, "元宵节")
			l.events.important = append(l.events.important, "上元节")
			l.events.other = append(l.events.other, "壮族歌墟节", "苗族踩山节", "达斡尔族卡钦")
		}
		if d == "正月十六" {
			l.events.other = append(l.events.other, "侗族芦笙节(至正月二十)")
		}
		if d == "正月廿五" {
			l.events.other = append(l.events.other, "填仓节")
		}

		if d == "二月初一" {
			l.events.other = append(l.events.other, "瑶族忌鸟节")
		}
		if d == "二月初二" {
			l.events.important = append(l.events.important, "春龙节(龙抬头)")
			l.events.other = append(l.events.other, "畲族会亲节")
		}
		if d == "二月初八" {
			l.events.other = append(l.events.other, "傈傈族刀杆节")
		}
		if d == "三月初三" {
			l.events.important = append(l.events.important, "北帝诞")
			l.events.other = append(l.events.other, "苗族黎族歌墟节")
		}
		if d == "三月十五" {
			l.events.other = append(l.events.other, "白族三月街(至三月二十)")
		}
		if d == "三月廿三" {
			l.events.important = append(l.events.important, "天后诞", "妈祖诞")
		}
		if d == "四月初八" {
			l.events.important = append(l.events.important, "牛王诞")
		}
		if d == "四月十八" {
			l.events.other = append(l.events.other, "锡伯族西迁节")
		}
		if d == "五月十三" {
			l.events.important = append(l.events.important, "关帝诞")
			l.events.other = append(l.events.other, "阿昌族泼水节")
		}
		if d == "五月廿二" {
			l.events.other = append(l.events.other, "鄂温克族米阔鲁节")
		}
		if d == "五月廿九" {
			l.events.other = append(l.events.other, "瑶族达努节")
		}
		if d == "六月初六" {
			l.events.important = append(l.events.important, "姑姑节", "天贶节")
			l.events.other = append(l.events.other, "壮族祭田节", "瑶族尝新节")
		}
		if d == "六月廿四" {
			l.events.other = append(l.events.other, "火把节、星回节(彝、白、佤、阿昌、纳西、基诺族)")
		}
		if d == "七月初七" {
			l.events.important = append(l.events.important, "七夕(中国情人节,乞巧节,女儿节)")
		}
		if d == "七月十三" {
			l.events.other = append(l.events.other, "侗族吃新节")
		}
		if d == "七月十五" {
			l.events.important = append(l.events.important, "中元节、鬼节")
		}
		if d == "九月初九" {
			l.events.important = append(l.events.important, "重阳节")
		}
		if d == "十月初一" {
			l.events.important = append(l.events.important, "祭祖节(十月朝)")
		}
		if d == "十月十五" {
			l.events.important = append(l.events.important, "下元节")
		}
		if d == "十月十六" {
			l.events.other = append(l.events.other, "瑶族盘王节")
		}
		if d == "十二初八" {
			l.events.important = append(l.events.important, "腊八节")
		}
	}
	if l.Lmc2 == "正" { //最后一月
		if d == "十二三十" && l.Ldn == 30 {
			l.events.festival = append(l.events.festival, "除夕")
		}
		if d == "十二廿九" && l.Ldn == 29 {
			l.events.festival = append(l.events.festival, "除夕")
		}
		if d == "十二廿三" {
			l.events.important = append(l.events.important, "北方小年")
		}
		if d == "十二廿四" {
			l.events.important = append(l.events.important, "南方小年")
		}
	}
	if l.Ljq != "" {
		l.events.important = append(l.events.important, l.Ljq)
	}

	//农历杂节
	var w, w2 string
	if l.CurDZ >= 0 && l.CurDZ < 81 { //数九
		w = numCn[l.CurDZ/9+1]
		if l.CurDZ%9 == 0 {
			l.events.important = append(l.events.important, w+"九")
		} else {
			l.events.other = append(l.events.other, w+"九第"+strconv.Itoa(l.CurDZ%9+1)+"天")
		}
	}

	w = string([]rune(l.Lday2)[0])
	w2 = string([]rune(l.Lday2)[1])
	if l.CurXZ >= 20 && l.CurXZ < 30 && w == "庚" {
		l.events.important = append(l.events.important, "初伏")
	}
	if l.CurXZ >= 30 && l.CurXZ < 40 && w == "庚" {
		l.events.important = append(l.events.important, "中伏")
	}
	if l.CurLQ >= 0 && l.CurLQ < 10 && w == "庚" {
		l.events.important = append(l.events.important, "末伏")
	}
	if l.CurMZ >= 0 && l.CurMZ < 10 && w == "丙" {
		l.events.important = append(l.events.important, "入梅")
	}
	if l.CurXS >= 0 && l.CurXS < 12 && w2 == "未" {
		l.events.important = append(l.events.important, "出梅")
	}
}

// jd 应靠近所要取得的气朔日,isQi 表示是否为'气'，为'气'时算节气的儒略日
func getShuoQiDay(jd float64, isQi bool) int {
	jd += float64(J2000)
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
		return dInt - J2000
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
	var t = S_aLon_t2(qi) * 36525
	t = t - dt_T(t) + 8.0/24
	var v = math.Mod(t+0.5, 1.0) * 86400
	if v < 1200 || v > 86400-1200 {
		t = S_aLon_t(qi)*36525 - dt_T(t) + 8.0/24
	}
	return t
}

// 较高精度朔
func soHigh(shuo float64) float64 {
	var t = MS_aLon_t2(shuo) * 36525
	t = t - dt_T(t) + 8.0/24
	var v = math.Mod(t+0.5, 1) * 86400
	if v < 1800 || v > 86400-1800 {
		t = MS_aLon_t(shuo)*36525 - dt_T(t) + 8.0/24
	}
	return t
}
