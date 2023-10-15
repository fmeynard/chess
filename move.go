package main

var (
	whitePawnNormalMovesMap  = make(map[int][]int)
	blackPawnNormalMovesMap  = make(map[int][]int)
	whitePawnCaptureMovesMap = make(map[int][]int)
	blackPawnCaptureMovesMap = make(map[int][]int)
)

func generateBaseMoves() {
	generatePawnMoves()
}

func generatePawnMoves() {
	for i := 0; i < 64; i++ {
		r, c := cellIdxToCoordinates(i)

		if r == 1 {
			whitePawnNormalMovesMap[i] = append(whitePawnNormalMovesMap[i], i-8)
			blackPawnNormalMovesMap[i] = append(blackPawnNormalMovesMap[i], i+8)
		} else if r == 6 {
			whitePawnNormalMovesMap[i] = append(whitePawnNormalMovesMap[i], i-8)
			blackPawnNormalMovesMap[i] = append(blackPawnNormalMovesMap[i], i+8)
		} else {
			if r != 7 {
				blackPawnNormalMovesMap[i] = append(blackPawnNormalMovesMap[i], i+8)
			}
			if r != 0 {
				whitePawnNormalMovesMap[i] = append(whitePawnNormalMovesMap[i], i-8)
			}
		}

		if c == 0 {
			if r != 0 {
				whitePawnCaptureMovesMap[i] = append(whitePawnCaptureMovesMap[i], i-7)
			}
			if r != 7 {
				blackPawnCaptureMovesMap[i] = append(blackPawnCaptureMovesMap[i], i+9)
			}
		} else if c == 7 {
			if r != 0 {
				whitePawnCaptureMovesMap[i] = append(whitePawnCaptureMovesMap[i], i-9)
			}
			if r != 7 {
				blackPawnCaptureMovesMap[i] = append(blackPawnCaptureMovesMap[i], i+7)
			}
		} else {
			if r != 0 {
				whitePawnCaptureMovesMap[i] = append(whitePawnCaptureMovesMap[i], i-7, i-9)
			}
			if r != 7 {
				blackPawnCaptureMovesMap[i] = append(blackPawnCaptureMovesMap[i], i+7, i+9)
			}
		}
	}
}

func pawnPossiblesMoves(pos *Position, pieceToMoveIdx int, isWhite bool) []int {
	normalMovesMap := blackPawnNormalMovesMap
	capturesMovesMap := blackPawnCaptureMovesMap
	if isWhite {
		normalMovesMap = whitePawnNormalMovesMap
		capturesMovesMap = whitePawnCaptureMovesMap
	}

	moves := []int{}
	for _, pMove := range normalMovesMap[pieceToMoveIdx] {
		if pos.board[Cell(pMove)] == NoPiece {
			moves = append(moves, pMove)
		}
	}

	r, _ := cellIdxToCoordinates(pieceToMoveIdx)
	if r == 6 &&
		isWhite &&
		pos.board[Cell(pieceToMoveIdx-16)] == NoPiece &&
		pos.board[Cell(pieceToMoveIdx-8)] == NoPiece {
		moves = append(moves, pieceToMoveIdx-16)
	} else if r == 1 && !isWhite &&
		pos.board[Cell(pieceToMoveIdx+16)] == NoPiece &&
		pos.board[Cell(pieceToMoveIdx+8)] == NoPiece {
		moves = append(moves, pieceToMoveIdx+16)
	}

	for _, pMove := range capturesMovesMap[pieceToMoveIdx] {
		if pos.board[Cell(pMove)] != NoPiece {
			moves = append(moves, pMove)
		}
	}

	// if pos.enPassantIdx != nil {
	// 	if enPassantMap[pos.enPassantIdx] == pieceToMoveIdx {
	// 		moves = append(moves, enPassantIdx)
	// 	}
	// }

	return moves
}

func (piece Piece) possibleMoves(pieceIdx int, pos *Position) []int {
	if piece.toPieceType() == PieceTypePawn {
		return pawnPossiblesMoves(
			pos,
			pieceIdx,
			piece.toColor() == PieceColorWhite,
		)
	}

	return []int{}
}
