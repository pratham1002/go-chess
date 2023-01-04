package chess

import (
	"errors"
	"reflect"
)

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

func (c *Chess) Move(from Position, to Position) error {
	// sanity checks - from and to are on the board and not the same
	if !from.IsValid() || !to.IsValid() {
		return errors.New("invalid position")
	}

	// query the board for the piece at the from position
	from_piece := c.Board.GetPieceAt(from)

	if from_piece == nil {
		return errors.New("no piece at from position")
	}

	// check if the piece is of the same color as the current turn
	if from_piece.getColor() != c.turn {
		return errors.New("not your turn")
	}

	to_piece := c.Board.GetPieceAt(to)

	// check if the move is legal
	if from == to ||
		(to_piece != nil && from_piece.getColor() == to_piece.getColor()) ||
		!from_piece.IsLegalMove(from, to, *c.Board) {
		return errors.New("illegal move")
	}

	// if the to piece is a king, the game is over
	if to_piece != nil && reflect.TypeOf(to_piece) == reflect.TypeOf(&King{}) {
		// TODO: Notify Game Over
		return nil
	}

	// move the piece
	c.Board.SpotAt(from).Piece = nil
	c.Board.SpotAt(to).Piece = from_piece

	// check if the game is in check, revert the move if it is
	if c.Board.IsInCheck(c.turn) {
		c.Board.SpotAt(from).Piece = from_piece
		c.Board.SpotAt(to).Piece = to_piece
		return errors.New("illegal move")
	}

	// check if the game is in checkmate
	if c.Board.IsInCheckmate(c.turn) {
		// TODO: Notify checkmate
		return nil
	}

	// check if the game is in stalemate
	if c.Board.IsInStalemate(c.turn) {
		// TODO: Notify stalemate
		return nil
	}

	// change the turn
	c.turn = c.turn.Opposite()

	return nil
}

func (c *Chess) String() string {
	return "Turn: " + c.turn.String() + " " + c.Board.String()
}
