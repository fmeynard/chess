package main

func yolo(isWhite bool) (map[int][]int, map[int][]int) {
	normalMovesMap := make(map[int][]int)
	capturesMovesMap := make(map[int][]int)

	for i := 0; i < 64; i++ {
		r, c := cellIdxToCoordinates(i)

		if isWhite {
			if r == 6 {
				normalMovesMap[i] = append(normalMovesMap[i], i-8, i-16)
			} else {
				normalMovesMap[i] = append(normalMovesMap[i], i-8)
			}
		} else {
			if r == 1 {
				normalMovesMap[i] = append(normalMovesMap[i], i+8)
			} else {
				normalMovesMap[i] = append(normalMovesMap[i], i+8, i+16)
			}
		}

		if c == 0 {
			if isWhite {
				capturesMovesMap[i] = append(capturesMovesMap[i], i-7)
			} else {
				capturesMovesMap[i] = append(capturesMovesMap[i], i+9)
			}
		} else if c == 7 {
			if isWhite {
				capturesMovesMap[i] = append(capturesMovesMap[i], i+-9)
			} else {
				capturesMovesMap[i] = append(capturesMovesMap[i], i+7)
			}
		} else {
			if isWhite {
				capturesMovesMap[i] = append(capturesMovesMap[i], i+-7, i+-9)
			} else {
				capturesMovesMap[i] = append(capturesMovesMap[i], i+7, i+9)
			}
		}
	}

	return normalMovesMap, capturesMovesMap
}

func pawnPossiblesMoves(pos *Position, pieceToMoveIdx int, isWhite bool) []int {
	normalMovesMap, capturesMovesMap := yolo(true)

	moves := []int{}
	for _, pMove := range normalMovesMap[pieceToMoveIdx] {
		if pos.board[Cell(pMove)] == NoPiece {
			moves = append(moves, pMove)
		}
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
