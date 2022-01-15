package lgame

import (
	"testing"
)

/*
A simple state.
Red's turn.
L-shape moves: 4.
Neutral moves: 14.
Total possible moves: 56.

  x 0 1 2 3
y ┌─────────┐
0 │ □ R N □ │
1 │ □ R B □ │
2 │ R R B □ │
3 │ □ N B B │
  └─────────┘
*/
func getSimpleState() GameState {
	return GameState{
		Turn: PlayerRed,
		Players: []LPiece{
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
		Neutrals: []NeutralPiece{
			{2, 0},
			{1, 3},
		},
	}
}

/*
A more difficult state.
Blue's turn.
L-shape moves: 17.
Neutral moves: 14.
Total possible moves: 238.

  x 0 1 2 3
y ┌─────────┐
0 │ R R N □ │
1 │ R □ □ □ │
2 │ R B B B │
3 │ N □ □ B │
  └─────────┘
*/
func getDifficultState() GameState {
	return GameState{
		Turn: PlayerBlue,
		Players: []LPiece{
			{
				{0, 2},
				{0, 1},
				{0, 0},
				{1, 0},
			},
			{
				{1, 2},
				{2, 2},
				{3, 2},
				{3, 3},
			},
		},
		Neutrals: []NeutralPiece{
			{2, 0},
			{0, 3},
		},
	}
}

func TestGetLShapeMoves(t *testing.T) {
	state := getSimpleState()
	expectedStatesCount := 4

	t.Log("Starting state:")
	t.Log(drawState(DefaultSettings(), state))

	nextStates := getLShapeMoves(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible L-shape moves:")
	for _, s := range nextStates {
		t.Log(drawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d L-shape moves but got %d", expectedStatesCount, nextStatesCount)
	}
}

func TestGetNeutralMoves(t *testing.T) {
	state := getSimpleState()
	expectedStatesCount := 14

	t.Log("Starting state:")
	t.Log(drawState(DefaultSettings(), state))

	nextStates := getNeutralMoves(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible neutral moves:")
	for _, s := range nextStates {
		t.Log(drawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d neutral moves but got %d", expectedStatesCount, nextStatesCount)
	}
}

func TestGetMoves(t *testing.T) {
	state := getSimpleState()
	expectedStatesCount := 56

	t.Log("Starting state:")
	t.Log(drawState(DefaultSettings(), state))

	nextStates := getPossibleNextStates(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible moves:")
	for _, s := range nextStates {
		t.Log(drawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d moves but got %d", expectedStatesCount, nextStatesCount)
	}
}

func TestGetMovesDifficult(t *testing.T) {
	state := getDifficultState()
	expectedStatesCount := 238

	t.Log("Starting state:")
	t.Log(drawState(DefaultSettings(), state))

	nextStates := getPossibleNextStates(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible moves:")
	for _, s := range nextStates {
		t.Log(drawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d moves but got %d", expectedStatesCount, nextStatesCount)
	}
}
