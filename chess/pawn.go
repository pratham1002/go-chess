package chess

type Pawn struct {
	Color   Color
	IsMoved bool
}

func NewPawn(color Color, pos Position) *Pawn {
	return &Pawn{
		Color:   color,
		IsMoved: false,
	}
}

func (p *Pawn) getColor() Color {
	return p.Color
}

func (p *Pawn) IsLegalMove(from Position, to Position, b Board) bool {
	// check if the move is one square in any direction
	if p.Color == White {
		if from.X == to.X && from.Y == to.Y-1 {
			return true
		}
	} else {
		if from.X == to.X && from.Y == to.Y+1 {
			return true
		}
	}

	// if the pawn has not moved, it can move two squares
	if !p.IsMoved {
		if p.Color == White {
			if from.X == to.X && from.Y == to.Y-2 {
				return true
			}
		} else {
			if from.X == to.X && from.Y == to.Y+2 {
				return true
			}
		}
	}

	// TODO: check if any piece is in the way and if the move is diagonal
	if b.SpotAt(to).Piece != nil && b.SpotAt(to).Piece.getColor() != p.Color {
		if p.Color == White {
			if from.X == to.X-1 && from.Y == to.Y-1 {
				return true
			}
			if from.X == to.X+1 && from.Y == to.Y-1 {
				return true
			}
		} else {
			if from.X == to.X-1 && from.Y == to.Y+1 {
				return true
			}
			if from.X == to.X+1 && from.Y == to.Y+1 {
				return true
			}
		}
	}

	return false
}

func (p *Pawn) String() string {
	if p.Color == White {
		return "♙"
	} else {
		return "♟"
	}
}
