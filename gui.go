package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	WhiteSqBgColor    = color.White
	BlackSqBgColor    = color.Gray{0x30}
	PossibleSqBgColor = color.NRGBA{0, 0xff, 0, 0x28}
)

type ChessGUI struct {
	window        fyne.Window
	game          *Game
	gameContainer *fyne.Container
}

type UIPiece struct {
	widget.Icon
	cell Cell
	gui  *ChessGUI
}

func (pos *Position) NewUIPiece(gui *ChessGUI, cell Cell) *UIPiece {
	uiPiece := &UIPiece{cell: cell, gui: gui}
	uiPiece.ExtendBaseWidget(uiPiece)
	currentPiece := pos.board[cell]
	if currentPiece != NoPiece {
		uiPiece.Resource = currentPiece.toSvg()
	}
	return uiPiece
}

func (uiPiece *UIPiece) Tapped(event *fyne.PointEvent) {
	// fmt.Println(uiPiece.cell.toNotation())

	// currentPiece := uiPiece.gui.game.currentPos.board[uiPiece.cell]
	// fmt.Println("Current piece : ", currentPiece, int(uiPiece.cell))

	// if currentPiece.toPieceType() == PieceTypePawn {
	// 	pMoves := pawnPossiblesMoves(uiPiece.gui.game.currentPos, int(uiPiece.cell))
	// 	r, c := cellIdxToCoordinates(int(uiPiece.cell))
	// 	fmt.Println(int(uiPiece.cell), pMoves, r, c)

	// 	// img := cell.(*fyne.Container).Objects[1].(*UIPiece)
	// }
}

var dragStartIdx int = -1

var dragEndPos fyne.Position

func (uiPiece *UIPiece) Dragged(event *fyne.DragEvent) {
	dragStartIdx = int(uiPiece.cell)
	dragEndPos = event.Position

	pMoves := uiPiece.gui.game.currentPos.board[uiPiece.cell].possibleMoves(
		int(uiPiece.cell),
		uiPiece.gui.game.currentPos,
	)
	for _, cMove := range pMoves {
		cellRec := uiPiece.gui.gameContainer.Objects[cMove].(*fyne.Container).Objects[0].(*canvas.Rectangle)
		cellRec.FillColor = PossibleSqBgColor
		cellRec.Refresh()
	}
}

func (uiPiece *UIPiece) DragEnd() {
	r, c := cellIdxToCoordinates(dragStartIdx)
	verticalOffset := dragEndPos.Y / (uiPiece.gui.gameContainer.Size().Height / 8)
	horizontalOffset := dragEndPos.X / (uiPiece.gui.gameContainer.Size().Width / 8)
	newRow := r
	if !(verticalOffset >= 0 && verticalOffset <= 1) {
		newRow = r + int(verticalOffset)
		if verticalOffset < 0 {
			newRow--
		}
	}

	newCol := c
	if !(horizontalOffset >= 0 && horizontalOffset <= 1) {
		newCol = c + int(horizontalOffset)
		if horizontalOffset < 0 {
			newCol--
		}
	}

	newIdx := newRow*8 + newCol
	if dragStartIdx != newIdx {
		uiPiece.gui.game.Move(NewMove(dragStartIdx, newIdx))
	}

	uiPiece.gui.refreshGrid()
}

func NewChessGUI(window fyne.Window, game *Game) *ChessGUI {

	gui := &ChessGUI{
		game:   game,
		window: window,
	}
	grid := container.NewGridWithColumns(8)

	for i := 0; i < 64; i++ {
		grid.Add(container.NewStack(
			canvas.NewRectangle(nil),
			game.currentPos.NewUIPiece(gui, Cell(i))),
		)
	}

	window.SetContent(grid)
	window.Resize(fyne.NewSize(500, 500))

	gui.window = window
	gui.gameContainer = grid
	return gui
}

func (ui *ChessGUI) refreshGrid() {
	for i, cell := range ui.gameContainer.Objects {
		img := cell.(*fyne.Container).Objects[1].(*UIPiece)
		bgCell := cell.(*fyne.Container).Objects[0].(*canvas.Rectangle)

		r, c := cellIdxToCoordinates(i)

		bgCell.FillColor = BlackSqBgColor
		if r%2 != c%2 {
			bgCell.FillColor = WhiteSqBgColor
		}

		p := ui.game.currentPos.board[Cell(i)]
		if p == NoPiece {
			img.Resource = nil
			if img.Visible() {
				img.Hide()
			}
		} else {
			img.Resource = p.toSvg()
			if !img.Visible() {
				img.Show()
			}
		}
		img.Refresh()
	}
}
