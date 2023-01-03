package chess

type Queen struct {
	color Color
	pos   Position
}

func NewQueen(color Color, pos Position) *Queen {
	return &Queen{
		color: color,
		pos:   pos,
	}
}

func (q *Queen) IsLegalMove(pos Position) bool {
	if !pos.isValid() {
		return false
	}

	if q.pos.Equals(&pos) {
		return false
	}

	// check if the queen is moving in a straight line, with no pieces in the way
	if q.pos.X == pos.X {
		return true
	}

	if q.pos.Y == pos.Y {
		return true
	}

	if q.pos.X-q.pos.Y == pos.X-pos.Y {
		return true
	}

	// TODO: check if any piece is in the way

	return false
}

func (q *Queen) Move(pos Position) {
	q.pos = pos
}

func (q *Queen) String() string {
	if q.color == White {
		return "♕"
	} else {
		return "♛"
	}
}
