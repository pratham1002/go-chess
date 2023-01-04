package chess

type Bishop struct {
	Color Color
}

func NewBishop(color Color) *Bishop {
	return &Bishop{
		Color: color,
	}
}

func (b *Bishop) getColor() Color {
	return b.Color
}

func (b *Bishop) IsLegalMove(from Position, to Position, board Board) bool {
	return diagonal(from, to, board)
}

func (b *Bishop) String() string {
	if b.Color == White {
		return "♗"
	} else {
		return "♝"
	}
}
