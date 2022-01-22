package bot

import (
	"infinibattle-l-game/internal/lgame"
	"math"
	"os"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// Setup.
	allIdealStates = getAllIdealStateTransforms(lgame.DefaultSettings())

	code := m.Run()
	os.Exit(code)
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
func getDifficultState() lgame.GameState {
	return lgame.GameState{
		PlayerTurn: lgame.PlayerBlue,
		Players: [2]lgame.Player{
			{
				Piece: lgame.LPiece{
					{0, 2},
					{0, 1},
					{0, 0},
					{1, 0},
				},
				Score: 0,
			},
			{
				Piece: lgame.LPiece{
					{1, 2},
					{2, 2},
					{3, 2},
					{3, 3},
				},
				Score: 0,
			},
		},
		Neutrals: [2]lgame.NeutralPiece{
			{2, 0},
			{0, 3},
		},
	}
}

/*
A state where it's possible for red to prevent blue from scoring next turn.

  x 0 1 2 3
y ┌─────────┐
0 │ R □ □ N │
1 │ R R R N │
2 │ □ □ B □ │
3 │ B B B □ │
  └─────────┘

Should become:

  x 0 1 2 3
y ┌─────────┐
0 │ R □ □ N │
1 │ R □ □ □ │
2 │ R R B □ │
3 │ B B B N │
  └─────────┘
*/
func getScorePreventableState() (lgame.GameState, lgame.GameState) {
	return lgame.GameState{
			PlayerTurn: lgame.PlayerRed,
			Players: [2]lgame.Player{
				{
					Piece: lgame.LPiece{
						{0, 1},
						{1, 1},
						{2, 1},
						{0, 0},
					},
				},
				{
					Piece: lgame.LPiece{
						{0, 3},
						{1, 3},
						{2, 3},
						{2, 2},
					},
				},
			},
			Neutrals: [2]lgame.NeutralPiece{
				{3, 0},
				{3, 1},
			},
		},
		lgame.GameState{
			PlayerTurn: lgame.PlayerBlue,
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
						{0, 3},
						{1, 3},
						{2, 3},
						{2, 2},
					},
				},
			},
			Neutrals: [2]lgame.NeutralPiece{
				{3, 0},
				{3, 3},
			},
		}
}

func BenchmarkGetDifficultNextState(b *testing.B) {
	worstTime := 0.0

	for i := 0; i < b.N; i++ {
		startTime := time.Now()

		getNextState(lgame.DefaultSettings(), getDifficultState())

		worstTime = math.Max(worstTime, time.Now().Sub(startTime).Seconds())
	}

	if b.N > 1 {
		b.Logf("Worst time: %.3f seconds", worstTime)
	}
}

func TestGetScorePreventableState(t *testing.T) {
	state, expectedState := getScorePreventableState()
	result := getNextState(lgame.DefaultSettings(), state)

	for i, p := range result.Players {
		if p.Piece != expectedState.Players[i].Piece {
			t.Errorf("Unexpected player %d state.", i)
		}
	}

	for i, n := range result.Neutrals {
		if n != expectedState.Neutrals[i] {
			t.Errorf("Unexpected neutral %d state.", i)
		}
	}

	if t.Failed() {
		t.Log(lgame.DrawState(lgame.DefaultSettings(), result))
	}
}

func BenchmarkGetScorePreventableState(b *testing.B) {
	worstTime := 0.0

	for i := 0; i < b.N; i++ {
		startTime := time.Now()

		state, _ := getScorePreventableState()
		getNextState(lgame.DefaultSettings(), state)

		worstTime = math.Max(worstTime, time.Now().Sub(startTime).Seconds())
	}

	if b.N > 1 {
		b.Logf("Worst time: %.3f seconds", worstTime)
	}
}
