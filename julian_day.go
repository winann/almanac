package almanac

import (
	"math"
)

type JulianDay = float64

// julianDay 公历转儒略日，传入参数年、月参数为整型；日的参数为浮点型，小数部分表示一天中的时刻
// 儒略日数是指从公元 -4712 年开始连续计算日数得出的天数及不满一日的小数，通常记为 JD。
// 传统上儒略日的计数是从格林尼治平午，即世界时12点开始的。
// 返回值 JulianDay 即为儒略日
func julianDay(year, month int, day JulianDay) (jd JulianDay) {
	var n = 0
	// 判断是否为格里高利历日1582*372+10*31+15
	var isGregorian = year*372+month*31+int(day) >= 588829
	if month <= 2 {
		month += 12
		year--
	}
	if isGregorian {
		n = int(year / 100)
		//加百年闰
		n = 2 - n + int(n/4)
	}
	jd = math.Floor(365.25*float64(year+4716)) + math.Floor(30.6001*float64(month+1)) + day + float64(n) - 1524.5
	return
}

// timeFromJD 根据儒略日得到当前的日期时间
// 返回 AlmanacTime 对象
func timeFromJD(julianDay JulianDay) (time Time) {
	// 取得日数的整数部分i及小数部分f
	var i = int(julianDay + 0.50000001)
	var f = julianDay + 0.50000001 - float64(i)
	var c int
	if i >= 2299161 {
		c = int((float64(i) - 1867216.25) / 36524.25)
		i += 1 + c - int(c/4)
	}
	i += 1524

	// 年
	time.year = int((float64(i) - 122.1) / 365.25)
	// 月
	i -= int(365.25 * float64(time.year))
	time.month = int(float64(i) / 30.601)
	// 日
	i -= int(30.601 * float64(time.month))
	time.day = i

	if time.month > 13 {
		time.month -= 13
		time.year -= 4715
	} else {
		time.month -= 1
		time.year -= 4716
	}
	// 日的小数转化为时分秒
	// 时
	f *= 24
	time.hour = int(f)
	// 分
	f -= float64(time.hour)
	f *= 60
	time.minute = int(f)
	// 秒
	f -= float64(time.minute)
	f *= 60
	time.second = int(f)
	return
}

//func timeFromJD(julianDay JulianDay) (Time AlmanacTime) {
//	//// 取得日数的整数部分i及小数部分f
//	jd := big.NewFloat(julianDay)
//	var fi = jd.Add(jd, big.NewFloat(0.500001))
//	var bi, _ = fi.Int(nil)
//	fi = fi.Sub(fi, new(big.Float).SetInt(bi))
//
//	if bi.Cmp(big.NewInt(2299161)) > -1 {
//		i := bi.Int64()
//		c := int64((float64(i) - 1867216.25) / 36524.25)
//		i += 1 + c - int64(c/4)
//
//		a := new(big.Float).SetInt(bi)
//		a = a.Sub(a, big.NewFloat(1867216.25))
//		a = a.Quo(a, big.NewFloat(36524.25))
//		bc, _ := a.Int(nil)
//
//		bi = bi.Add(bi, big.NewInt(1))
//		bi = bi.Add(bi, bc)
//		bi = bi.Sub(bi, bc.Quo(bc, big.NewInt(4)))
//	}
//	bi = bi.Add(bi, big.NewInt(1524))
//	fi = fi.SetInt(bi)
//
//	y := fi.Sub(fi, big.NewFloat(122.1))
//	year, _ := y.Quo(y, big.NewFloat(365.25)).Int64()
//	Time.year = int(year)
//	offSet, _ := new(big.Float).Mul(big.NewFloat(365.25), new(big.Float).SetInt64(year)).Int(nil)
//	bi = bi.Sub(bi, offSet)
//	month, _ := new(big.Float).SetInt(bi).Quo(new(big.Float).SetInt(bi), big.NewFloat(30.601)).Int64()
//	Time.month = int(month)
//	offSet, _ = new(big.Float).Mul(big.NewFloat(30.601), new(big.Float).SetInt64(month)).Int(nil)
//	bi = bi.Sub(bi, offSet)
//	Day := bi.Int64()
//	Time.Day = int(Day)
//
//	if Time.month > 13 {
//		Time.month -= 13
//		Time.year -= 4715
//	} else {
//		Time.month -= 1
//		Time.year -= 4716
//	}
//
//	//// 日的小数转化为时分秒
//	jd2 := big.NewFloat(julianDay)
//	fi2 := jd2.Add(jd2, big.NewFloat(0.500001))
//	bi2, _ := fi2.Int(nil)
//
//	f := fi2.Sub(fi2, new(big.Float).SetInt(bi2))
//	f = f.Mul(f, big.NewFloat(24))
//	hour, _ := f.Int64()
//	Time.hour = int(hour)
//
//	f = f.Sub(f, new(big.Float).SetInt64(hour))
//	f = f.Mul(f, big.NewFloat(60))
//	minute, _ := f.Int64()
//	Time.minute = int(minute)
//
//	f = f.Sub(f, new(big.Float).SetInt64(minute))
//	f = f.Mul(f, big.NewFloat(60))
//	second, _ := f.Int64()
//	Time.second = int(second)
//	return
//}
