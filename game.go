package main

type Game struct {
	currentPos *Position
	positions  []*Position
}

type Move struct {
	startPos int
	endPos   int
}

func NewMove(startPos int, endPos int) Move {
	return Move{
		startPos: startPos,
		endPos:   endPos,
	}
}

// type NotationMove struct {
// 	startPos string
// 	endPos   string
// }

// func (notationMove *NotationMove) ToMove() *Move {

// }

func NewGame() (*Game, error) {
	pos, err := fenToPosition("rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w KQkq - 0 1")
	if err != nil {
		return nil, err
	}

	return &Game{
		currentPos: pos,
		positions:  []*Position{pos},
	}, nil
}

func (game *Game) Move(move Move) {
	oldBoard := game.currentPos.board
	newBoard := oldBoard
	newBoard[Cell(move.endPos)] = oldBoard[Cell(move.startPos)]
	newBoard[Cell(move.startPos)] = NoPiece

	game.currentPos.board = newBoard
	game.positions = append(game.positions, game.currentPos)
}
