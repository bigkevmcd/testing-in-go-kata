package fizzbuzz

import "testing"

//START OMIT
func TestFizz(t *testing.T) {
	t.Run("With multiple of 3", func(t *testing.T) {
		if msg := FizzBuzz(3); msg != "Fizz" {
			t.Errorf("incorrect message for 3: got %s, expected Fizz", msg)
		}
	})

	t.Run("With non-multiple of 3", func(t *testing.T) {
		if msg := FizzBuzz(4); msg == "Fizz" {
			t.Errorf("incorrect message for 4: got %s, expected ''", msg)
		}
	})
}

//END OMIT
