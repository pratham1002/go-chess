package chess

type King struct {
	Color   Color
	IsMoved bool
}

func NewKing(color Color) *King {
	return &King{
		Color:   color,
		IsMoved: false,
	}
}

func (k *King) getColor() Color {
	return k.Color
}

func (k *King) IsLegalMove(from Position, to Position, b Board) bool {
	if from.X == to.X+1 && from.Y == to.Y {
		return true
	} else if from.X == to.X+1 && from.Y == to.Y+1 {
		return true
	} else if from.X == to.X+1 && from.Y == to.Y-1 {
		return true
	} else if from.X == to.X-1 && from.Y == to.Y {
		return true
	} else if from.X == to.X-1 && from.Y == to.Y+1 {
		return true
	} else if from.X == to.X-1 && from.Y == to.Y-1 {
		return true
	} else if from.X == to.X && from.Y == to.Y+1 {
		return true
	} else if from.X == to.X && from.Y == to.Y-1 {
		return true
	} else {
		return false
	}
}

func (k *King) String() string {
	if k.Color == White {
		return "♔"
	} else {
		return "♚"
	}
}
