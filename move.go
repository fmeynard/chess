package main

type MovesMap map[int][]int

var (
	whitePawnNormalMovesMap  = make(MovesMap)
	blackPawnNormalMovesMap  = make(MovesMap)
	whitePawnCaptureMovesMap = make(MovesMap)
	blackPawnCaptureMovesMap = make(MovesMap)
	knightMoves              = make(MovesMap)
	diagonalMoves            = make(map[int]MovesMap) // [direction][cellIdx] -> possiblesCellsIdx[]
	linearMoves              = make(map[int]MovesMap) // [direction][cellIdx] -> possiblesCellsIdx[]
)

const ( // Move directions
	North = iota
	South
	West
	East
	NorthWest
	NorthEast
	SouthWest
	SouthEast
)

func generateBaseMoves() {
	generatePawnMoves()
	generateKnightMoves()
	generateDiagonalSliderMoves()
	generateLinearSliderMoves()
}

func generateDiagonalSliderMoves() {
	diagonalMoves[SouthEast] = make(MovesMap)
	diagonalMoves[SouthWest] = make(MovesMap)
	diagonalMoves[NorthEast] = make(MovesMap)
	diagonalMoves[NorthWest] = make(MovesMap)

	for i := 0; i < 64; i++ {
		r, c := cellIdxToCoordinates(i)

		inc := 1
		for y := 1; r+y < 8; y++ {
			if c+inc < 8 {
				diagonalMoves[SouthEast][i] = append(diagonalMoves[SouthEast][i], i+(y*8)+inc)
			}

			if c-inc >= 0 {
				diagonalMoves[SouthWest][i] = append(diagonalMoves[SouthWest][i], i+(y*8)-inc)
			}

			inc++
		}

		inc = 1
		for y := 1; r-y >= 0; y++ {
			if c+inc < 8 {
				diagonalMoves[NorthEast][i] = append(diagonalMoves[NorthEast][i], i-(y*8)+inc)
			}

			if c-inc >= 0 {
				diagonalMoves[NorthWest][i] = append(diagonalMoves[NorthWest][i], i-(y*8)-inc)
			}

			inc++
		}
	}
}

func generateLinearSliderMoves() {
	linearMoves[North] = make(MovesMap)
	linearMoves[South] = make(MovesMap)
	linearMoves[West] = make(MovesMap)
	linearMoves[East] = make(MovesMap)

	for i := 0; i < 64; i++ {
		r, c := cellIdxToCoordinates(i)

		for y := 1; y+r < 8; y++ {
			linearMoves[South][i] = append(linearMoves[South][i], i+y*8)
		}

		for y := 1; r-y >= 0; y++ {
			linearMoves[North][i] = append(linearMoves[North][i], i-y*8)
		}

		for x := 1; x+c < 8; x++ {
			linearMoves[East][i] = append(linearMoves[East][i], i+x)
		}

		for x := 1; c-x >= 0; x++ {
			linearMoves[West][i] = append(linearMoves[West][i], i-x)
		}
	}
}

func generateKnightMoves() {
	for i := 0; i < 64; i++ {
		r, c := cellIdxToCoordinates(i)

		// top left
		if r-2 >= 0 && c-1 >= 0 {
			knightMoves[i] = append(knightMoves[i], i-17)
		}
		// top right
		if r-2 >= 0 && c+1 <= 7 {
			knightMoves[i] = append(knightMoves[i], i-15)
		}
		// bottom left
		if r+2 <= 7 && c-1 >= 0 {
			knightMoves[i] = append(knightMoves[i], i+17)
		}
		// bottom right
		if r+2 <= 7 && c+1 <= 7 {
			knightMoves[i] = append(knightMoves[i], i+15)
		}

		// left top
		if c-2 >= 0 && r-1 >= 0 {
			knightMoves[i] = append(knightMoves[i], i-10)
		}
		// left bottom
		if c-2 >= 0 && r+1 <= 7 {
			knightMoves[i] = append(knightMoves[i], i+6)
		}
		// right top
		if c+2 <= 7 && r-1 >= 0 {
			knightMoves[i] = append(knightMoves[i], i-6)
		}
		if c+2 <= 7 && r+1 <= 7 {
			knightMoves[i] = append(knightMoves[i], i+10)
		}
	}
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

func pawnPossiblesMoves(pos *Position, pieceToMoveIdx int, pieceColor int) []int {
	isWhite := pieceColor == PieceColorWhite
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
		if pos.board[Cell(pMove)] != NoPiece && pos.board[Cell(pMove)].toColor() != pieceColor {
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

func rookPossibleMoves(pos *Position, pieceToMoveIdx int, pieceColor int) []int {
	return linearSliderPossibleMoves(pos, pieceToMoveIdx, pieceColor)
}

func knightPossiblesMove(pos *Position, pieceToMoveIdx int, pieceColor int) []int {
	legalMoves := []int{}
	for _, pMove := range knightMoves[pieceToMoveIdx] {
		if pos.board[Cell(pMove)] == NoPiece || pos.board[Cell(pMove)].toColor() != pieceColor {
			legalMoves = append(legalMoves, pMove)
		}
	}

	return legalMoves
}

func bishopPossibleMoves(pos *Position, pieceToMoveIdx int, pieceColor int) []int {
	return diagonalSliderPossibleMoves(pos, pieceToMoveIdx, pieceColor)
}

func queenPossibleMoves(pos *Position, pieceToMoveIdx int, pieceColor int) []int {
	return sliderPossibleMoves(
		pos,
		pieceToMoveIdx,
		pieceColor,
		[]int{North, South, West, East, SouthEast, SouthWest, NorthEast, NorthWest},
	)
}

func linearSliderPossibleMoves(pos *Position, pieceToMoveIdx int, pieceColor int) []int {
	return sliderPossibleMoves(
		pos,
		pieceToMoveIdx,
		pieceColor,
		[]int{North, South, West, East},
	)
}

func diagonalSliderPossibleMoves(pos *Position, pieceToMoveIdx int, pieceColor int) []int {
	return sliderPossibleMoves(
		pos,
		pieceToMoveIdx,
		pieceColor,
		[]int{SouthEast, SouthWest, NorthEast, NorthWest},
	)
}

func sliderPossibleMoves(pos *Position, pieceToMoveIdx int, pieceColor int, possibleDirections []int) []int {
	legalMoves := []int{}

	for _, direction := range possibleDirections {
		var directionMoves []int
		if direction < 4 {
			directionMoves = linearMoves[direction][pieceToMoveIdx]
		} else {
			directionMoves = diagonalMoves[direction][pieceToMoveIdx]
		}

		for _, pMove := range directionMoves {
			if pos.board[Cell(pMove)].toColor() == pieceColor {
				break
			}
			legalMoves = append(legalMoves, pMove)
			if pos.board[Cell(pMove)] != NoPiece {
				break
			}
		}
	}

	return legalMoves
}

func (piece Piece) possibleMoves(pieceIdx int, pos *Position) []int {
	switch piece.toPieceType() {
	case PieceTypePawn:
		return pawnPossiblesMoves(
			pos,
			pieceIdx,
			piece.toColor(),
		)
	case PieceTypeKnight:
		return knightPossiblesMove(pos, pieceIdx, piece.toColor())
	case PieceTypeBishop:
		return bishopPossibleMoves(pos, pieceIdx, piece.toColor())
	case PieceTypeRook:
		return rookPossibleMoves(pos, pieceIdx, piece.toColor())
	case PieceTypeQueen:
		return queenPossibleMoves(pos, pieceIdx, piece.toColor())
	}

	return []int{}
}
