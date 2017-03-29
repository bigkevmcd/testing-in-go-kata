package fizzbuzz

import (
	"fmt"
	"strings"
	"testing"
)

type stringPrinter struct {
	Entries []string
}

func (sp *stringPrinter) Print(v ...interface{}) {
	sp.Entries = append(sp.Entries, fmt.Sprint(v...))
}

func (sp stringPrinter) String() string {
	return strings.Join(sp.Entries, "\n")
}

func newStringPrinter() *stringPrinter {
	return &stringPrinter{Entries: make([]string, 0)}
}

//START OMIT
func TestFizz(t *testing.T) {
	t.Run("With multiple of 3", func(t *testing.T) {
		printer := newStringPrinter()
		FizzBuzz(printer, 3)

		if msg := printer.String(); msg != "Fizz" {
			t.Errorf("incorrect message for 3: got %s, expected Fizz", msg)
		}
	})

	t.Run("With non-multiple of 3", func(t *testing.T) {
		printer := newStringPrinter()
		FizzBuzz(printer, 4)
		if msg := printer.String(); msg == "Fizz" {
			t.Errorf("incorrect message for 4: got %s, expected ''", msg)
		}
	})
}

//END OMIT
