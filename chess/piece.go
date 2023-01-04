package chess

type Piece interface {
	getColor() Color
	IsLegalMove(from Position, to Position, b Board) bool
	String() string
}
