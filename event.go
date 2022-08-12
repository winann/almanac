package almanac

// Event 上班/放假事件
type Event struct {
	// 重要喜庆日子名称(可将日子名称置红)
	Festival []string

	// 重要日子名称
	Important []string

	// 非重要、喜庆日子名称
	Other []string

	// 是否是周末（如果 IsWorkday 为 true 表示需要上班）
	IsWeekend bool

	// 是否是假期
	IsHoliday bool

	// 是否需要上班（如果 IsWeekend 为 true 则说明是补班）
	IsWorkday bool

	// 如果是假期，则显示假期名称
	Holiday string
}

// GetFestivalCount 获取重要喜庆节日的个数
// 为 iOS 提供的接口
func (e *Event) GetFestivalCount() int {
	return len(e.Festival)
}

// GetFestival 获取重要喜庆节日
// 为 iOS 提供的接口
func (e *Event) GetFestival(index int) (f string) {
	if 0 <= index && len(e.Festival) > index {
		f = e.Festival[index]
	}
	return
}

// GetImportantCount 获取重要节日的个数
// 为 iOS 提供的接口
func (e *Event) GetImportantCount() int {
	return len(e.Important)
}

// GetImportant 获取重要节日
// 为 iOS 提供的接口
func (e *Event) GetImportant(index int) (f string) {
	if 0 <= index && len(e.Important) > index {
		f = e.Important[index]
	}
	return
}

// GetOtherCount 获取其他节日的个数
// 为 iOS 提供的接口
func (e *Event) GetOtherCount() int {
	return len(e.Other)
}

// GetOther 获取其它节日
// 为 iOS 提供的接口
func (e *Event) GetOther(index int) (f string) {
	if 0 <= index && len(e.Other) > index {
		f = e.Other[index]
	}
	return
}
