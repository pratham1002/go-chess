package chess

type Chess struct {
	Board *Board
	turn  Color
}

func NewChess() *Chess {
	return &Chess{
		Board: NewBoard(),
		turn:  White,
	}
}

func (c *Chess) Move(from Position, to Position) {
	// TODO

	// query the board for the piece at the from position

	// check if the piece is of the same color as the current turn

	// check if the move is legal

	// move the piece

	// check if the game is in check, revert the move if it is

	// check if the game is over

	// check if the game is in checkmate

	// change the turn
}

func (c *Chess) String() string {
	return "Turn: " + c.turn.String() + " " + c.Board.String()
}
