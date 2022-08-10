package almanac

import "strconv"

type Dynasty struct {
	// 起始年
	startYear int

	// 使用年数
	duration int

	// 已用年数
	used int

	// 朝代名
	Name string

	// 朝代号
	Mark string

	// 皇帝
	Emperor string

	// 年号
	EraName string
}

// NewDynastyInfo 通过公元年获取朝代信息
func NewDynastyInfo(year int) (d *Dynasty) {

	var i = 0
	var info [7]string
	for ; i < len(eraNames); i++ {
		info = eraNames[i]
		startYear, _ := strconv.Atoi(info[0])
		duration, _ := strconv.Atoi(info[1])
		if startYear <= year && (startYear+duration) >= year {
			d = new(Dynasty)
			d.startYear = startYear
			d.duration = duration
			d.used, _ = strconv.Atoi(info[2])
			d.Name = info[3]
			d.Mark = info[4]
			d.Emperor = info[5]
			d.EraName = info[6]
		}
	}
	return
}
