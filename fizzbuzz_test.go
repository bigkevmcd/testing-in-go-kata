package main

import "testing"

func FizzBuzz(n int) string {
	switch {
	case (n%3 == 0 && n%5 == 0):
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return ""
	}
}

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

func TestFizzWithMultipleOf5(t *testing.T) {
	if msg := FizzBuzz(5); msg != "Buzz" {
		t.Errorf("incorrect message for 5: got %s, expected Buzz", msg)
	}
}

func TestFizzBuzzTable(t *testing.T) {
	tests := []struct {
		n   int
		msg string
	}{
		{1, ""},
		{2, ""},
		{3, "Fizz"},
		{4, ""},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, ""},
		{8, ""},
		{9, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz"},
	}

	for _, test := range tests {
		if m := FizzBuzz(test.n); m != test.msg {
			t.Errorf("incorrect message for %d: got %s, expected %s", test.n, m, test.msg)
		}
	}
}
