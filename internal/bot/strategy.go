package bot

import "infinibattle-l-game/internal/lgame"

type moveOption struct {
	ourMove       lgame.GameState
	opponentMoves []opponentMoveOption
}

type opponentMoveOption struct {
	move lgame.GameState
	//ourMoves []lgame.GameState // TODO: Is this required?
}

func getNextState(settings lgame.GameSettings, cur lgame.GameState) lgame.GameState {
	moveOptions := buildMoveOptions(settings, cur)

	curScore := cur.Players[cur.PlayerTurn].Score
	var cornerPointMoves []moveOption

	for _, move := range moveOptions {
		// TODO: Check if this actually wins, or if it is dependent on the score.
		// Check for a winning move.
		if len(move.opponentMoves) == 0 {
			return move.ourMove
		}

		// Check if this move increases our score.
		if move.ourMove.Players[cur.PlayerTurn].Score > curScore {
			cornerPointMoves = append(cornerPointMoves, move)
		}
	}

	// TODO: Find best corner point move?
	if len(cornerPointMoves) > 0 {
		return cornerPointMoves[0].ourMove
	}

	// TODO: If all else fails.
	return moveOptions[0].ourMove
}

func buildMoveOptions(settings lgame.GameSettings, cur lgame.GameState) []moveOption {
	ourMoves := lgame.GetPossibleNextStates(settings, cur)

	// Get our possible moves with all possible opponent reactions.
	moveOptions := make([]moveOption, len(ourMoves))
	for i := 0; i < len(ourMoves); i++ {
		opponentOptions := lgame.GetPossibleNextStates(settings, ourMoves[i])

		moveOptions[i] = moveOption{
			ourMove:       ourMoves[i],
			opponentMoves: make([]opponentMoveOption, len(opponentOptions)),
		}

		for j := 0; j < len(opponentOptions); j++ {
			moveOptions[i].opponentMoves[j] = opponentMoveOption{
				move: opponentOptions[j],
			}
		}
	}

	return moveOptions
}
