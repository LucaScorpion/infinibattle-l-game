package bot

import (
	"infinibattle-l-game/internal/lgame"
	"math/rand"
)

type moveOption struct {
	ourMove       lgame.GameState
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
		// TODO: Check if this actually wins, or if it is dependent on the score.
		// Check for a winning move.
		if len(move.opponentMoves) == 0 {
			return move.ourMove
		}

		// Check if this move increases our score.
		if move.ourMove.Players[cur.PlayerTurn].Score > curScore {
			scoringMoves = append(scoringMoves, move)
		}
	}

	// TODO: Check if getting locked in matters, or if it is dependent on the score.
	// Return the first scoring move that doesn't lock us in.
	for _, move := range scoringMoves {
		if canMoveAfterAnyOpponentMove(settings, &move) {
			return move.ourMove
		}
	}

	// Find any move that doesn't lock us in.
	for _, move := range moveOptions {
		if canMoveAfterAnyOpponentMove(settings, &move) {
			return move.ourMove
		}
	}

	// Uhm... we're fucked I guess? Just return any move.
	return moveOptions[rand.Intn(len(moveOptions))].ourMove
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
