package chess

type Spot struct {
	Position
	Piece
}

func NewSpot(p Position) *Spot {
	return &Spot{Position: p}
}

func (s *Spot) String() string {
	if s.Piece == nil {
		return "   "
	} else {
		return " " + s.Piece.String() + " "
	}
}
