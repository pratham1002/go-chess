package chess

// define a structure for a chess board which has pointers to pieces
type Board struct {
	// 8x8 array of pointers to pieces
	Pieces [32]*Piece
}
