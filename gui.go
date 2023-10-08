package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

type ChessGUI struct {
	window        fyne.Window
	game          *Game
	gameContainer *fyne.Container
}

type Cell int

func NewCell(row int, col int) Cell {
	return Cell(row*8 + col)
}

func cellIdxToCoordinates(cellIdx int) (int, int) {
	return 7 - (cellIdx / 8), cellIdx % 8
}

func NewChessGUI(window fyne.Window, game *Game) *ChessGUI {
	grid := container.NewGridWithColumns(8)

	for i := 0; i < 64; i++ {
		r, c := cellIdxToCoordinates(i)

		cell := canvas.NewRectangle(color.Gray{0x30})
		if r%2 == c%2 {
			cell.FillColor = color.White
		}

		grid.Add(container.NewStack(cell, canvas.NewImageFromResource(nil)))
	}

	window.SetContent(grid)
	window.Resize(fyne.NewSize(500, 500))

	return &ChessGUI{
		gameContainer: grid,
		game:          game,
		window:        window,
	}
}

func (ui *ChessGUI) refreshGrid() {
	for i, cell := range ui.gameContainer.Objects {
		r, c := cellIdxToCoordinates(i)
		img := cell.(*fyne.Container).Objects[1].(*canvas.Image)
		p := ui.game.currentPos.board[NewCell(r, c)]
		if p == NoPiece {
			continue
		}
		img.Resource = p.toSvg()
		img.Refresh()
	}
}
