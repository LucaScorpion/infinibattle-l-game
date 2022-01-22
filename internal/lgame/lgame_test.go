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
		PlayerTurn: PlayerRed,
		Players: [2]Player{
			{
				Piece: LPiece{
					{1, 0},
					{1, 1},
					{1, 2},
					{0, 2},
				},
				Score: 0,
			},
			{
				Piece: LPiece{
					{2, 1},
					{2, 2},
					{2, 3},
					{3, 3},
				},
				Score: 0,
			},
		},
		Neutrals: [2]NeutralPiece{
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
		PlayerTurn: PlayerBlue,
		Players: [2]Player{
			{
				Piece: LPiece{
					{0, 2},
					{0, 1},
					{0, 0},
					{1, 0},
				},
				Score: 0,
			},
			{
				Piece: LPiece{
					{1, 2},
					{2, 2},
					{3, 2},
					{3, 3},
				},
				Score: 0,
			},
		},
		Neutrals: [2]NeutralPiece{
			{2, 0},
			{0, 3},
		},
	}
}

func TestGetLShapeMoves(t *testing.T) {
	state := getSimpleState()
	expectedStatesCount := 4

	t.Log("Starting state:")
	t.Log(DrawState(DefaultSettings(), state))

	nextStates := GetLShapeMoves(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible L-shape moves:")
	for _, s := range nextStates {
		t.Log(DrawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d L-shape moves but got %d", expectedStatesCount, nextStatesCount)
	}

	if nextStates[0].Players[state.PlayerTurn].Piece == nextStates[1].Players[state.PlayerTurn].Piece {
		t.Errorf("Got multiple states which are the same")
	}

	assertAllStatesValid(t, nextStates)
}

func TestGetNeutralMoves(t *testing.T) {
	state := getSimpleState()
	expectedStatesCount := 14

	t.Log("Starting state:")
	t.Log(DrawState(DefaultSettings(), state))

	nextStates := GetNeutralMoves(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible neutral moves:")
	for _, s := range nextStates {
		t.Log(DrawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d neutral moves but got %d", expectedStatesCount, nextStatesCount)
	}

	assertAllStatesValid(t, nextStates)
}

func TestGetMoves(t *testing.T) {
	state := getSimpleState()
	expectedStatesCount := 56

	t.Log("Starting state:")
	t.Log(DrawState(DefaultSettings(), state))

	nextStates := GetPossibleNextStates(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible moves:")
	for _, s := range nextStates {
		t.Log(DrawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d moves but got %d", expectedStatesCount, nextStatesCount)
	}

	assertAllStatesValid(t, nextStates)
}

func TestGetMovesDifficult(t *testing.T) {
	state := getDifficultState()
	expectedStatesCount := 238

	t.Log("Starting state:")
	t.Log(DrawState(DefaultSettings(), state))

	nextStates := GetPossibleNextStates(DefaultSettings(), state)
	nextStatesCount := len(nextStates)

	t.Log("Possible moves:")
	for _, s := range nextStates {
		t.Log(DrawState(DefaultSettings(), s))
	}

	if nextStatesCount != expectedStatesCount {
		t.Errorf("Expected %d moves but got %d", expectedStatesCount, nextStatesCount)
	}

	assertAllStatesValid(t, nextStates)
}

func assertAllStatesValid(t *testing.T, states []GameState) {
	for _, state := range states {
		occupied := OccupationGrid{}

		for i, p := range state.Players {
			for _, c := range p.Piece {
				if _, ok := occupied[c]; ok {
					t.Errorf("Grid space is already occupied: %d,%d", c.X, c.Y)
				}

				occupied[c] = playerIndexToOccupation[PlayerIndex(i)]
			}
		}

		for _, n := range state.Neutrals {
			if _, ok := occupied[Coordinate(n)]; ok {
				t.Errorf("Grid space is already occupied: %d,%d", n.X, n.Y)
			}

			occupied[Coordinate(n)] = occupiedNeutral
		}

		if len(occupied) != 10 {
			t.Errorf("10 Grid spaces should be occupied, but got %d", len(occupied))
		}
	}
}

func TestIsValidMove(t *testing.T) {
	/*
		From:

		  x 0 1 2 3
		y ┌─────────┐
		0 │ □ R N □ │
		1 │ □ R B □ │
		2 │ R R B □ │
		3 │ □ N B B │
		  └─────────┘
	*/
	simpleStateMoves := map[GameState]bool{
		/*
			  x 0 1 2 3
			y ┌─────────┐
			0 │ R R □ □ │
			1 │ □ R B N │
			2 │ □ R B □ │
			3 │ □ N B B │
			  └─────────┘
		*/
		{
			PlayerTurn: PlayerBlue,
			Players: [2]Player{
				{
					Piece: LPiece{
						{1, 0},
						{1, 1},
						{1, 2},
						{0, 0},
					},
				},
				{
					Piece: LPiece{
						{2, 1},
						{2, 2},
						{2, 3},
						{3, 3},
					},
				},
			},
			Neutrals: [2]NeutralPiece{
				{1, 3},
				{3, 1},
			},
		}: true,
		/*
			  x 0 1 2 3
			y ┌─────────┐
			0 │ N R R R │
			1 │ □ □ B R │
			2 │ □ □ B □ │
			3 │ □ N B B │
			  └─────────┘
		*/
		{
			PlayerTurn: PlayerBlue,
			Players: [2]Player{
				{
					Piece: LPiece{
						{3, 0},
						{2, 0},
						{1, 0},
						{3, 1},
					},
				},
				{
					Piece: LPiece{
						{2, 1},
						{2, 2},
						{2, 3},
						{3, 3},
					},
				},
			},
			Neutrals: [2]NeutralPiece{
				{0, 0},
				{1, 3},
			},
		}: false, // We moved our piece over the neutral piece.
		/*
			  x 0 1 2 3
			y ┌─────────┐
			0 │ R R □ N │
			1 │ □ R B □ │
			2 │ □ R B □ │
			3 │ N □ B B │
			  └─────────┘
		*/
		{
			PlayerTurn: PlayerBlue,
			Players: [2]Player{
				{
					Piece: LPiece{
						{1, 0},
						{1, 1},
						{1, 2},
						{0, 0},
					},
				},
				{
					Piece: LPiece{
						{2, 1},
						{2, 2},
						{2, 3},
						{3, 3},
					},
				},
			},
			Neutrals: [2]NeutralPiece{
				{3, 0},
				{0, 3},
			},
		}: false, // We moved both neutral pieces.
		/*
			  x 0 1 2 3
			y ┌─────────┐
			0 │ □ R N □ │
			1 │ □ R B □ │
			2 │ □ R X □ │
			3 │ □ N B B │
			  └─────────┘
		*/
		{
			PlayerTurn: PlayerBlue,
			Players: [2]Player{
				{
					Piece: LPiece{
						{1, 2},
						{1, 1},
						{1, 0},
						{2, 2},
					},
				},
				{
					Piece: LPiece{
						{2, 1},
						{2, 2},
						{2, 3},
						{3, 3},
					},
				},
			},
			Neutrals: [2]NeutralPiece{
				{2, 0},
				{1, 3},
			},
		}: false, // We overlap with blue.
	}

	for move, valid := range simpleStateMoves {
		if IsValidMove(getSimpleState(), move) != valid {
			validString := "valid"
			if !valid {
				validString = "invalid"
			}
			t.Errorf("Move should be %s:", validString)
			t.Log(DrawState(DefaultSettings(), move))
		}
	}
}
