package bot

import "infinibattle-l-game/internal/lgame"

type moveOption struct {
	ourMove       lgame.GameState
	opponentMoves []lgame.GameState
}

func GetNextState(settings lgame.GameSettings, cur lgame.GameState) lgame.GameState {
	ourMoves := lgame.GetPossibleNextStates(settings, cur)

	// Get our possible moves with all possible opponent reactions.
	moveOptions := make([]moveOption, len(ourMoves))
	for i := 0; i < len(ourMoves); i++ {
		moveOptions[i] = moveOption{
			ourMove:       ourMoves[i],
			opponentMoves: lgame.GetPossibleNextStates(settings, ourMoves[i]),
		}
	}

	return moveOptions[0].ourMove
}
