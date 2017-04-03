package fizzbuzz

import (
	"fmt"
	"strings"
	"testing"
)

type stringPrinter struct {
	Entries []string
}

func (sp *stringPrinter) Print(v string) {
	sp.Entries = append(sp.Entries, fmt.Sprint(v))
}

func (sp stringPrinter) String() string {
	return strings.Join(sp.Entries, "\n")
}

func newStringPrinter() *stringPrinter {
	return &stringPrinter{Entries: make([]string, 0)}
}

func TestFizzBuzz(t *testing.T) {
	t.Run("Fizz with non-multiple of 3", setup(func(t *testing.T, p *stringPrinter) {

		FizzBuzz(p, 4)

		if msg := p.String(); msg != "4" {
			t.Errorf("incorrect message for 4: got %s, expected 4", msg)
		}

	}))

	t.Run("Fizz with multiple of 3", setup(func(t *testing.T, p *stringPrinter) {
		FizzBuzz(p, 3)

		if msg := p.String(); msg != "Fizz" {
			t.Errorf("incorrect message for 3: got %s, expected Fizz", msg)
		}
	}))

}

func setup(test func(t *testing.T, p *stringPrinter)) func(t *testing.T) {
	return func(t *testing.T) {
		p := newStringPrinter()
		test(t, p)
	}
}
