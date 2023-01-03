package chess

type Pawn struct {
	Color   Color
	Pos     Position
	isMoved bool
}

func NewPawn(color Color, pos Position) *Pawn {
	return &Pawn{
		Color:   color,
		Pos:     pos,
		isMoved: false,
	}
}

func (p *Pawn) IsMoved() bool {
	return p.isMoved
}

func (p *Pawn) IsLegalMove(pos Position) bool {
	if !pos.isValid() {
		return false
	}

	if p.Pos.Equals(&pos) {
		return false
	}

	// check if the move is one square in any direction
	if p.Color == White {
		if p.Pos.X == pos.X && p.Pos.Y == pos.Y-1 {
			return true
		}
	} else {
		if p.Pos.X == pos.X && p.Pos.Y == pos.Y+1 {
			return true
		}
	}

	// if the pawn has not moved, it can move two squares
	if !p.isMoved {
		if p.Color == White {
			if p.Pos.X == pos.X && p.Pos.Y == pos.Y-2 {
				return true
			}
		} else {
			if p.Pos.X == pos.X && p.Pos.Y == pos.Y+2 {
				return true
			}
		}
	}

	// TODO: check if any piece is in the way and if the move is diagonal

	return false
}

func (p *Pawn) Move(pos Position) {
	p.Pos = pos
	p.isMoved = true
}

func (p *Pawn) String() string {
	if p.Color == White {
		return "♙"
	} else {
		return "♟"
	}
}
