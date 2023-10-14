package main

import "strconv"

type Cell int

func NewCell(row int, col int) Cell {
	return Cell(row*8 + col)
}

func (cell Cell) toNotation() string {
	letter := ""
	switch int(cell) % 8 {
	case 0:
		letter += "a"
	case 1:
		letter += "b"
	case 2:
		letter += "c"
	case 3:
		letter += "d"
	case 4:
		letter += "e"
	case 5:
		letter += "f"
	case 6:
		letter += "g"
	case 7:
		letter += "h"
	}

	return letter + strconv.Itoa(8-int(cell)/8)
}

func cellIdxToCoordinates(cellIdx int) (int, int) {
	return cellIdx / 8, cellIdx % 8
}
