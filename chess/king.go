package chess

type King struct {
	Color   Color
	Pos     Position
	isMoved bool
}

func NewKing(color Color, pos Position) *King {
	return &King{
		Color:   color,
		Pos:     pos,
		isMoved: false,
	}
}

func (k *King) IsMoved() bool {
	return k.isMoved
}

func (k *King) IsLegalMove(pos Position) bool {
	if !pos.isValid() {
		return false
	}

	if k.Pos.Equals(&pos) {
		return false
	}

	// check if the move is one square in any direction
	if k.Pos.X == pos.X && k.Pos.Y == pos.Y-1 {
		return true
	}

	if k.Pos.X == pos.X && k.Pos.Y == pos.Y+1 {
		return true
	}

	if k.Pos.X == pos.X-1 && k.Pos.Y == pos.Y {
		return true
	}

	if k.Pos.X == pos.X+1 && k.Pos.Y == pos.Y {
		return true
	}

	if k.Pos.X == pos.X-1 && k.Pos.Y == pos.Y-1 {
		return true
	}

	if k.Pos.X == pos.X+1 && k.Pos.Y == pos.Y-1 {
		return true
	}

	if k.Pos.X == pos.X-1 && k.Pos.Y == pos.Y+1 {
		return true
	}

	if k.Pos.X == pos.X+1 && k.Pos.Y == pos.Y+1 {
		return true
	}

	return false
}

func (k *King) Move(pos Position) {
	k.Pos = pos
	k.isMoved = true
}

func (k *King) String() string {
	if k.Color == White {
		return "♔"
	} else {
		return "♚"
	}
}
