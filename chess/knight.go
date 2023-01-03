package chess

type Knight struct {
	Color Color
	Pos   Position
}

func NewKnight(color Color, pos Position) *Knight {
	return &Knight{
		Color: color,
		Pos:   pos,
	}
}

func (k *Knight) IsLegalMove(pos Position) bool {
	if !pos.isValid() {
		return false
	}

	if k.Pos.Equals(&pos) {
		return false
	}

	// check if the move is two squares in one direction and one square in the other
	if k.Pos.X == pos.X-2 && k.Pos.Y == pos.Y-1 {
		return true
	}

	if k.Pos.X == pos.X-2 && k.Pos.Y == pos.Y+1 {
		return true
	}

	if k.Pos.X == pos.X+2 && k.Pos.Y == pos.Y-1 {
		return true
	}

	if k.Pos.X == pos.X+2 && k.Pos.Y == pos.Y+1 {
		return true
	}

	if k.Pos.X == pos.X-1 && k.Pos.Y == pos.Y-2 {
		return true
	}

	if k.Pos.X == pos.X-1 && k.Pos.Y == pos.Y+2 {
		return true
	}

	if k.Pos.X == pos.X+1 && k.Pos.Y == pos.Y-2 {
		return true
	}

	if k.Pos.X == pos.X+1 && k.Pos.Y == pos.Y+2 {
		return true
	}

	return false
}

func (k *Knight) Move(pos Position) {
	k.Pos = pos
}

func (k *Knight) String() string {
	if k.Color == White {
		return "♘"
	} else {
		return "♞"
	}
}
