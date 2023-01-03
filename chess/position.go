package chess

type Position struct {
	X int
	Y int
}

func NewPosition(x, y int) *Position {
	return &Position{
		X: x,
		Y: y,
	}
}

func (p *Position) isValid() bool {
	return p.X >= 0 && p.X < 8 && p.Y >= 0 && p.Y < 8
}

func (p *Position) Equals(other *Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func (p *Position) String() string {
	return string(rune(p.X+'A')) + string(rune(p.Y+'1'))
}
