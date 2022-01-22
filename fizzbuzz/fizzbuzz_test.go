package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	t.Run("should returns 1 when input 1", func(t *testing.T) {
		input := 1
		want := "1"

		got := FB(input)

		if got != want {
			t.Errorf("FizzBuzz(%v) should return %v", input, want)
		}
	})

	t.Run("should return 2 when input 2", func(t *testing.T) {
		input := 2
		want := "2"

		got := FB(input)

		if got != want {
			t.Errorf("FizzBuzz(%v) should return %v", input, want)
		}
	})

	t.Run("should returns Fizz when input 3", func(t *testing.T) {
		input := 3
		want := "Fizz"

		got := FB(input)

		if got != want {
			t.Errorf("FizzBuzz(%v) should return %v", input, want)
		}
	})

	t.Run("should returns Buzz when input 5", func(t *testing.T) {
		input := 5
		want := "Buzz"

		got := FB(input)

		if got != want {
			t.Errorf("FizzBuzz(%v) should return %v", input, want)
		}
	})

	t.Run("should returns FizzBuzz when input 15", func(t *testing.T) {
		input := 15
		want := "FizzBuzz"

		got := FB(input)

		if got != want {
			t.Errorf("FizzBuzz(%v) should return %v", input, want)
		}
	})
}
