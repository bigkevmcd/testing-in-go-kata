package main

import "testing"

var numberTests = []struct {
	n        int
	expected string
}{
	{1, "1"},
	{2, "2"},
	{3, "Fizz"},
	{4, "4"},
	{5, "Buzz"},
	{15, "FizzBuzz"},
}

func TestTable(t *testing.T) {
	for _, tt := range numberTests {
		if msg := FizzBuzz(tt.n); msg != tt.expected {
			t.Errorf("Fizzbuzz(%v) incorrect: got %v, expected %v", tt.n, msg, tt.expected)
		}

	}

}

func TestSimpleNumber(t *testing.T) {
	if FizzBuzz(1) != "1" {
		t.Errorf("incorrect response: got %s, expected 1", FizzBuzz(1))
	}
}

func TestFizzNumber(t *testing.T) {
	if FizzBuzz(3) != "Fizz" {
		t.Errorf("incorrect response: got %s, expected Fizz", FizzBuzz(3))
	}
}

func TestBuzzNumber(t *testing.T) {
	if FizzBuzz(5) != "Buzz" {
		t.Errorf("incorrect response: got %s, expected Buzz", FizzBuzz(5))
	}
}
