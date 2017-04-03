package fizzbuzz

import "strconv"

// PRINTER OMIT
type Printer func(string)

// ENDPRINTER OMIT

func FizzBuzz(p Printer, n int) {
	switch {
	case (n%3 == 0 && n%5 == 0):
		p("FizzBuzz")
	case n%3 == 0:
		p("Fizz")
	case n%5 == 0:
		p("Buzz")
	default:
		p(strconv.Itoa(n))
	}
}
