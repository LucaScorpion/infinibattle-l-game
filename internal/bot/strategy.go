package bot

import (
	"infinibattle-l-game/internal/lgame"
	"math"
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

	// Get some stats from the current state.
	ourPlayer := cur.PlayerTurn
	opponentPlayer := lgame.OtherPlayer(cur.PlayerTurn)
	ourScore := cur.Players[ourPlayer].Score
	opponentScore := cur.Players[opponentPlayer].Score

	// Keep track of all scoring moves, check those first.
	var scoringMoves []moveOption

	for _, move := range moveOptions {
		// Check for a blocking move, but only if our score is higher.
		if len(move.opponentMoves) == 0 && weAreWinning(ourPlayer, move.state) {
			return move.state
		}

		// Store all moves that increase our score.
		if move.state.Players[ourPlayer].Score > ourScore {
			scoringMoves = append(scoringMoves, move)
		}
	}

	// Check if we can move to one of the ideal states.
	if ok, state := checkIdealStates(cur); ok {
		return state
	}

	// Return the best option out of the scoring moves.
	if len(scoringMoves) > 0 {
		return findBestMove(scoringMoves, opponentPlayer, opponentScore).state
	}

	// Return the best option out of the rest.
	return findBestMove(moveOptions, opponentPlayer, opponentScore).state
}

func checkIdealStates(cur lgame.GameState) (bool, lgame.GameState) {
	for _, goal := range allIdealStates {
		// Copy the relevant information and apply it to our actual current state.
		move := cur
		move.Neutrals = goal.Neutrals
		move.Players[cur.PlayerTurn] = goal.Players[goal.PlayerTurn]
		move.PlayerTurn = lgame.OtherPlayer(cur.PlayerTurn)

		// If the move is valid, return it.
		if lgame.IsValidMove(cur, move) {
			return true, move
		}
	}

	return false, lgame.GameState{}
}

func findBestMove(options []moveOption, opponentPlayer lgame.PlayerIndex, opponentScore int) moveOption {
	if len(options) == 0 {
		panic("Cannot find best move when options is empty.")
	}

	// Find the move that gives our opponent the least scoring move possibilities.
	var bestMove moveOption
	var opponentScoringOptionsAfterBest = math.MaxInt

	for _, move := range options {
		// Find how many scoring options the opponent will have.
		opponentScoringOptions := 0
		for _, o := range move.opponentMoves {
			// Check if the opponent score increased.
			if o.move.Players[opponentPlayer].Score > opponentScore {
				opponentScoringOptions++
			}
		}

		// Compare, store the new option.
		if opponentScoringOptions < opponentScoringOptionsAfterBest {
			bestMove = move
			opponentScoringOptionsAfterBest = opponentScoringOptions
		}

		// TODO: Watch out for killer positions?
		// TODO: Try to move towards a state where we can potentially block the opponent in future moves?
	}

	return bestMove
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

func weAreWinning(thisPlayer lgame.PlayerIndex, state lgame.GameState) bool {
	return state.Players[thisPlayer].Score > state.Players[lgame.OtherPlayer(thisPlayer)].Score
}
