package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	app := app.New()
	window := app.NewWindow("Chess")

	game, err := NewGame()
	if err != nil {
		panic(err)
	}

	gui := NewChessGUI(window, game)

	game.Move(NewMove(8, 16))
	gui.refreshGrid()
	game.Move(NewMove(16, 24))
	gui.refreshGrid()

	window.ShowAndRun()
}
