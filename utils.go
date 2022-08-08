package almanac

import "math"

const MIN = 0.000001

// isEqual
// 判断 float 是否相等
func isEqual(x, y float64) bool {
	return math.Abs(x-y) < MIN
}

// floorInt
// 获取小于等于某 float64 的最大整数
func floorInt(x float64) int {
	return int(math.Floor(x))
}
