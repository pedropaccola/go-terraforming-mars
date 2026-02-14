package game

type GameStateOpts struct {
	BoardName BoardName
}

type GameState struct {
	Board           *Board
	GlobalParameter map[GlobalParameter]*GlobalParameterTrack
}

func NewGameState(opts GameStateOpts) *GameState {
	return &GameState{
		Board:           NewBoard(opts.BoardName),
		GlobalParameter: NewGlobalParameters(),
	}
}
