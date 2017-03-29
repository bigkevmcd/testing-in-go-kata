package marsrover

type Location struct {
	X, Y   int64
	Facing string
}

type Robot struct {
	XRange   int64
	YRange   int64
	Position Location
}

func (r *Robot) MoveForward() {
	switch r.Position.Facing {
	case "N":
		r.Position.Y++
	case "E":
		r.Position.X++
	case "W":
		r.Position.X--
	case "S":
		r.Position.Y--
	}
}

func (r *Robot) TurnRight() {
	switch r.Position.Facing {
	case "N":
		r.Position.Facing = "E"
	case "E":
		r.Position.Facing = "S"
	case "W":
		r.Position.Facing = "N"
	case "S":
		r.Position.Facing = "W"
	}
}

func (r *Robot) Execute(c ...string) Location {
	for _, command := range c {
		switch command {
		case "F":
			r.MoveForward()
		}
	}
	return r.Position
}

func New(x, y int64, p Location) *Robot {
	return &Robot{XRange: x, YRange: y, Position: p}
}
