package fizzbuzz_test

import (
	"testing"

	"github.com/anuchito/highQ/fizzbuzz"
)

func TestFizzBuzzExpose(t *testing.T) {
	t.Run("should returns 1 when input 1", func(t *testing.T) {
		input := 1
		want := "1"

		got := fizzbuzz.FB(input)

		if got != want {
			t.Errorf("FizzBuzz(%v) should return %v", input, want)
		}
	})
}
