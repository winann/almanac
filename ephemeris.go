package almanac

import "math"

//type ephemeris struct {
//}

////中精度章动计算,t是世纪数
//func nutation2(t float64) [2]float64 {
//	var (
//		i  = 0
//		c  = 0.0
//		a  = 0.0
//		t2 = t * t
//		B  = nutB
//		dL = 0.0
//		dE = 0.0
//	)
//	for i = 0; i < len(B); i += 5 {
//		c = B[i] + B[i+1]*t + B[i+2]*t2
//		if i == 0 {
//			a = -1.742 * t
//		} else {
//			a = 0
//		}
//		dL += (B[i+3] + a) * math.Sin(c)
//		dE += B[i+4] * math.Cos(c)
//	}
//	return [...]float64{dL / 100 / rad, dE / 100 / rad} //黄经章动,交角章动
//}

//只计算黄经章动
func nutationLon2(t float64) float64 {
	var (
		i  = 0
		a  = 0.0
		t2 = t * t
		dL = 0.0
		B  = nutB
	)

	for i = 0; i < len(B); i += 5 {
		if i == 0 {
			a = -1.742 * t
		} else {
			a = 0
		}
		dL += (B[i+3] + a) * math.Sin(B[i]+B[i+1]*t+B[i+2]*t2)
	}
	return dL / 100 / rad
}

// xl0Calc xt星体,zn坐标号,t儒略世纪数,n计算项数
func xl0Calc(xt, zn int, t, n float64) float64 {
	t /= 10 //转为儒略千年数
	var j int
	var v, tn, c float64 = 0, 1, 0
	var F = xl0[xt]
	var n1, n2, N float64
	var n0 float64
	var pn = zn*6 + 1
	var N0 = F[pn+1] - F[pn] //N0序列总数
	for i := 0; i < 6; i++ {
		n1 = F[pn+i]
		n2 = F[pn+1+i]
		n0 = n2 - n1
		if isEqual(n0, 0) {
			continue
		}
		if n < 0 {
			//确定项数
			N = n2
		} else {
			N = math.Floor(3*n*n0/N0+0.5) + n1
			if i != 0 {
				N += 3
			}
			if N > n2 {
				N = n2
			}
		}
		j = int(n1)
		c = 0
		for ; j < int(N); j += 3 {
			c += F[j] * math.Cos(F[j+1]+t*F[j+2])
		}
		v += c * tn
		tn *= t
	}
	v /= F[0]
	if xt == 0 { //地球
		var t2 = t * t
		var t3 = t2 * t //千年数的各次方
		if zn == 0 {
			v += (-0.0728 - 2.7702*t - 1.1019*t2 - 0.0996*t3) / rad
		}
		if zn == 1 {
			v += (+0.0000 + 0.0004*t + 0.0004*t2 - 0.0026*t3) / rad
		}
		if zn == 2 {
			v += (-0.0020 + 0.0044*t + 0.0213*t2 - 0.0250*t3) / 1000000
		}
	} else { //其它行星
		var dv = xl0Xzb[(xt-1)*3+zn]
		if zn == 0 {
			v += -3 * t / rad
		}
		if zn == 2 {
			v += dv / 1000000
		} else {
			v += dv / rad
		}
	}
	return v
}

func xl1Calc(zn int, t float64, n int) float64 { //计算月亮
	var ob = xl1[zn]
	var i, j, N int
	var F []float64
	var v float64
	var tn float64 = 1
	var c float64
	var (
		t2 = t * t
		t3 = t2 * t
		t4 = t3 * t
		t5 = t4 * t
		tx = t - 10
	)
	if zn == 0 {
		v += (3.81034409 + 8399.684730072*t - 3.319e-05*t2 + 3.11e-08*t3 - 2.033e-10*t4) * rad //月球平黄经(弧度)
		v += 5028.792262*t + 1.1124406*t2 + 0.00007699*t3 - 0.000023479*t4 - 0.0000000178*t5   //岁差(角秒)
		if tx > 0 {
			//对公元3000年至公元5000年的拟合,最大误差小于10角秒
			v += -0.866 + 1.43*tx + 0.054*tx*tx
		}
	}
	t2 /= 1e4
	t3 /= 1e8
	t4 /= 1e8

	n *= 6
	if n < 0 {
		n = len(ob[0])
	}
	for i = 0; i < len(ob); i++ {
		F = ob[i]
		N = int(float64(n*len(F))/float64(len(ob[0])) + 0.5)
		if i != 0 {
			N += 6
			if N >= len(F) {
				N = len(F)
			}
		}
		j = 0
		c = 0
		for ; j < N; j += 6 {
			c += F[j] * math.Cos(F[j+1]+t*F[j+2]+t2*F[j+3]+t3*F[j+4]+t4*F[j+5])
		}
		v += c * tn
		tn *= t
	}
	if zn != 2 {
		v /= rad
	}
	return v
}

// eLon 地球经度计算,返回Date分点黄经,传入世纪数、取项数
func eLon(t, n float64) float64 {
	return xl0Calc(0, 0, t, n)
}

// mLon 月球经度计算,返回Date分点黄经,传入世纪数,n是项数比例
func mLon(t float64, n int) float64 {
	return xl1Calc(0, t, n)
}

// msALonT2 已知月日视黄经差求时间,高速低精度,误差不超过600秒(只验算了几千年)
func msALonT2(w float64) float64 {
	var t, v float64 = 0, 7771.37714500204
	t = (w + 1.08472) / v
	var L float64
	var t2 = t * t
	t -= (-0.00003309*t2 + 0.10976*math.Cos(0.784758+8328.6914246*t+0.000152292*t2) + 0.02224*math.Cos(0.18740+7214.0628654*t-0.00021848*t2) - 0.03342*math.Cos(4.669257+628.307585*t)) / v
	L = mLon(t, 20) - (4.8950632 + 628.3319653318*t + 0.000005297*t*t + 0.0334166*math.Cos(4.669257+628.307585*t) + 0.0002061*math.Cos(2.67823+628.307585*t)*t + 0.000349*math.Cos(4.6261+1256.61517*t) - 20.5/rad)
	v = 7771.38 - 914*math.Sin(0.7848+8328.691425*t+0.0001523*t*t) - 179*math.Sin(2.543+15542.7543*t) - 160*math.Sin(0.1874+7214.0629*t)
	t += (w - L) / v
	return t
}

// sALonT2 已知太阳视黄经反求时间,高速低精度,最大误差不超过600秒
func sALonT2(w float64) float64 {
	var t, v float64 = 0, 628.3319653318
	t = (w - 1.75347 - math.Pi) / v
	t -= (0.000005297*t*t + 0.0334166*math.Cos(4.669257+628.307585*t) + 0.0002061*math.Cos(2.67823+628.307585*t)*t) / v
	t += (w - eLon(t, 8) - math.Pi + (20.5+17.2*math.Sin(2.1824-33.75705*t))/rad) / v
	return t
}

// eV 地球速度,t是世纪数,误差小于万分3
func eV(t float64) float64 {
	var f = 628.307585 * t
	return 628.332 + 21*math.Sin(1.527+f) + 0.44*math.Sin(1.48+f*2) + 0.129*math.Sin(5.82+f)*t + 0.00055*math.Sin(4.21+f)*t*t
}

// mV 月球速度计算,传入世经数
func mV(t float64) float64 {
	var v = 8399.71 - 914*math.Sin(0.7848+8328.691425*t+0.0001523*t*t) //误差小于5%
	v -= 179*math.Sin(2.543+15542.7543*t) +                            //误差小于0.3%+
		160*math.Sin(0.1874+7214.0629*t) +
		62*math.Sin(3.14+16657.3828*t) +
		34*math.Sin(4.827+16866.9323*t) +
		22*math.Sin(4.9+23871.4457*t) +
		12*math.Sin(2.59+14914.4523*t) +
		7*math.Sin(0.23+6585.7609*t) +
		5*math.Sin(0.9+25195.624*t) +
		5*math.Sin(2.32-7700.3895*t) +
		5*math.Sin(3.88+8956.9934*t) +
		5*math.Sin(0.49+7771.3771*t)
	return v
}

// msALonT 已知月日视黄经差求时间
func msALonT(w float64) float64 {
	var t float64
	var v = 7771.37714500204
	t = (w + 1.08472) / v
	t += (w - MsALon(t, 3, 3)) / v
	v = mV(t) - eV(t) //v的精度0.5%，详见原文
	t += (w - MsALon(t, 20, 10)) / v
	t += (w - MsALon(t, -1, 60)) / v
	return t
}

// sALonT 已知太阳视黄经反求时间
func sALonT(w float64) float64 {
	var t, v = 0.0, 628.3319653318
	t = (w - 1.75347 - math.Pi) / v
	v = eV(t) //v的精度0.03%，详见原文
	t += (w - sALon(t, 10)) / v
	v = eV(t) //再算一次v有助于提高精度,不算也可以
	t += (w - sALon(t, -1)) / v
	return t
}

//二次曲线外推
func dtExt(y, jsd float64) float64 {
	var dy = (y - 1820) / 100
	return -20 + jsd*dy*dy
}

//计算世界时与原子时之差,传入年
func dtCalc(y float64) float64 {
	var y0 = dtAt[len(dtAt)-2] //表中最后一年
	var t0 = dtAt[len(dtAt)-1] //表中最后一年的deltaT
	if y >= y0 {
		var jsd float64 = 31 //sjd是y1年之后的加速度估计。瑞士星历表jsd=31,NASA网站jsd=32,skmap 的jsd=29
		if y > y0+100 {
			return dtExt(y, jsd)
		}
		var v = dtExt(y, jsd)        //二次曲线外推
		var dv = dtExt(y0, jsd) - t0 //ye年的二次外推与te的差
		return v - dv*(y0+100-y)/100
	}
	var i, d = 0, dtAt
	for i = 0; i < len(d); i += 5 {
		if y < d[i+5] {
			break
		}
	}
	var (
		t1 = (y - d[i]) / (d[i+5] - d[i]) * 10
		t2 = t1 * t1
		t3 = t2 * t1
	)
	return d[i+1] + d[i+2]*t1 + d[i+3]*t2 + d[i+4]*t3
}

//传入儒略日(J2000起算),计算TD-UT(单位:日)
func dtT(t float64) float64 {
	return dtCalc(t/365.2425+2000) / 86400.0
}

// MsALon 月日视黄经的差值
func MsALon(t float64, Mn int, Sn float64) float64 {
	return mLon(t, Mn) + gxcMoonLon() - (eLon(t, Sn) + gxcSunLon(t) + math.Pi)
}

// sALon 太阳视黄经
func sALon(t, n float64) float64 {
	return eLon(t, n) + nutationLon2(t) + gxcSunLon(t) + math.Pi //注意，这里的章动计算很耗时
}

//=============================一些天文基本问题=====================================
//==================================================================================
//返回朔日的编号,jd应在朔日附近，允许误差数天
//func suoN(jd float64) int {
//	return int((jd + 8) / 29.5306)
//}

//太阳光行差,t是世纪数
func gxcSunLon(t float64) float64 {
	var v = -0.043126 + 628.301955*t - 0.000002732*t*t //平近点角
	var e = 0.016708634 - 0.000042037*t - 0.0000001267*t*t
	return (-20.49552 * (1 + e*math.Cos(v))) / rad //黄经光行差
}

////黄纬光行差
//func gxc_sunLat(t float64) float64 {
//	return 0
//}

//月球经度光行差,误差0.07"
func gxcMoonLon() float64 {
	return -3.4e-6
}

////月球纬度光行差,误差0.006"
//func gxc_moonLat(t float64) float64 {
//	return 0.063 * math.Sin(0.057+8433.4662*t+0.000064*t*t) / rad
//}
