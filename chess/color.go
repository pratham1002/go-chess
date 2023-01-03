package chess

type Color int

const (
	White Color = iota
	Black
)

func (c Color) String() string {
	switch c {
	case Black:
		return "Black"
	case White:
		return "White"
	}
	return "Invalid color"
}
