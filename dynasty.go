package almanac

import "strconv"

type dynasty struct {
	// 起始年
	startYear int

	// 使用年数
	duration int

	// 已用年数
	used int

	// 朝代名
	name string

	// 朝代号
	mark string

	// 皇帝
	emperor string

	// 年号
	eraName string
}

// NewDynastyInfo 通过公元年获取朝代信息
func NewDynastyInfo(year int) (d *dynasty) {

	var i = 0
	var info [7]string
	for ; i < len(eraNames); i++ {
		info = eraNames[i]
		startYear, _ := strconv.Atoi(info[0])
		duration, _ := strconv.Atoi(info[1])
		if startYear <= year && (startYear+duration) >= year {
			d = new(dynasty)
			d.startYear = startYear
			d.duration = duration
			d.used, _ = strconv.Atoi(info[2])
			d.name = info[3]
			d.mark = info[4]
			d.emperor = info[5]
			d.eraName = info[6]
		}
	}
	return
}
