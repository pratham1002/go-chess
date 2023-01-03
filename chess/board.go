package chess

type Board struct {
	Pieces [8][8]Piece
}

func NewBoard() *Board {
	// Create the board
	board := &Board{}

	// Create the pieces
	var piece Piece = NewKing(White, Position{0, 4})
	board.Pieces[0][4] = piece

	piece = NewKing(Black, Position{7, 4})
	board.Pieces[7][4] = piece

	piece = NewQueen(White, Position{0, 3})
	board.Pieces[0][3] = piece

	piece = NewQueen(Black, Position{7, 3})
	board.Pieces[7][3] = piece

	piece = NewRook(White, Position{0, 0})
	board.Pieces[0][0] = piece
	piece = NewRook(White, Position{0, 7})
	board.Pieces[0][7] = piece

	piece = NewRook(Black, Position{7, 0})
	board.Pieces[7][0] = piece
	piece = NewRook(Black, Position{7, 7})
	board.Pieces[7][7] = piece

	piece = NewKnight(White, Position{0, 1})
	board.Pieces[0][1] = piece
	piece = NewKnight(White, Position{0, 6})
	board.Pieces[0][6] = piece

	piece = NewKnight(Black, Position{7, 1})
	board.Pieces[7][1] = piece
	piece = NewKnight(Black, Position{7, 6})
	board.Pieces[7][6] = piece

	piece = NewBishop(White, Position{0, 2})
	board.Pieces[0][2] = piece
	piece = NewBishop(White, Position{0, 5})
	board.Pieces[0][5] = piece

	piece = NewBishop(Black, Position{7, 2})
	board.Pieces[7][2] = piece
	piece = NewBishop(Black, Position{7, 5})
	board.Pieces[7][5] = piece

	for i := 0; i < 8; i++ {
		piece = NewPawn(White, Position{1, i})
		board.Pieces[1][i] = piece
	}

	for i := 0; i < 8; i++ {
		piece = NewPawn(Black, Position{6, i})
		board.Pieces[6][i] = piece
	}

	return board
}

func (b *Board) String() string {
	var s string
	s = "\n+---+---+---+---+---+---+---+---+\n|"
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			if b.Pieces[i][j] == nil {
				s += "   "
			} else {
				s += " " + b.Pieces[i][j].String() + " "
			}
			s += "|"
		}
		s += "\n+---+---+---+---+---+---+---+---+\n"

		if i < 7 {
			s += "|"
		}
	}
	return s
}
