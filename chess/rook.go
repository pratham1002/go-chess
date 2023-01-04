package chess

type Rook struct {
	Color   Color
	IsMoved bool
}

func NewRook(color Color) *Rook {
	return &Rook{
		Color:   color,
		IsMoved: false,
	}
}

func (r *Rook) getColor() Color {
	return r.Color
}

func (r *Rook) IsLegalMove(from Position, to Position, b Board) bool {
	return vertical(from, to, b) || horizontal(from, to, b)
}

func (r *Rook) String() string {
	if r.Color == White {
		return "♖"
	} else {
		return "♜"
	}
}
