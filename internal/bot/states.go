package bot

import "infinibattle-l-game/internal/lgame"

// States in which we (red) are scoring, and the opponent (blue) cannot score on their turn.
// The PlayerTurn is "our" player index, i.e. the piece which we can move.
var idealStates = []lgame.GameState{
	/*
		┌─────────┐
		│ R □ □ N │
		│ R □ B B │
		│ R R □ B │
		│ □ □ N B │
		└─────────┘
	*/
	{
		PlayerTurn: 0,
		Players: [2]lgame.Player{
			{
				Piece: lgame.LPiece{
					{0, 2},
					{0, 1},
					{0, 0},
					{1, 2},
				},
			},
			{
				Piece: lgame.LPiece{
					{3, 1},
					{3, 2},
					{3, 3},
					{2, 1},
				},
			},
		},
		Neutrals: [2]lgame.NeutralPiece{
			{3, 0},
			{2, 3},
		},
	},
	/*
		┌─────────┐
		│ N □ □ N │
		│ □ □ B B │
		│ □ □ R B │
		│ R R R B │
		└─────────┘
	*/
	{
		PlayerTurn: 0,
		Players: [2]lgame.Player{
			{
				Piece: lgame.LPiece{
					{2, 3},
					{1, 3},
					{0, 3},
					{2, 2},
				},
			},
			{
				Piece: lgame.LPiece{
					{3, 1},
					{3, 2},
					{3, 3},
					{2, 1},
				},
			},
		},
		Neutrals: [2]lgame.NeutralPiece{
			{0, 0},
			{3, 0},
		},
	},
}

func getAllIdealStateTransforms(settings lgame.GameSettings) []lgame.GameState {
	var result []lgame.GameState
	for _, state := range idealStates {
		result = append(result, lgame.AllTransforms(settings, state)...)
	}
	return result
}
