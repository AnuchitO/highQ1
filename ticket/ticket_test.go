package ticket

import "testing"

var testCases = []struct {
	order Order
	want  float64
}{
	{Order{Age: 0}, 0},
	{Order{Age: 4}, 12},
	{Order{Age: 12}, 12},
	{Order{Age: 13}, 25},
	{Order{Age: 65}, 25},
	{Order{Age: 66}, 5},
}

func TestPriceAll(t *testing.T) {
	for _, tt := range testCases {
		got := Price(tt.order)

		if got != tt.want {
			t.Fatalf("actual %v, expected %v", got, tt.want)
		}
	}
}