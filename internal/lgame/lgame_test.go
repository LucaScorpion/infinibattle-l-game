package lgame

import (
	"testing"
)

func TestGetLShapeMoves(t *testing.T) {
	/*
		  x 0 1 2 3
		y ┌─────────┐
		0 │ □ R N □ │
		1 │ □ R B □ │
		2 │ R R B □ │
		3 │ □ N B B │
		  └─────────┘
	*/
	state := GameState{
		Players: [2]LPiece{
			{
				{1, 0},
				{1, 1},
				{1, 2},
				{0, 2},
			},
			{
				{2, 1},
				{2, 2},
				{2, 3},
				{3, 3},
			},
		},
		Neutrals: [2]NeutralPiece{
			{2, 0},
			{1, 3},
		},
	}

	t.Log("Starting state:")
	t.Log(drawState(DefaultSettings(), state))

	expectedStatesCount := 4
	lMoves := getLShapeMoves(DefaultSettings(), state, PlayerRed)
	lMovesCount := len(lMoves)

	t.Log("Possible L-shape moves:")
	for _, s := range lMoves {
		t.Log(drawState(DefaultSettings(), s))
	}

	if lMovesCount != expectedStatesCount {
		t.Errorf("Expected %d L-shape moves but got %d", expectedStatesCount, lMovesCount)
	}
}
