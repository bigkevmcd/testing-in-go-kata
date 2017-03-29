package fizzbuzz

// PRINTER OMIT
type Printer interface {
	Print(v ...interface{})
}

// ENDPRINTER OMIT

func FizzBuzz(p Printer, n int) {
	switch {
	case (n%3 == 0 && n%5 == 0):
		p.Print("FizzBuzz?")
	case n%3 == 0:
		p.Print("Fizz")
	case n%5 == 0:
		p.Print("Buzz")
	default:
		return
	}
}
