package chess

type Knight struct {
	Color Color
}

func NewKnight(color Color) *Knight {
	return &Knight{
		Color: color,
	}
}

func (k *Knight) getColor() Color {
	return k.Color
}

func (k *Knight) IsLegalMove(from Position, to Position, b Board) bool {
	if from.X == to.X+1 && from.Y == to.Y+2 {
		return true
	} else if from.X == to.X+1 && from.Y == to.Y-2 {
		return true
	} else if from.X == to.X-1 && from.Y == to.Y+2 {
		return true
	} else if from.X == to.X-1 && from.Y == to.Y-2 {
		return true
	} else if from.X == to.X+2 && from.Y == to.Y+1 {
		return true
	} else if from.X == to.X+2 && from.Y == to.Y-1 {
		return true
	} else if from.X == to.X-2 && from.Y == to.Y+1 {
		return true
	} else if from.X == to.X-2 && from.Y == to.Y-1 {
		return true
	} else {
		return false
	}
}

func (k *Knight) String() string {
	if k.Color == White {
		return "♘"
	} else {
		return "♞"
	}
}
