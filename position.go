package main

import (
	"errors"
	"strconv"
	"strings"
)

type Position struct {
	board map[Cell]Piece
}

// https://en.wikipedia.org/wiki/Forsyth%E2%80%93Edwards_Notation
// ex : rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR
func fenToPosition(fenStr string) (*Position, error) {
	parts := strings.Split(fenStr, " ")
	if len(parts) != 6 {
		return nil, errors.New("invalid fen")
	}

	board := map[Cell]Piece{}
	rows := strings.Split(parts[0], "/")
	for i, row := range rows {
		incJ := 0
		for j, fendPieceId := range row {
			blankCount, err := strconv.Atoi(string(fendPieceId))
			if err == nil {
				for x := 0; x < blankCount; x++ {
					board[NewCell(7-i, 7-j+x+incJ)] = NoPiece
				}
				incJ += blankCount

				continue
			}

			board[NewCell(7-i, 7-j+incJ)] = NewPiece(string(fendPieceId))
			// board[Cell{i, j + incJ}] = NewPiece(string(fendPieceId))
		}
	}

	return &Position{
		board: board,
	}, nil
}
