package fizzbuzz

import "testing"

func TestFizzBuzzTable(t *testing.T) {
	tests := []struct {
		n   int
		msg string
	}{
		{1, "1"}, {2, "2"},
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"}, {8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{15, "FizzBuzz?"},
	}

	for _, test := range tests {
		if m := FizzBuzz(test.n); m != test.msg {
			t.Errorf("incorrect message for %d: got %s, expected %s", test.n, m, test.msg)
		}
	}
}
