package fizzbuzz

import (
	"strings"
	"testing"
)

func createPrinter(printed *[]string) func(string) {
	return func(s string) {
		*printed = append(*printed, s)
	}
}

func TestFizz(t *testing.T) {
	t.Run("With multiple of 3", func(t *testing.T) {
		var printed []string

		FizzBuzz(createPrinter(&printed), 3)

		if msg := strings.Join(printed, "\n"); msg != "Fizz" {
			t.Errorf("incorrect message for 3: got %s, expected Fizz", msg)
		}
	})

	t.Run("With non-multiple of 3", func(t *testing.T) {
		var printed []string

		FizzBuzz(createPrinter(&printed), 4)

		if msg := strings.Join(printed, "\n"); msg != "4" {
			t.Errorf("incorrect message for 4: got %s, expected 4", msg)
		}
	})
}
