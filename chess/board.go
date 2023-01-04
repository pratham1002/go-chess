package chess

type Board struct {
	Spots [8][8]*Spot
}

func NewBoard() *Board {
	// Create the board
	board := &Board{
		Spots: [8][8]*Spot{},
	}

	// Create the spots
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			board.Spots[i][j] = NewSpot(Position{i, j})
		}
	}

	// Create the pieces
	var piece Piece

	board.Spots[0][4].Piece = NewKing(White)

	board.Spots[7][4].Piece = NewKing(Black)

	board.Spots[0][3].Piece = NewQueen(White)

	board.Spots[7][3].Piece = NewQueen(Black)

	board.Spots[0][2].Piece = NewBishop(White)

	board.Spots[0][5].Piece = NewBishop(White)

	board.Spots[7][2].Piece = NewBishop(Black)

	board.Spots[7][5].Piece = NewBishop(Black)

	board.Spots[0][1].Piece = NewKnight(White)

	board.Spots[0][6].Piece = NewKnight(White)

	board.Spots[7][1].Piece = NewKnight(Black)

	board.Spots[7][6].Piece = NewKnight(Black)

	board.Spots[0][0].Piece = NewRook(White)

	board.Spots[0][7].Piece = NewRook(White)

	board.Spots[7][0].Piece = NewRook(Black)

	board.Spots[7][7].Piece = NewRook(Black)

	for i := 0; i < 8; i++ {
		piece = NewPawn(White, Position{1, i})
		board.Spots[1][i].Piece = piece

		piece = NewPawn(Black, Position{6, i})
		board.Spots[6][i].Piece = piece
	}
	return board
}

func (b *Board) SpotAt(p Position) *Spot {
	return b.Spots[p.X][p.Y]
}

func (b *Board) GetPieceAt(p Position) Piece {
	return b.SpotAt(p).Piece
}

func (b *Board) IsInCheck(color Color) bool {
	// TODO: Implement
	return false
}

func (b *Board) IsInCheckmate(color Color) bool {
	// TODO: Implement
	return false
}

func (b *Board) IsInStalemate(color Color) bool {
	// TODO: Implement
	return false
}

func (b *Board) String() string {
	var s string
	s = "\n+---+---+---+---+---+---+---+---+\n|"
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			s += b.SpotAt(Position{i, j}).String() + "|"
		}
		s += "\n+---+---+---+---+---+---+---+---+\n"

		if i < 7 {
			s += "|"
		}
	}
	return s
}
