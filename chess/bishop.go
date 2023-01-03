package chess

type Bishop struct {
	Color Color
	Pos   Position
}

func NewBishop(color Color, pos Position) *Bishop {
	return &Bishop{
		Color: color,
		Pos:   pos,
	}
}

func (b *Bishop) IsLegalMove(pos Position) bool {
	if !pos.isValid() {
		return false
	}

	if b.Pos.Equals(&pos) {
		return false
	}

	// check if the bishop is moving diagonally, with no pieces in the way
	if b.Pos.X-b.Pos.Y == pos.X-pos.Y {
		return true
	}

	// TODO: check if any piece is in the way

	return false
}

func (b *Bishop) Move(pos Position) {
	b.Pos = pos
}

func (b *Bishop) String() string {
	if b.Color == White {
		return "♗"
	} else {
		return "♝"
	}
}
