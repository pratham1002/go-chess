package chess

type Piece interface {
	IsLegalMove(pos Position) bool
	Move(pos Position)
	String() string
}
