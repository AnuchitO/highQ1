package sum

import "testing"

func TestSum(t *testing.T) {
	r := Sum(3, 5)

	if r != 8 {
		t.Error("Expected 8, got", r)
	}
}

func TestSumAll(t *testing.T) {
	xs := []int{1, -2, 3}

	r := Sum(xs...)

	if r != 2 {
		t.Error("Expected 6, got", r)
	}
}
