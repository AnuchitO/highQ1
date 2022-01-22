package binary

import "testing"

var tests = []struct {
	name  string
	input string
	want  int
}{
	{"one", "1", 1},
	{"two", "10", 2},
	{"three", "11", 3},
	{"four", "100", 4},
	{"nine", "1001", 9},
	{"twentysix", "11010", 26},
	{"name", "10001101000", 1128},
	{"name", "0", 0},
}

func TestParseBinary(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := ParseBinary(tt.input)

			if got != tt.want {
				t.Error("Expected ", tt.want, ", got", got)
			}
		})
	}
}

func TestParseInvalidBinary(t *testing.T) {
	want := `"a" is not a vaid digit`

	_, err := ParseBinary("1a")

	if err.Error() != want {
		t.Error("should return error:", want, "but got:", err.Error())
	}
}
