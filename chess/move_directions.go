package chess

func vertical(from Position, to Position, b Board) bool {
	if to.X != from.X {
		return false
	}

	if from.Y < to.Y {
		for i := from.Y; i < to.Y; i++ {
			if b.SpotAt(Position{X: from.X, Y: i}).Piece != nil {
				return false
			}
		}
	} else {
		for i := from.Y; i > to.Y; i-- {
			if b.SpotAt(Position{X: from.X, Y: i}).Piece != nil {
				return false
			}
		}
	}
	return true
}

func horizontal(from Position, to Position, b Board) bool {
	if to.Y != from.Y {
		return false
	}

	if from.X < to.X {
		for i := from.X; i < to.X; i++ {
			if b.SpotAt(Position{X: i, Y: from.Y}).Piece != nil {
				return false
			}
		}
	} else {
		for i := from.X; i > to.X; i-- {
			if b.SpotAt(Position{X: i, Y: from.Y}).Piece != nil {
				return false
			}
		}
	}
	return true
}

func diagonal(from Position, to Position, b Board) bool {
	if from.X-from.Y != to.X-to.Y {
		return false
	}

	if from.X < to.X && from.Y < to.Y {
		for i := 1; i < to.X-from.X; i++ {
			if b.SpotAt(Position{X: from.X + i, Y: from.Y + i}).Piece != nil {
				return false
			}
		}
		return true
	} else if from.X > to.X && from.Y < to.Y {
		for i := 1; i < from.X-to.X; i++ {
			if b.SpotAt(Position{X: from.X - i, Y: from.Y + i}).Piece != nil {
				return false
			}
		}
		return true
	} else if from.X < to.X && from.Y > to.Y {
		for i := 1; i < to.X-from.X; i++ {
			if b.SpotAt(Position{X: from.X + i, Y: from.Y - i}).Piece != nil {
				return false
			}
		}

		return true
	} else {
		for i := 1; i < from.X-to.X; i++ {
			if b.SpotAt(Position{X: from.X - i, Y: from.Y - i}).Piece != nil {
				return false
			}
		}

		return true
	}
}
