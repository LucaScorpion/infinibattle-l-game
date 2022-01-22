package bot

import "infinibattle-l-game/internal/lgame"

// States in which we (red) are scoring, and the opponent cannot score next turn.
var idealStates = []lgame.GameState{
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
