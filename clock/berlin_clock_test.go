package clock

import (
	"reflect"
	"testing"
	"time"
)

func TestSeconds(t *testing.T) {
	var tests = []struct {
		second   int
		expected string
	}{
		{0, "Y"},
		{1, "O"},
		{59, "O"}}

	for _, tt := range tests {
		if o := Seconds(tt.second); o != tt.expected {
			t.Errorf("Seconds(%d) incorrect, got %s, expected %s", tt.second, o, tt.expected)
		}
	}
}

func TestHours(t *testing.T) {
	var tests = []struct {
		hours    int
		expected []string
	}{
		{0, []string{"OOOO", "OOOO"}},
		{7, []string{"ROOO", "RROO"}},
		{13, []string{"RROO", "RRRO"}},
		{23, []string{"RRRR", "RRRO"}},
	}

	for _, tt := range tests {
		if o := Hours(tt.hours); !reflect.DeepEqual(o, tt.expected) {
			t.Errorf("Hours(%d) incorrect, got %v, expected %v", tt.hours, o, tt.expected)
		}
	}
}

func TestMinutes(t *testing.T) {
	var tests = []struct {
		minutes  int
		expected []string
	}{
		{0, []string{"OOOOOOOOOOO", "OOOO"}},
		{17, []string{"YYROOOOOOOO", "YYOO"}},
		{59, []string{"YYRYYRYYRYY", "YYYY"}},
	}

	for _, tt := range tests {
		if o := Minutes(tt.minutes); !reflect.DeepEqual(o, tt.expected) {
			t.Errorf("Minutes(%d) incorrect, got %v, expected %v", tt.minutes, o, tt.expected)
		}
	}
}

func TestTimes(t *testing.T) {
	var tests = []struct {
		testTime string
		expected []string
	}{
		{"16:37:16", []string{"Y", "RRRO", "ROOO", "YYRYYRYOOOO", "YYOO"}},
	}
	for _, tt := range tests {
		clock, err := time.Parse("15:04:05", tt.testTime)
		if err != nil {
			t.Errorf("failed to parse %s to a valid timestamp: got %s", err)
		}
		if o := BerlinClock(clock); !reflect.DeepEqual(o, tt.expected) {
			t.Errorf("BerlinClock(%s) incorrect, got %v, expected %v", tt.testTime, o, tt.expected)
		}
	}
}
