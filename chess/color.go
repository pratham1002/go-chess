package chess

type Color int

const (
	White Color = iota
	Black
)

func (c Color) Opposite() Color {
	if c == White {
		return Black
	}
	return White
}

func (c Color) Equals(other Color) bool {
	return c == other
}

func (c Color) String() string {
	switch c {
	case Black:
		return "Black"
	case White:
		return "White"
	}
	return "Invalid color"
}
