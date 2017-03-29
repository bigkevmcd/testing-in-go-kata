package marsrover

import "testing"

func TestNewRobot(t *testing.T) {
	l := Location{1, 1, "E"}
	r := New(5, 3, l)

	if r.XRange != 5 {
		t.Errorf("unexpected robot XRange: got %d, expected %d", r.XRange, 5)
	}

	if r.YRange != 3 {
		t.Errorf("unexpected robot XRange: got %d, expected %d", r.YRange, 3)
	}
	if r.Position != l {
		t.Errorf("unexpected robot Position: got %+v, expected %+v", r.Position, l)
	}
}

func TestMoveForward(t *testing.T) {
	forwardTests := []struct {
		facing string
		end    Location
	}{
		{"N", Location{1, 2, "N"}},
		{"E", Location{2, 1, "E"}},
		{"W", Location{0, 1, "W"}},
		{"S", Location{1, 0, "S"}},
	}

	for i, tt := range forwardTests {
		l := Location{1, 1, tt.facing}
		r := New(5, 3, l)

		r.MoveForward()

		if r.Position != tt.end {
			t.Errorf("failed test %d unexpected position: got %+v, expected %+v", i, r.Position, tt.end)
		}
	}
}

func TestTurnRight(t *testing.T) {
	forwardTests := []struct {
		facing string
		end    Location
	}{
		{"N", Location{1, 1, "E"}},
		{"E", Location{1, 1, "S"}},
		{"W", Location{1, 1, "N"}},
		{"S", Location{1, 1, "W"}},
	}

	for i, tt := range forwardTests {
		l := Location{1, 1, tt.facing}
		r := New(5, 3, l)

		r.TurnRight()

		if r.Position != tt.end {
			t.Errorf("failed test %d unexpected position: got %+v, expected %+v", i, r.Position, tt.end)
		}
	}
}

var sequences = []struct {
	commands []string
	expected Location
}{
	{[]string{"F"}, Location{2, 1, "E"}},
	{[]string{"F", "R"}, Location{2, 1, "S"}},
	{[]string{"R", "R"}, Location{1, 1, "W"}},
	{[]string{"F", "F"}, Location{2, 2, "E"}},
}

func TestExecuteSequence(t *testing.T) {
	for i, tt := range sequences {
		l := Location{1, 1, "E"}
		r := New(5, 3, l)
		new := r.Execute(tt.commands...)
		if new != tt.expected {
			t.Errorf("test %d incorrect location: got %+v, expected %+v", i, new, tt.expected)
		}
	}
}
