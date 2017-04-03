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

func TestFizzWithMultipleOf3(t *testing.T) {
	printer := newStringPrinter()
	FizzBuzz(printer, 3)

	if msg := printer.String(); msg != "Fizz" {
		t.Errorf("incorrect message for 3: got %s, expected Fizz", msg)
	}
}

func TestFizzWithNonMultipleOf3(t *testing.T) {
	printer := newStringPrinter()
	FizzBuzz(printer, 4)

	if msg := printer.String(); msg != "4" {
		t.Errorf("incorrect message for 4: got %s, expected 4", msg)
	}
}
