package fizzbuzz

func FizzBuzz(n int) string {
	switch {
	case (n%3 == 0 && n%5 == 0):
		return "FizzBuzz?"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return ""
	}
}
