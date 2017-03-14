package fizzbuzz

import "testing"

func TestFizzWithMultipleOf3(t *testing.T) {
	if msg := FizzBuzz(3); msg != "Fizz" {
		t.Errorf("incorrect message for 3: got %s, expected Fizz", msg)
	}
}

func TestFizzWithNonMultipleOf3(t *testing.T) {
	if msg := FizzBuzz(4); msg == "Fizz" {
		t.Errorf("incorrect message for 4: got %s, expected ''", msg)
	}
}
