package almanac

import "math"

const MIN = 0.000001

// isEqual
// 判断 float 是否相等
func isEqual(x, y float64) bool {
	return math.Abs(x-y) < MIN
}
