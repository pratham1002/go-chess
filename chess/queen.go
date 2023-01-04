package chess

type Queen struct {
	color Color
}

func NewQueen(color Color) *Queen {
	return &Queen{
		color: color,
	}
}

func (q *Queen) getColor() Color {
	return q.color
}

func (q *Queen) IsLegalMove(from Position, to Position, b Board) bool {
	return vertical(from, to, b) || horizontal(from, to, b) || diagonal(from, to, b)
}

func (q *Queen) String() string {
	if q.color == White {
		return "♕"
	} else {
		return "♛"
	}
}
