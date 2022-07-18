package almanac

import "testing"

func TestIsEqual(t *testing.T) {
	if !isEqual(30.6, 30.60000000001) {
		t.Error("isEqual error:", false, "But Expect:", true)
	}

	if isEqual(30.6, 30.60001) {
		t.Error("isEqual error:", true, "But Expect:", false)
	}
}
