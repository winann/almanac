package almanac

// Event 上班/放假事件
type Event struct {
	festival  []string // 重要喜庆日子名称(可将日子名称置红)
	important []string // 重要日子名称
	other     []string // 非重要、喜庆日子名称
	isWeekend bool     // 是否是周末（如果 isWorkday 为 true 表示需要上班）
	isHoliday bool     // 是否是假期
	isWorkday bool     // 是否需要上班（如果 isWeekend 为 true 则说明是补班）
	holiday   string   // 如果是假期，则显示假期名称
}
