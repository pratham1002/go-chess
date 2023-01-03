package chess

type Rook struct {
	Color   Color
	Pos     Position
	isMoved bool
}

func NewRook(color Color, pos Position) *Rook {
	return &Rook{
		Color:   color,
		Pos:     pos,
		isMoved: false,
	}
}

func (r *Rook) IsMoved() bool {
	return r.isMoved
}

func (r *Rook) IsLegalMove(pos Position) bool {
	if !pos.isValid() {
		return false
	}

	if r.Pos.Equals(&pos) {
		return false
	}

	// check if the move either in the same row or column
	if r.Pos.X == pos.X {
		return true
	}

	if r.Pos.Y == pos.Y {
		return true
	}

	// TODO: check if there is a piece in the way

	return false
}

func (r *Rook) Move(pos Position) {
	r.Pos = pos
	r.isMoved = true
}

func (r *Rook) String() string {
	if r.Color == White {
		return "♖"
	} else {
		return "♜"
	}
}
