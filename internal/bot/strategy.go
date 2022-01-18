package bot

import (
	"infinibattle-l-game/internal/lgame"
	"math/rand"
)

type moveOption struct {
	state         lgame.GameState
	opponentMoves []opponentMoveOption
}

type opponentMoveOption struct {
	move      lgame.GameState
	ourLMoves []lgame.GameState
}

func getNextState(settings lgame.GameSettings, cur lgame.GameState) lgame.GameState {
	moveOptions := buildMoveOptions(settings, cur)

	curScore := cur.Players[cur.PlayerTurn].Score
	var scoringMoves []moveOption

	for _, move := range moveOptions {
		// Check for a blocking move, but only if our score is higher.
		if len(move.opponentMoves) == 0 && weAreWinning(cur.PlayerTurn, move.state) {
			return move.state
		}

		// Check if this move increases our score.
		if move.state.Players[cur.PlayerTurn].Score > curScore {
			scoringMoves = append(scoringMoves, move)
		}
	}

	// Return the first scoring move that doesn't lock us in (unless our score is higher).
	for _, move := range scoringMoves {
		if isWinningOrFreeMove(settings, &move, cur.PlayerTurn) {
			return move.state
		}
	}

	// Find any move that doesn't lock us in (unless our score is higher).
	for _, move := range moveOptions {
		if isWinningOrFreeMove(settings, &move, cur.PlayerTurn) {
			return move.state
		}
	}

	// Uhm... we're fucked I guess? Just return any move.
	return moveOptions[rand.Intn(len(moveOptions))].state
}

func buildMoveOptions(settings lgame.GameSettings, cur lgame.GameState) []moveOption {
	ourMoves := lgame.GetPossibleNextStates(settings, cur)

	// Get our possible moves with all possible opponent reactions.
	moveOptions := make([]moveOption, len(ourMoves))
	for i := 0; i < len(ourMoves); i++ {
		opponentOptions := lgame.GetPossibleNextStates(settings, ourMoves[i])

		moveOptions[i] = moveOption{
			state:         ourMoves[i],
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

func canMoveAfterAnyOpponentMove(settings lgame.GameSettings, move *moveOption) bool {
	for _, opponentMove := range move.opponentMoves {
		loadOurMovesAfterOpponentMove(settings, &opponentMove)

		if len(opponentMove.ourLMoves) == 0 {
			return false
		}
	}

	return true
}

func loadOurMovesAfterOpponentMove(settings lgame.GameSettings, opponentMove *opponentMoveOption) {
	if opponentMove.ourLMoves == nil {
		// Here we only need to look at our possible L shape moves.
		opponentMove.ourLMoves = lgame.GetLShapeMoves(settings, opponentMove.move)
	}
}

func isWinningOrFreeMove(settings lgame.GameSettings, move *moveOption, thisPlayer lgame.PlayerIndex) bool {
	// We only care about being locked in if we are behind in score.
	return weAreWinning(thisPlayer, move.state) || canMoveAfterAnyOpponentMove(settings, move)
}

func weAreWinning(thisPlayer lgame.PlayerIndex, state lgame.GameState) bool {
	return state.Players[thisPlayer].Score > state.Players[lgame.OtherPlayer(thisPlayer)].Score
}
