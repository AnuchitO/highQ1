package sum

import "testing"

func TestSum(t *testing.T) {
	r := Sum(3, 5)

	if r != 8 {
		t.Error("Expected 8, got", r)
	}
}
