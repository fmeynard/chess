//go:generate fyne bundle -o piece-bundle.go pieces

package main

import "fyne.io/fyne/v2"

type Piece int8

const (
	NoPiece Piece = iota
	BlackRook
	BlackKnight
	BlackBishop
	BlackQueen
	BlackKing
	BlackPawn
	WhiteRook
	WhiteKnight
	WhiteBishop
	WhiteQueen
	WhiteKing
	WhitePawn
)

func NewPiece(fenId string) Piece {
	switch fenId {
	case "r":
		return WhiteRook
	case "n":
		return WhiteKnight
	case "b":
		return WhiteBishop
	case "q":
		return WhiteQueen
	case "k":
		return WhiteKing
	case "p":
		return WhitePawn
	case "R":
		return BlackRook
	case "N":
		return BlackKnight
	case "B":
		return BlackBishop
	case "Q":
		return BlackQueen
	case "K":
		return BlackKing
	case "P":
		return BlackPawn
	}

	return NoPiece
}

func (p Piece) toSvg() *fyne.StaticResource {
	switch p {
	case WhiteRook:
		return resourceWhiteRookSvg
	case WhiteKnight:
		return resourceWhiteKnightSvg
	case WhiteBishop:
		return resourceWhiteBishopSvg
	case WhiteQueen:
		return resourceWhiteQueenSvg
	case WhiteKing:
		return resourceWhiteKingSvg
	case WhitePawn:
		return resourceWhitePawnSvg
	case BlackRook:
		return resourceBlackRookSvg
	case BlackKnight:
		return resourceBlackKnightSvg
	case BlackBishop:
		return resourceBlackBishopSvg
	case BlackQueen:
		return resourceBlackQueenSvg
	case BlackKing:
		return resourceBlackKingSvg
	case BlackPawn:
		return resourceBlackPawnSvg
	}

	panic("piece to ressource mapping not defined")
}
