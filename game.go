package main

type Game struct {
	currentPos *Position
	positions  []*Position
}

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
